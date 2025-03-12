package daily

import (
	"fmt"
	"io/fs"
	"regexp"
	"strconv"
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
	r, err := regexp.Compile("(?P<year>[0-9]{4})-(?P<month>[0-9]{2})-(?P<day>[0-9]{2}).md")
	if err != nil {
		return nil, err
	}
	// pattern generate
	if year == -1 && month == -1 {
		r, err = regexp.Compile("(?P<year>[0-9]{4})-(?P<month>[0-9]{2})-(?P<day>[0-9]{2}).md")
		if err != nil {
			return nil, err
		}
	} else if year == -1 {
		r, err = regexp.Compile("(?P<year>[0-9]{4})-" + fmt.Sprintf("%02d", month) + "-[0-9]{2}.md")
		if err != nil {
			return nil, err
		}
	} else if month == -1 {
		r, err = regexp.Compile(strconv.Itoa(year) + "-[0-9]{2}-[0-9]{2}.md")
		if err != nil {
			return nil, err
		}
	} else {
		r, err = regexp.Compile(strconv.Itoa(year) + "-" + fmt.Sprintf("%02d", month) + "-[0-9]{2}.md")
		if err != nil {
			return nil, err
		}
	}

	for _, d := range dailies {
		matched := r.MatchString(d.Name())
		if !matched {
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
