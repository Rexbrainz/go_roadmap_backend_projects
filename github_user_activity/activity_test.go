package activity_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go_roadmap_backend_projects/activity"
)

func TestFormatPushEvent(t *testing.T) {
	t.Parallel()

	payload := []byte(`{
		"ref": "refs/heads/main",
		"before": "abc",
		"head": "def"
	}`)

	got, err := activity.FormatPushEvent(payload, "Rexbrainz/repo")
	if err != nil {
		t.Fatal(err)
	}
	want := "Pushed updates to Rexbrainz/repo"
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
