package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func totalDurationCalculate(inputPath string, locImageCounter int) string {
	_totalVideoDuration := totalDurationVideo(inputPath)
	splitTotal := int(_totalVideoDuration / float32(myFlag.second))
	lenSplitTotal := len(strconv.Itoa(splitTotal)) + 1
	preload := len(strconv.Itoa(locImageCounter + 1))
	updateZero := lenSplitTotal - preload
	prefixZero := ""
	for i := 1; i <= updateZero; i++ {
		prefixZero = prefixZero + "0"
	}
	return prefixZero
	// fmt.Println("---------------> ", lenSplitTotal)
}

func maxDurationPerSecond(inputPath string, durable int) int {
	_totalVideoDuration := int(totalDurationVideo(inputPath))
	if durable == -1 {
		return _totalVideoDuration
	} else if durable > _totalVideoDuration {
		log.Panic("duration is greater than total duration of video")
		return _totalVideoDuration
	} else {
		return durable
	}

}

func execute() {

	cmd, err := exec.Command("/bin/sh", "./pdf.sh").Output()
	if err != nil {
		fmt.Printf("error %s", err)
	}
	output := string(cmd)
	log.Info(output)
}

func cleanOutPut(outPath string) {
	// v := outPath + "*.jpg"
	_path := outPath + "*.jpg"
	files, err := filepath.Glob(_path)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if err := os.Remove(f); err != nil {
			panic(err)
		}
	}
}

func runMode(condition bool) runnerCore {
	if condition {
		return &RunReadTimePositionAsJpeg{}
	} else {

		return &RunReadFrameAsJpeg{}
	}

}

func runner(isParallel bool, f runnerCore, arg flags) {
	if isParallel {
		go func() {
			err := f.Execute(arg)
			if err != nil {
				logrus.Error(err)
			}
		}()
	} else {
		err := f.Execute(arg)
		if err != nil {
			logrus.Error(err)
		}
	}
}
