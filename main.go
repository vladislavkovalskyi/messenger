package main

import (
	"messenger/gui"
	"messenger/tools"
)

func main() {
	println(tools.GetScreenResolution())
	gui.Application()
}
