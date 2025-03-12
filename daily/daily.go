package daily

import (
	"fmt"
	"io/fs"
	"strconv"
	"strings"
	"time"
)

type Daily struct {
	Path string
	Data []byte
}

// Filter returns Daily slice filtered by year and month
// year and month is optional, which means no year returns dailies which have
// the assigned month. If you want to get all dailies, you should set year and month
// to be -1.
func (d Daily) Filter(fileSystem fs.FS, year, month int) ([]Daily, error) {
	filteredDailies := make([]Daily, 0)
	dailies, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}

	for _, d := range dailies {
		if !strings.HasPrefix(d.Name(), strconv.Itoa(year)) {
			continue
		}
		monthFormat := fmt.Sprintf("%02d", month)
		if month != -1 && !strings.Contains(d.Name(), monthFormat) {
			continue
		}
		data, _ := fs.ReadFile(fileSystem, d.Name())
		filteredDailies = append(filteredDailies, Daily{Path: d.Name(), Data: data})
	}
	return filteredDailies, nil
}

func (d Daily) currentDate() (year, month, day int) {
	year, _, day = time.Now().Date()
	month = int(time.Now().Month())
	return
}
