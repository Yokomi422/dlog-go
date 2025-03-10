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
		"2024-02-04.md": {Data: []byte("test")},
	}
	t.Run("Filter by year", func(t *testing.T) {
		fetchedDailies := dailies.Filter(vfs, 2024)
		expected := []daily.Daily{
			daily.Daily{Path: "2024-01-01.md", Data: []byte("test")},
			daily.Daily{Path: "2024-01-02.md", Data: []byte("test")},
			daily.Daily{Path: "2024-02-04.md", Data: []byte("test")},
		}
		if !reflect.DeepEqual(fetchedDailies, expected) {
			t.Errorf("got %v, want %v", fetchedDailies, expected)
		}
	})
}
