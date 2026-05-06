package client

type Request struct {
	Method  string
	URL     string
	Body    []byte
	Headers map[string]string
}
