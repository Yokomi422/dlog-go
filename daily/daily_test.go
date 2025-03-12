package daily_test

import (
	"github.com/Yokomi422/dlog-go/daily"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestFetchSpecifiedDailies(t *testing.T) {
	dailies := daily.Daily{}
	vfs := fstest.MapFS{
		"2024-01-01.md": {},
		"2024-01-02.md": {},
		"2025-01-03.md": {},
		"2024-03-04.md": {},
	}
	t.Run("Filter by year and month", func(t *testing.T) {
		fetchedDailies, err := dailies.FetchSpecifiedDailies(vfs, 2024, 3, -1)
		if err != nil {
			t.Fatal(err)
		}
		expected := []daily.Daily{
			{Path: "2024-03-04.md"},
		}
		assertFilteredDailies(t, fetchedDailies, expected)
	})
}

// TODO:refactoring replace reflect.DeepEqual with slices.Equal
func assertFilteredDailies(t *testing.T, filtered []daily.Daily, expected []daily.Daily) {
	t.Helper()
	if !reflect.DeepEqual(filtered, expected) {
		t.Errorf("got %v, want %v", filtered, expected)
	}
}
