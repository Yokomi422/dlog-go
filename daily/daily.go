package daily

type Daily struct {
	Path     string
	Metadata Metadata
	Content  Content
}

type Metadata struct {
	Title       string
	Description string
	Tags        []string
}

type Content struct {
	Body string
}
