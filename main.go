package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
)

var flagvar int
var ImageCouner int = 0

type flags struct {
	inputFile  string
	outputFile string
	time_sec   int
	duration   int
}

var myFlag = flags{}

func init() {

	flag.IntVar(&myFlag.time_sec, "s", 1, "split time per second")
	flag.StringVar(&myFlag.inputFile, "i", "./inFile/1280.mp4", "please insert video path ")
	flag.StringVar(&myFlag.outputFile, "o", "./outFile/", "please insert output file path ")
	flag.IntVar(&myFlag.duration, "d", -1, "please insert duration of video ")
	flag.Parse()
	fmt.Println(myFlag.time_sec, myFlag.inputFile, myFlag.outputFile)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {
	// inputpath := "./inFile/1280.mp4"
	// var nFlag = flag.Int("n", 1234, "help message for flag n")
	// flag.Parse()

	fmt.Println(myFlag.outputFile)
	inputpath := myFlag.inputFile
	outputPath := myFlag.outputFile
	_duration := myFlag.duration
	durableAnyVideo := maxDurationPerSecond(inputpath, _duration)
	for i := 1; i < durableAnyVideo; i++ {
		_sec := myFlag.time_sec * i
		zeroAdder := totalDurationCalculate(inputpath, ImageCouner)
		log.Info(_sec)
		if _sec >= durableAnyVideo {
			break
		}
		RunReadTimePositionAsJpeg(inputpath, outputPath, _sec, zeroAdder)
	}
	execute()
	cleanOutPut(outputPath)

}
