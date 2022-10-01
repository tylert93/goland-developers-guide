package main

import (
	"coursecontent/demo/pkg/display"
	"coursecontent/demo/pkg/msg"
)

func main() {
	display.Display("Hello from display")
	msg.Hi()
	msg.Emote("An exciting message")
}