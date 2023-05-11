package mdhttp

import "net/http"

// ResponseWriterCatcher catches and stores the response
type ResponseWriterCatcher struct {
	http.ResponseWriter
	StatusCode int
	Body       []byte
}

func (r *ResponseWriterCatcher) WriteHeader(sc int) {
	r.StatusCode = sc

	r.ResponseWriter.WriteHeader(sc)
}

func (r *ResponseWriterCatcher) Write(b []byte) (int, error) {
	r.Body = b

	return r.ResponseWriter.Write(b)
}
