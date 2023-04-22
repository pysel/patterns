package cor

import (
	"fmt"
	"strings"
)

// Extractor extracts a data from a request and returns it.
type Extractor struct {
	builtRequest *Request
}

// NewExtractor creates a new Extractor.
func NewExtractor() *Extractor {
	return &Extractor{}
}

func (e *Extractor) extractRequest(reqString string) error {
	if !validateReqString(reqString) {
		return fmt.Errorf("invalid <%s> request string", reqString)
	}

	handlerType, data := strings.Split(reqString, "|")[0], strings.Split(reqString, "|")[1]
	e.builtRequest = newRequest(handlerType, data)

	return nil
}

func validateReqString(reqString string) bool {
	if strings.Count(reqString, "|") != 1 && (strings.Index(reqString, "|") == 0 || strings.Index(reqString, "|") != len(reqString)-1) {
		return false
	}

	return true
}
