package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/corona10/goimagehash"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func ReadTimePositionAsJpeg(inFileName string, seconds int) io.Reader {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(inFileName, ffmpeg.KwArgs{"ss": seconds}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		panic(err)
	}
	return buf
}

func RunReadTimePositionAsJpeg(inputFile string, outputFile string, second int) {
	reader_video := ReadTimePositionAsJpeg(inputFile, second)
	reader_replixa := ReadTimePositionAsJpeg(inputFile, second)
	// calculateHash := NewHasherReader(reader_video)
	// fmt.Print("=========================>")
	// fmt.Print(calculateHash.Hash())
	// fmt.Print("<=========================")
	info := CalculateBasicHashes(reader_replixa)

	// fmt.Println("md5    :", info.Md5)
	// inUnique := insertHash(string(calculateHash.Hash()))
	inUnique := insertHash(info.Md5)
	// inUnique = true
	// data, _ := io.ReadAll(reader)
	// fmt.Println(string(data))
	if inUnique {
		img, err := imaging.Decode(reader_video)
		hash1, _ := goimagehash.AverageHash(img)

		if err != nil {
			log.Fatal(err)
		}
		completeFullPath := outputFile + strconv.Itoa(second) + ".jpg"
		fmt.Println("=========================>")
		fmt.Println(completeFullPath)
		fmt.Println(hash1.GetHash())
		fmt.Println("<=========================")
		err = imaging.Save(img, completeFullPath)
		if err != nil {
			fmt.Println("===============error==========")
			log.Fatal(err)
		}
	}
}

func main() {
	// inputpath := "./inFile/1280.mp4"
	inputpath := "./inFile/a.webm"
	outputPath := "./outFile/"
	for i := 1; i < 15; i++ {
		time.Sleep(time.Second * 1)
		RunReadTimePositionAsJpeg(inputpath, outputPath, i)
	}

}
