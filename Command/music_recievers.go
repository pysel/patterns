package main

import "fmt"

type Music struct{}

func (Music) on() {
	fmt.Println("Music is ON")
}

func (Music) off() {
	fmt.Println("Music is OFF")
}

type MusicOnCommand struct {
	Command
	Music
}

func (MOC MusicOnCommand) exec() {
	MOC.Music.on()
}

type MusicOffCommand struct {
	Command
	Music
}

func (MOC MusicOffCommand) exec() {
	MOC.Music.off()
}
