package daily_test

import (
	"github.com/Yokomi422/dlog-go/daily"
	"path/filepath"

	"testing"
)

func TestCreateDailyDir(t *testing.T) {
	tempDir := t.TempDir()
	dailyDir := daily.CreateDailyDir(tempDir)

	expected := filepath.Join(tempDir, "daily")

	if dailyDir != expected {
		t.Errorf("got %q, want %q", dailyDir, expected)
	}
}
