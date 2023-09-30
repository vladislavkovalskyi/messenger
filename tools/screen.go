package tools

import (
	"fmt"
	"syscall"
)

var (
	user32           = syscall.NewLazyDLL("user32.dll")
	getSystemMetrics = user32.NewProc("GetSystemMetrics")
)

const (
	ScreenX = 0
	ScreenY = 1
)

func GetScreenResolution(scale float32) (float32, float32, error) {
	// default param
	if scale == .0 {
		scale = 1.0
	}

	screenWidth, _, _ := getSystemMetrics.Call(ScreenX)
	screenHeight, _, _ := getSystemMetrics.Call(ScreenY)

	if screenWidth == 0 || screenHeight == 0 {
		return 0, 0, fmt.Errorf("error getting screen resolution")
	}
	return float32(screenWidth) * scale, float32(screenHeight) * scale, nil
}
