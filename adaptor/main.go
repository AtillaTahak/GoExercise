package main

import "fmt"

type MusicPlayer interface {
	playMusic(file string)
}

type MP3Player struct {
}

func (m MP3Player) playMusic(file string) {
	fmt.Println("Playing MP3 file", file)
}


type MP4Player struct {
}

func (m MP4Player) playMusic(file string) {
	fmt.Println("Playing MP4 file", file)
}


type MusicPlayerAdapter struct {
	player *MusicPlayer
}

func (m *MusicPlayerAdapter) playMusic(file string) {
	if file[len(file)-3:] == "mp3" {
		MP3Player{}.playMusic(file)
	} else if file[len(file)-3:] == "mp4" {
		MP4Player{}.playMusic(file)
	} else {
		fmt.Println("Invalid media type", file[len(file)-3:])
	}
}

func main() {
	adapter := MusicPlayerAdapter{}
	adapter.playMusic("file.mp3")
	adapter.playMusic("file.mp4")
	adapter.playMusic("file.wav")
}