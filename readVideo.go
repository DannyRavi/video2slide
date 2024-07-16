package main

import (
	"bytes"
	"io"
	"log"
	"strconv"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	"github.com/tidwall/gjson"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ReadTimePositionAsJpeg(inFileName string, seconds int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName, ffmpeg.KwArgs{"ss": seconds}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func RunReadTimePositionAsJpeg(inputFile string, outputFile string, second int, zeroAdd string) {
	reader_video := ReadTimePositionAsJpeg(inputFile, second)
	img, err := imaging.Decode(reader_video)
	if err != nil {
		log.Fatal(err)
	}
	hash1, _ := goimagehash.AverageHash(img)
	getHash := hash1.ToString()[2:]
	inUnique := insertHash(getHash)

	if inUnique {
		ImageCouner++
		completeFullPath := outputFile + zeroAdd + strconv.Itoa(ImageCouner) + ".jpg"
		err = imaging.Save(img, completeFullPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func totalDurationVideo(fileName string) float32 {
	a, err := ffmpeg.Probe(fileName)
	if err != nil {
		panic(err)
	}
	totalDuration := gjson.Get(a, "format.duration").Float()
	return float32(totalDuration)
}

