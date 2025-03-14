package daily

import (
	"log"
	"os"
	"path/filepath"
)

func CreateDailyDir(name string) string {
	if _, err := os.Stat(name); err != nil {
		err = os.Mkdir(name, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}

	return filepath.Join(name, "daily")
}
