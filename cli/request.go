package cli


// Request represents an API request information
type Request struct {
	Name string
	Description string
	Path string
	Method string
	UseServerHeaders bool
	Headers map[string]string
	Body map[string]string
}


func (r *Request) ReadRequest() error {
	return nil
}


func (r *Request) WriteRequest() error {
    return nil
}
