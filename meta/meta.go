package meta

type Metadata struct {
	Title       string
	Description string
}

func (m Metadata) ParseMetadata(text string) (Metadata, error) {

	return m, nil
}
