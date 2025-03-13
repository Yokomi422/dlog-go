package daily_test

import (
	"github.com/Yokomi422/dlog-go/daily"
	"reflect"
	"strings"
	"testing"
)

func TestParseMetadata(t *testing.T) {
	const body = `Title: test
Description: description
Tags: tag1, tag2`

	metadata := daily.Metadata{}
	got, err := metadata.ParseMetadata(strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	expected := daily.Metadata{
		Title: "test", Description: "description", Tags: []string{"tag1", "tag2"},
	}
	assertMetadata(t, got, expected)
}

func assertMetadata(t *testing.T, got, expected daily.Metadata) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %#v, want %#v", got, expected)
	}
}

func TestParseWholeFile(t *testing.T) {
	const body = `Title: test
Description: description
Tags: tag1, tag2
---
# Heading1
Hello, world!
## Heading2
I'm a test
`
	got := daily.ParseContent(strings.NewReader(body))
	expected := daily.Daily{
		Metadata: daily.Metadata{
			Title:       "test",
			Description: "description",
			Tags:        []string{"tag1", "tag2"},
		},
		Content: daily.Body{
			Content: "# Heading1\nHello, world!\n## Heading2\nI'm a test\n",
		},
	}
	assertDaily(t, got, expected)
}

func assertDaily(t *testing.T, got, expected daily.Daily) {
	t.Helper()
	if reflect.DeepEqual(got, expected) {
		t.Errorf("got %s, want %s", got, expected)
	}
}
