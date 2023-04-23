package main

import "fmt"

type OS interface {
	GetMediaFormat() MediaFormat
	SetMediaFormat(MediaFormat)
	Play(string) string
}

type Windows struct {
	mediaFormat MediaFormat
}

func (w *Windows) GetMediaFormat() MediaFormat {
	return w.mediaFormat
}

func (w *Windows) SetMediaFormat(mediaFormat MediaFormat) {
	w.mediaFormat = mediaFormat
}

func (w *Windows) Play(filename string) string {
	return fmt.Sprintf("Windows: %s ", w.mediaFormat.Play(filename))
}

type Linux struct {
	mediaFormat MediaFormat
}

func (l *Linux) GetMediaFormat() MediaFormat {
	return l.mediaFormat
}

func (l *Linux) SetMediaFormat(mediaFormat MediaFormat) {
	l.mediaFormat = mediaFormat
}

func (l *Linux) Play(filename string) string {
	return fmt.Sprintf("Linux: %s ", l.mediaFormat.Play(filename))
}

type ArchLinux struct {
	mediaFormat MediaFormat
}

func (a *ArchLinux) GetMediaFormat() MediaFormat {
	return a.mediaFormat
}

func (a *ArchLinux) SetMediaFormat(mediaFormat MediaFormat) {
	a.mediaFormat = mediaFormat
}

func (a *ArchLinux) Play(filename string) string {
	return fmt.Sprintf("Arch Linux: %s ", a.mediaFormat.Play(filename))
}
