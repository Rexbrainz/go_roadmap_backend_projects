package activity_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"go_roadmap_backend_projects/activity"
)

func TestGetActivity_ReturnsLatestUserGithubActivity(t *testing.T) {
	t.Parallel()

	want := &activity.Activity{
		Events: []byte{},
	}
	got, err := activity.GetActivities("Rexbrainz")
	if err != nil {
		t.Fatal(err)
	}

	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
