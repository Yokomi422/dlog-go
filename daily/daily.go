package daily

type Daily struct {
	Path     string
	Metadata Metadata
	Content  Body
}

type Metadata struct {
	Title       string
	Description string
	Tags        []string
}

type Body struct {
	Content string
}
