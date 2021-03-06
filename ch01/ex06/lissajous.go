package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0xFF, 0x00, 0x00, 0xff}, //#ff0000
	color.RGBA{0xFF, 0xA5, 0x00, 0xff}, //#ffa500
	color.RGBA{0xFF, 0xFF, 0x00, 0xff}, //#ffff00
	color.RGBA{0x00, 0x80, 0x00, 0xff}, //#008000
	color.RGBA{0x00, 0xFF, 0xFF, 0xff}, //#00ffff
	color.RGBA{0x00, 0x00, 0xFF, 0xff}, //#0000ff
	color.RGBA{0x80, 0x00, 0x80, 0xff}, //#800080
}

const (
	backgraundIndex = 0 // パレットの最初の色
	redIndex        = 1 // 虹を割り当て
	orangeIndex     = 2
	yellowIndex     = 3
	greenIndex      = 4
	skyIndex        = 5
	blueIndex       = 6
	purpleIndex     = 7
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	filename := "lissajous.gif"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	lissajous(file)
	file.Close()
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // 発振器xが完了する周回の回数
		res     = 0.001 // 回転の分解能
		size    = 100   // 画像キャンパス（-size..+size）
		nframes = 63    // アニメーションフレーム数
		delay   = 8     // 10ms 単位でのフレーム間の遅延
	)
	freq := rand.Float64() * 5.0 // 発振器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			colorIndex := i%7 + 1
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(colorIndex))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
