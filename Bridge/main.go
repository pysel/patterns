// I avoid using factory pattern for simplicity

package main

import "fmt"

func main() {
	windowsOS := &Windows{}
	windowsOS.SetMediaFormat(&AudioFormat{})
	fmt.Println(windowsOS.Play("song.mp3"))

	windowsOS.SetMediaFormat(&VideoFormat{})
	fmt.Println(windowsOS.Play("movie.mp4"))

	linuxOS := &Linux{}
	linuxOS.SetMediaFormat(&AudioFormat{})
	fmt.Println(linuxOS.Play("song.mp3"))

	linuxOS.SetMediaFormat(&VideoFormat{})
	fmt.Println(linuxOS.Play("movie.mp4"))

	archLinuxOS := &ArchLinux{}
	archLinuxOS.SetMediaFormat(&AudioFormat{})
	fmt.Println(archLinuxOS.Play("song.mp3"))
}
