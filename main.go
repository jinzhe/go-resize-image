package main

import (
	"flag"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strconv"
	"strings"

	R "github.com/nfnt/resize"
)

var (
	// P => The target folder to be compressed
	P *string
	// W => Image compression width
	W *string
	// Q => Image compression quality
	Q *string
	// M => Watermark file name (png)
	M *string
	// T => Generate thumbnail size eg:200x200
	T *string
	// WW => Convert width to uint
	WW uint
	// WWW => Convert width(percent) to int
	WWW int
	// QQ =>
	QQ int
	// WT =>
	WT bool
)

func init() {
	P = flag.String("p", "./data/", "The target folder to be compressed")
	W = flag.String("w", "1920", "Image compression width")
	Q = flag.String("q", "80", "Image compression quality")
	M = flag.String("m", "", "Watermark file name (png)")
	T = flag.String("t", "", "Generate thumbnail size")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if strings.Contains(*W, "%") {
		WT = true
		WWuint, _ := strconv.ParseUint(strings.Replace(*W, "%", "", -1), 10, 64)
		WW = uint(WWuint)
	} else {
		WT = false
		WWuint, _ := strconv.ParseUint(*W, 10, 64)
		WW = uint(WWuint)
	}

	QQuint, _ := strconv.ParseInt(*Q, 10, 64)
	QQ = int(QQuint)

	cmd(*P)
}

func cmd(path string) {
	files, _ := ioutil.ReadDir(path)
	for _, file := range files {
		if file.IsDir() {
			cmd(path + file.Name() + "/")
		} else {
			if strings.Contains(strings.ToLower(file.Name()), ".jpg") {
				to := path + file.Name()
				origin := path + file.Name()

				fmt.Println(origin)
				// 根据百分比生成宽度
				if WT {
					fileOrigin, _ := os.Open(origin)
					defer fileOrigin.Close()
					img, _, err := image.DecodeConfig(fileOrigin)
					if err != nil {
						fmt.Println(err)
						return
					}
					WWW = img.Width * int(WW) / 100
					resize(origin, uint(WWW), 0, to)
				} else {
					resize(origin, WW, 0, to)
				}
				if *M != "" {
					watermark(to, strings.Replace(to, ".jpg", "@watermark.jpg", 1))
				}
				if *T != "" {
					TS := strings.Split(*T, "x")
					TWuint, _ := strconv.ParseUint(TS[0], 10, 64)
					THuint, _ := strconv.ParseUint(TS[1], 10, 64)
					TW := uint(TWuint)
					TH := uint(THuint)
					thumbnail(to, TW, TH, strings.Replace(strings.ToLower(to), ".jpg", "@"+*T+".jpg", 1))
				}
			}

		}
	}
}

func resize(file string, width uint, height uint, to string) {
	fileOrigin, _ := os.Open(file)
	origin, _ := jpeg.Decode(fileOrigin)
	defer fileOrigin.Close()

	canvas := R.Resize(width, height, origin, R.Lanczos3)

	fileOutput, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOutput.Close()

	jpeg.Encode(fileOutput, canvas, &jpeg.Options{QQ})
}

func thumbnail(file string, width uint, height uint, to string) {
	fmt.Println(width, height, to)
	fileOrigin, _ := os.Open(file)
	origin, _ := jpeg.Decode(fileOrigin)
	defer fileOrigin.Close()

	canvas := R.Thumbnail(width, height, origin, R.Lanczos3)
	fileOutput, err := os.Create(to)
	if err != nil {
		log.Fatal(err)
	}
	defer fileOutput.Close()
	jpeg.Encode(fileOutput, canvas, &jpeg.Options{QQ})
}

func watermark(file string, to string) {

	fileOrigin, _ := os.Open(file)
	origin, _ := jpeg.Decode(fileOrigin)
	defer fileOrigin.Close()

	fileWatermark, _ := os.Open(*M)
	watermark, _ := png.Decode(fileWatermark)
	defer fileWatermark.Close()

	originSize := origin.Bounds()

	canvas := image.NewNRGBA(originSize)
	draw.Draw(canvas, originSize, origin, image.ZP, draw.Src)
	draw.Draw(canvas, watermark.Bounds().Add(image.Pt(30, 30)), watermark, image.ZP, draw.Over)

	fileOutput, _ := os.Create(to)
	jpeg.Encode(fileOutput, canvas, &jpeg.Options{100})
	defer fileOutput.Close()
}
