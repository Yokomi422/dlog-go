package meta

import (
	"bufio"
	"io"
	"strings"
)

type Metadata struct {
	Title       string
	Description string
}

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
)

func (m Metadata) ParseMetadata(r io.Reader) (Metadata, error) {
	scanner := bufio.NewScanner(r)

	readMetaLine := func(separator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), separator)
	}
	title := readMetaLine(titlePrefix)
	description := readMetaLine(descriptionPrefix)

	return Metadata{Title: title, Description: description}, nil
}
