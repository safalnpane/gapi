package cli


// Server represents information related to an API server
type Server struct {
    Name string
    Description string
    BaseUrl string
    Headers map[string]string
}
