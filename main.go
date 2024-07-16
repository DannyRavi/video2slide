package main

import (
	"time"
)

func main() {
	// inputpath := "./inFile/1280.mp4"
	inputpath := "./inFile/a.webm"
	outputPath := "./outFile/"
	for i := 1; i < 15; i++ {
		time.Sleep(time.Second * 1)
		RunReadTimePositionAsJpeg(inputpath, outputPath, i)
	}

}
