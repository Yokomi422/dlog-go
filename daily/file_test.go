package daily_test

import (
	"fmt"
	"github.com/Yokomi422/dlog-go/daily"
	"path/filepath"
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

func TestCreateDaily(t *testing.T) {
	tempDir := t.TempDir()
	dailyDir := daily.CreateDailyDir(tempDir)

	year, month, day := 2024, 1, 1
	dailyPath, err := daily.CreateDaily(dailyDir, year, month, day)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("daily path:", dailyPath)
	expected := filepath.Join(dailyDir, "2024-01-01.md")

	if dailyPath != expected {
		t.Errorf("got %q, want %q", dailyPath, expected)
	}
}
