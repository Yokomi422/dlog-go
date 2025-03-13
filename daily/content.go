package daily

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titlePrefix       = "Title: "
	descriptionPrefix = "Description: "
	tagsPrefix        = "Tags: "
)

func (m Metadata) ParseMetadata(r io.Reader) (Metadata, error) {
	scanner := bufio.NewScanner(r)

	readMetaLine := func(separator string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), separator)
	}
	title := readMetaLine(titlePrefix)
	description := readMetaLine(descriptionPrefix)
	tags := make([]string, 0)
	concatTags := readMetaLine(tagsPrefix)
	for _, tag := range strings.Split(concatTags, ",") {
		tags = append(tags, strings.TrimSpace(tag))
	}

	return Metadata{Title: title, Description: description, Tags: tags}, nil
}

func (b Body) ParseBody(scanner *bufio.Scanner) Body {
	// ignore a line
	scanner.Scan()

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	return Body{Content: strings.TrimSuffix(buf.String(), "\n")}
}

func ParseContent(r io.Reader) Daily {
	scanner := bufio.NewScanner(r)
	metadata := Metadata{}
	body := Body{}

	meta, err := metadata.ParseMetadata(r)
	if err != nil {
		panic(err)
	}

	return Daily{Metadata: meta, Content: body.ParseBody(scanner)}
}
