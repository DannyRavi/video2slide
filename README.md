[![Build workflow](https://img.shields.io/github/actions/workflow/status/DannyRavi/video2slide/go.yml?style=flat-square)](https://github.com//DannyRavi/DannyRavi/actions/workflows/go.yml?query=branch%3Amain)


# this is video to slide convert

this app can convert any video lecture to slide 
this app currently just run on linux machine.

## install
```

apt install ffmpeg
apt install img2pdf

```

## RUN

use git clone or any other ways to download source code
going to the source directory and run

```go

go run .  -s 1 -i ./inFile/1280.mp4
```

so you can see any frame of video on "out.pdf"

if you run binary file, remember you need "pdf.sh" script beside the binary file and ffmpeg, img2pdf should be install.
