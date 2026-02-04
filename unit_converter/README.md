# Unit Converter

HTTP service with a small HTML UI for converting length, weight, and temperature units. Includes JSON-free HTTP endpoints plus form-based pages.

## Features
- Convert length, weight, and temperature with query parameters.
- Simple HTML pages for each converter.
- Uses only the Go standard library.

## Supported Units
- Length: `millimeter`, `centimeter`, `meter`, `kilometer`, `inch`, `foot`, `yard`, `mile`
- Weight: `milligram`, `gram`, `kilogram`, `ounce`, `pound`
- Temperature: `celsius`, `fahrenheit`, `kelvin`

## Run
```bash
cd unit_converter
go run ./cmd/converter
```
Server listens on `:4000`.

## HTTP API
Query params: `value`, `from`, `to`

Examples:
```bash
curl "http://localhost:4000/convert/length?value=10&from=meter&to=foot"
curl "http://localhost:4000/convert/weight?value=2.5&from=kilogram&to=pound"
curl "http://localhost:4000/convert/temperature?value=100&from=celsius&to=fahrenheit"
```

Endpoints:
- `GET /convert/length`
- `GET /convert/weight`
- `GET /convert/temperature`

## HTML UI
- `GET /` index page
- `GET /length`
- `GET /weight`
- `GET /temperature`

## Tests
```bash
cd unit_converter
go test ./...
```

## Project Source
This project is one of the tasks from https://roadmap.sh/projects/unit-converter
