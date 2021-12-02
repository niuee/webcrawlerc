package parser

type parsedContent struct {
	Urls        []string
	Title       string
	Description string
	Paragraphs  []string
	OtherTags   map[string][]string
}
