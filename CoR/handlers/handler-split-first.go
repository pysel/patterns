package handlers

import "strings"

type HandlerSplitFirst struct{}

func (h *HandlerSplitFirst) ProcessRequest(data string) (response string) {
	// Splits data by " " and returns the first element
	response = strings.Split(data, " ")[0]
	return
}
