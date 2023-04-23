package main

import "fmt"

type MediaFormat interface {
	Play(string) string
}

type AudioFormat struct{}

func (a *AudioFormat) Play(filename string) string {
	return fmt.Sprintf("Playing audio file: %s", filename)
}

type VideoFormat struct{}

func (v *VideoFormat) Play(filename string) string {
	return fmt.Sprintf("Playing video file: %s", filename)
}
