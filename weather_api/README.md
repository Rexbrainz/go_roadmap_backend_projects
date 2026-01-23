# Weather API

Lightweight HTTP API that proxies OpenWeatherMap and caches responses (in-memory by default, optional Redis). Provides current weather for a given city via a single endpoint.

## Features
- `GET /weather?city={CITY,COUNTRY}` returns current conditions (temperature, min/max, summary, description).
- Caching layer: in-memory map; pluggable Redis adapter with TTL.
- Graceful error handling for missing city, provider failures, and malformed provider responses.
- Health check at `GET /ping`.

## Architecture
- **Handler layer** (`handlers.go`): HTTP endpoints, validation, JSON responses.
?- **Service layer** (`weather.go`): orchestrates fetch + cache, hides transport details.
- **Adapters** (`store.go`, `weather.go`): OpenWeatherMap client implements `WeatherFetcher`; in-memory and Redis stores implement `Storer`.
- **Models** (`weather_model.go`): typed structs for provider response and API output.
- **Errors** (`errors.go`): reusable typed errors + HTTP mapping.
- **Tests** (`weather_test.go`): cover caching behavior, client fetching, and parsing.

## Setup
1) Ensure Go installed (module targets Go 1.25.x).  
2) Get an OpenWeatherMap API key: https://openweathermap.org/api  
3) Export the key (used by the client):  
   ```bash
   export OPENWEATHER_API_KEY=<your_key>
   ```
4) (Optional) Redis cache: set `REDIS_URL` (e.g. `redis://localhost:6379/0`). If unset, the server uses in-memory cache.

## Run
```bash
cd weather_api
go run ./cmd/server   # if you add a main; currently start via go test/ custom harness
```

> The library exposes `weather.StartServer(app *Application)`. A minimal `main.go` could be:
> ```go
> package main
> import ("log"; "github.com/Rexbrainz/weather/weather")
> func main() {
>   key := os.Getenv("OPENWEATHER_API_KEY")
>   if key == "" { log.Fatal("OPENWEATHER_API_KEY required") }
>   fetcher := weather.NewClient(key)
>   store := weather.NewMemStore()
>   app := weather.NewWeatherApp(fetcher, store)
>   weather.StartServer(app)
> }
> ```

## API Reference

### GET /weather
- **Query params:** `city` (required) — format `City,CountryCode` e.g. `London,UK` or `Abia,NG`.
- **Spaces:** URL-encode spaces in city names (e.g. `New%20York,US` or `New+Delhi,IN`), otherwise the server will treat the request as malformed.
- **Examples:**
  - `curl "http://localhost:4000/weather?city=London,UK"`
  - `curl "http://localhost:4000/weather?city=New%20York,US"`
  - `curl --get --data-urlencode "city=New York,US" http://localhost:4000/weather`
- **Response 200:**
  ```json
  {
    "city": "Abia",
    "country": "NG",
    "temp": 311.69,
    "temp_min": 311.69,
    "temp_max": 311.69,
    "condition": "Clouds",
    "description": "overcast clouds"
  }
  ```
- **Error responses:**
  - `400 Bad Request` when `city` is missing.
  - `502 Bad Gateway` when the weather provider is unavailable.
  - `206 Partial Content` when provider response is malformed.
  - `500 Internal Server Error` for unexpected cases.

### GET /ping
- Returns `200 OK` with body `OK` for health checks.

## Caching
- In-memory cache: `MemStore` keyed by city; TTL set by the service (currently 15 minutes).
- Redis cache: `RedisStore` uses key pattern `weather:{city}`; failures are non-fatal (cache is best-effort).

## Testing
```bash
cd weather_api
go test ./...
```
Tests use fakes for fetcher/store plus an httptest server for the client.

## Notes
- Handlers set headers before writing status, preventing write-order bugs.
- Context passed through to the fetcher; if the client disconnects the request is canceled upstream.
