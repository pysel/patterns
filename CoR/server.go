package cor

/*
Server is the simulation of a server that can handle requests.

A user will send strings in a form of "handlerType|arbitraryData" to the server.
Assume arbitrary data cannot contain separators (|) or newlines (\n).
*/

type ServerI interface {
	HandleRequest(reqString string) (response string, err error)
}

type Server struct {
	Processor
}

func NewServer() *Server {
	return &Server{*newProcessor()}
}

func (s *Server) HandleRequest(reqString string) (response string, err error) {
	return s.processRequest(reqString)
}
