package handlers

import (
	"fmt"
)

type Handler interface {
	ProcessRequest(data string) (response string)
}

func GetRequiredHandler(handlerType string) (Handler, error) {
	switch handlerType {
	case "split-last":
		return &HandlerSplitLast{}, nil // better to use factory, avoid for simplicity
	case "split-first":
		return &HandlerSplitFirst{}, nil // ditto
	default:
		return nil, fmt.Errorf("unknown handler type: %s", handlerType)
	}
}
