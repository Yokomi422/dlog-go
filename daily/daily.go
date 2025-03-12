package daily

import (
	"io/fs"
	"regexp"
	"strconv"
	"time"
)

type Daily struct {
	Path string
	Data []byte
}

// FetchSpecifiedDailies retrieves Daily slices from the given fileSystem.
// If year, month, or day is -1, they will be treated as "not specified".
// - year = -1: defaults to current year
// - month = -1: defaults to current month
// - day = -1: means "do not filter by day" (i.e. get the entire month)
func (d Daily) FetchSpecifiedDailies(fileSystem fs.FS, year, month, day int) ([]Daily, error) {
	now := time.Now()
	if year == -1 {
		year = now.Year()
	}
	if month == -1 {
		month = int(now.Month())
	}

	entries, err := fs.ReadDir(fileSystem, ".")

	if err != nil {
		return nil, err
	}
	filtered := make([]Daily, 0)
	var re = regexp.MustCompile(`([0-9]{4})-([0-9]{2})-([0-9]{2}).md`)
	for _, d := range entries {
		name := d.Name()
		if !re.MatchString(name) {
			continue
		}
		match := re.FindStringSubmatch(name)
		fileYear, _ := strconv.Atoi(match[1])
		fileMonth, _ := strconv.Atoi(match[2])
		fileDay, _ := strconv.Atoi(match[3])

		if year != fileYear || month != fileMonth {
			continue
		}
		if day != -1 && day != fileDay {
			continue
		}

		data, err := fs.ReadFile(fileSystem, d.Name())
		if err != nil {
			return nil, err
		}
		filtered = append(filtered, Daily{Path: d.Name(), Data: data})
	}
	return filtered, nil
}
