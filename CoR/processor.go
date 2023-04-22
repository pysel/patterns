package cor

import "patterns/CoR/handlers"

type ProcessorI interface {
	ProcessRequest(reqString string) (response string)
}

type Processor struct {
	Extractor
}

func newProcessor() *Processor {
	return &Processor{*NewExtractor()}
}

func (p Processor) processRequest(reqString string) (response string, err error) {
	if err = p.extractRequest(reqString); err != nil {
		return "", err
	}

	handler, err := handlers.GetRequiredHandler(p.builtRequest.handlerType)
	if err != nil {
		return "", err
	}

	response = handler.ProcessRequest(p.builtRequest.data)

	return
}
