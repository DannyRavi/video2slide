package main

import (
	"flag"
	"fmt"
)

var flagvar int

type flags struct {
	inputFile  string
	outputFile string
	time_sec   int
}

var myFlag = flags{}

func init() {

	flag.IntVar(&myFlag.time_sec, "s", 1, "split time per second")
	flag.StringVar(&myFlag.inputFile, "i", "./inFile/a.webm", "please insert video ")
	flag.StringVar(&myFlag.outputFile, "o", "./outFile/", "please insert output file ")
	flag.Parse()
	fmt.Println(myFlag.time_sec, myFlag.inputFile, myFlag.outputFile)
}

func main() {
	// inputpath := "./inFile/1280.mp4"
	// var nFlag = flag.Int("n", 1234, "help message for flag n")
	// flag.Parse()
	fmt.Println(myFlag.outputFile)
	inputpath := "./inFile/a.webm"
	outputPath := "./outFile/"
	for i := 1; i < 10; i++ {
		RunReadTimePositionAsJpeg(inputpath, outputPath, i)
	}

}
