package daily_test

import (
	"github.com/Yokomi422/dlog-go/daily"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestDaily(t *testing.T) {
	dailies := daily.Daily{}
	vfs := fstest.MapFS{
		"2024-01-01.md": {Data: []byte("test")},
		"2024-01-02.md": {Data: []byte("test")},
		"2025-01-03.md": {Data: []byte("test")},
		"2024-03-04.md": {Data: []byte("test")},
	}
	t.Run("Filter by year", func(t *testing.T) {
		fetchedDailies, err := dailies.Filter(vfs, 2024, -1)

		if err != nil {
			t.Fatal(err)
		}
		expected := []daily.Daily{
			daily.Daily{Path: "2024-01-01.md", Data: []byte("test")},
			daily.Daily{Path: "2024-01-02.md", Data: []byte("test")},
			daily.Daily{Path: "2024-03-04.md", Data: []byte("test")},
		}
		assertFilteredDailies(t, fetchedDailies, expected)
	})
	t.Run("Filter by year and month", func(t *testing.T) {
		fetchedDailies, err := dailies.Filter(vfs, 2024, 3)
		if err != nil {
			t.Fatal(err)
		}
		expected := []daily.Daily{
			daily.Daily{Path: "2024-03-04.md", Data: []byte("test")},
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
