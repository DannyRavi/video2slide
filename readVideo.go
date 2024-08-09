package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type runnerCore interface {
	Execute(arg flags) error
}

func ReadTimePositionAsJpeg(inFileName string, seconds int) io.Reader {
	buf := bytes.NewBuffer(nil)
	percentage := float32(seconds) / float32(theHolder.max_sec)
	stePercentage := fmt.Sprintf("%f", percentage)
	logrus.Info(stePercentage + " %")
	err := ffmpeg.Input(inFileName, ffmpeg.KwArgs{"ss": seconds}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func ReadFrameAsJpeg(inFileName string, frameNum int) io.Reader {
	buf := bytes.NewBuffer(nil)
	percentage := float32(frameNum) / float32(theHolder.max_frame)
	stePercentage := fmt.Sprintf("%f", percentage)
	logrus.Info(stePercentage + " %")
	err := ffmpeg.Input(inFileName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

type RunReadFrameAsJpeg struct{}

func (f *RunReadFrameAsJpeg) Execute(arg flags) error {
	reader_video := ReadFrameAsJpeg(arg.inputFile, arg.second)
	img, err := imaging.Decode(reader_video)
	if err != nil {
		log.Fatal(err)
	}
	arg.counter++
	hash1, _ := goimagehash.AverageHash(img)
	getHash := hash1.ToString()[2:]
	map_keys[arg.counter] = getHash
	inUnique := insertHash(getHash, arg.diffImage, arg.counter)

	if inUnique {
		ImageCounter++
		completeFullPath := arg.outputFile + arg.zeroAdd + strconv.Itoa(ImageCounter) + ".jpg"
		err = imaging.Save(img, completeFullPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

type RunReadTimePositionAsJpeg struct{}

func (f *RunReadTimePositionAsJpeg) Execute(arg flags) error {

	reader_video := ReadTimePositionAsJpeg(arg.inputFile, arg.second)
	img, err := imaging.Decode(reader_video)
	if err != nil {
		log.Fatal(err)
	}
	arg.counter++
	hash1, _ := goimagehash.AverageHash(img)
	getHash := hash1.ToString()[2:]
	map_keys[arg.counter] = getHash
	inUnique := insertHash(getHash, arg.diffImage, arg.counter)

	if inUnique {
		ImageCounter++
		completeFullPath := arg.outputFile + arg.zeroAdd + strconv.Itoa(ImageCounter) + ".jpg"
		err = imaging.Save(img, completeFullPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func totalDurationVideo(fileName string) float32 {
	a, err := ffmpeg.Probe(fileName)
	if err != nil {
		panic(err)
	}
	totalDuration := gjson.Get(a, "format.duration").Float()
	return float32(totalDuration)
}

func getVideoSize(fileName string) int {
	log.Println("Getting video size for", fileName)
	data, err := ffmpeg.Probe(fileName)
	if err != nil {
		panic(err)
	}
	log.Println("got video info", data)
	type VideoInfo struct {
		Streams []struct {
			CodecType string `json:"codec_type"`
			Width     int
			Height    int
			NbFrames  string `json:"nb_frames"`
		} `json:"streams"`
	}
	vInfo := &VideoInfo{}
	err = json.Unmarshal([]byte(data), vInfo)
	if err != nil {
		panic(err)
	}
	for _, s := range vInfo.Streams {

		if s.CodecType == "video" {
			_frame, err := strconv.Atoi(s.NbFrames)
			if err != nil {
				// ... handle error
				panic(err)
			}
			return _frame
		}
	}
	return 0
}
