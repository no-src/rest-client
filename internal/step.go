package internal

const (
	stepInit = iota
	stepName
	stepRequestLine
	stepHeader
	stepHeaderEnd
	stepBody
)
