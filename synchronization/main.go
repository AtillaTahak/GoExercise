package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{} // 2 bytes
	msgch chan string // 128 bytes
}

func NewServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch: make(chan string, 128),
	}
}

func (s *Server) start() {
	fmt.Println("Starting server...")
	s.loop()
}

func (s *Server) stop() {
	s.quitch <- struct{}{}
}
func (s *Server) loop() {
mainloop:
	for{
		select {
			case <-s.quitch:
				fmt.Println("Quitting...")
				break mainloop
			case msg := <-s.msgch:
				s.handleMessage(msg)
		}
	}
	fmt.Println("Loop stopped.")

}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Handling message...", msg)
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func main() {
	sever := NewServer()
	go func() {
		time.Sleep(5 * time.Second)
		sever.stop()
	}()
	sever.start()

}