package daily

import (
	"io/fs"
	"strconv"
	"strings"
	"time"
)

type Daily struct {
	Path string
	Data []byte
}

func (d Daily) Filter(fileSystem fs.FS, year int) []Daily {
	filteredDailies := make([]Daily, 0)
	dailies, _ := fs.ReadDir(fileSystem, ".")

	for _, d := range dailies {
		if !strings.HasPrefix(d.Name(), strconv.Itoa(year)) {
			continue
		}
		data, _ := fs.ReadFile(fileSystem, d.Name())
		filteredDailies = append(filteredDailies, Daily{Path: d.Name(), Data: data})
	}
	return filteredDailies
}

func (d Daily) currentDate() (year, month, day int) {
	year, _, day = time.Now().Date()
	month = int(time.Now().Month())
	return
}
