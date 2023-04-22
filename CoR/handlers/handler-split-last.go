package handlers

import "strings"

type HandlerSplitLast struct{}

func (h *HandlerSplitLast) ProcessRequest(data string) (response string) {
	// Splits data by " " and returns the last element
	response = strings.Split(data, " ")[len(strings.Split(data, " "))-1]
	return
}
