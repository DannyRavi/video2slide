package main

import (
	"fmt"
	"os/exec"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func totalDurationCalculate(inputPath string, locImageCounter int) string {
	_totalVideoDuration := totalDurationVideo(inputPath)
	splitTotal := int(_totalVideoDuration / float32(myFlag.time_sec))
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

func execute(cmd string, args string) {
	out, err := exec.Command(cmd, args).Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	log.Info("Command Successfully Executed")
	rec_ := string(out[:])
	log.Info(rec_)
}
