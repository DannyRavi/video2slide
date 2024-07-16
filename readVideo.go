package main

import (
	"bytes"
	"io"
	"log"
	"strconv"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
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

func RunReadTimePositionAsJpeg(inputFile string, outputFile string, second int) {
	reader_video := ReadTimePositionAsJpeg(inputFile, second)
	img, err := imaging.Decode(reader_video)
	if err != nil {
		log.Fatal(err)
	}
	hash1, _ := goimagehash.AverageHash(img)
	getHash := hash1.ToString()[2:]
	inUnique := insertHash(getHash)

	if inUnique {
		completeFullPath := outputFile + strconv.Itoa(second) + ".jpg"

		err = imaging.Save(img, completeFullPath)
		if err != nil {
			log.Fatal(err)
		}
	}
}
