package main

import (
	"flag"
	"fmt"

	"github.com/corona10/goimagehash"
	log "github.com/sirupsen/logrus"
)

var flagvar int
var ImageCounter int = 0

type flags struct {
	inputFile     string
	outputFile    string
	second        int
	duration      int
	frame         int
	frame_or_time bool
	zeroAdd       string
	parallel      bool
	diffImage     int
	hash_old      *goimagehash.ImageHash
}

type holder struct {
	itr       int
	itr_ref   int
	max_frame int
	max_sec   int
}

var myFlag = flags{}
var theHolder = holder{}

func init() {

	flag.IntVar(&myFlag.second, "s", 1, "split time per second")
	flag.IntVar(&myFlag.frame, "f", 1, "split frame")
	flag.IntVar(&myFlag.diffImage, "k", 0, "difference set between 1 to 50")
	flag.StringVar(&myFlag.inputFile, "i", "./inFile/1280.mp4", "please insert video path ")
	flag.StringVar(&myFlag.outputFile, "o", "./outFile/", "please insert output file path ")
	flag.IntVar(&myFlag.duration, "d", -1, "please insert duration of video ")
	flag.BoolVar(&myFlag.frame_or_time, "frame", false, "false is time (second) mode - true is frame mode")
	flag.BoolVar(&myFlag.parallel, "p", false, "run parallel")
	flag.Parse()
	fmt.Println(myFlag.second, myFlag.inputFile, myFlag.outputFile)
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)

	theHolder.itr = 0
	if myFlag.frame_or_time {
		theHolder.itr_ref = 0
	} else {

		theHolder.itr_ref = 3
	}

}

func main() {

	setPointTime := myFlag.second
	fmt.Println(myFlag.frame_or_time)
	state := runMode(myFlag.frame_or_time)
	theHolder.max_frame = getVideoSize(myFlag.inputFile)
	log.Println(theHolder.max_frame)
	durableAnyVideo := maxDurationPerSecond(myFlag.inputFile, myFlag.duration)
	theHolder.max_sec = durableAnyVideo
	if !myFlag.frame_or_time {
		durableAnyVideo = theHolder.max_frame
	}

	for i := 1; i < durableAnyVideo; i++ {
		myFlag.second = setPointTime * i

		myFlag.zeroAdd = totalDurationCalculate(myFlag.inputFile, ImageCounter)
		log.Info(myFlag.second)
		if myFlag.second >= durableAnyVideo {
			break
		}

		if state != nil {
			runner(myFlag.parallel, state, myFlag)
		}
		// go RunReadTimePositionAsJpeg(myFlag.inputFile, myFlag.outputFile, _sec, myFlag.zeroAdd)
	}
	execute()
	cleanOutPut(myFlag.outputFile)

}
