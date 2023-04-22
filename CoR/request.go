package cor

type Request struct {
	handlerType string
	data        string
}

// NewRequest creates a new Request.
func newRequest(handlerType string, data string) *Request {
	return &Request{
		handlerType: handlerType,
		data:        data,
	}
}
