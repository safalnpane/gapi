package cli



type Header map[string]string

// Server represents information related to an API server
type Server struct {
    Name string
    Description string
    BaseURL string
    Headers []Header
}


func (s *Server) AddHeader(key, value string) {
    newHeader := map[string]string{key: key, value: value}
    s.Headers = append(s.Headers, newHeader)
}
