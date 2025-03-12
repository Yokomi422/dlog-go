package meta_test

import (
	"github.com/Yokomi422/dlog-go/meta"
	"reflect"
	"testing"
)

func TestParseMetadata(t *testing.T) {
	const body = `Title: test
Description: description`

	metadata := meta.Metadata{}
	got, err := metadata.ParseMetadata(body)
	if err != nil {
		t.Fatal(err)
	}

	expected := meta.Metadata{
		Title: "test", Description: "description",
	}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("got %#v, want %#v", got, expected)
	}
}
