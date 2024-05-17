package reporting

type Reporter interface {
	Report(report string, reportFormatting func(report string) string) (commentId string, error error)
	Flush() (string, error)
	Suppress() error
	SupportsMarkdown() bool
}
