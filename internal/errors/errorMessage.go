package error

type ErrorMessage struct {
	Path      string       `json:"path"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	Fields    []ErrorField `json:"fields"`
}
