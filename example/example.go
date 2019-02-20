package main

import (
	"image"
	"image/jpeg"
	"os"

	"github.com/ShoshinNikita/imgutils"
)

func main() {
	var img1, img2, img3 image.Image

	f, _ := os.Open("data/1.jpeg")
	img1, _ = jpeg.Decode(f)
	f.Close()
	f, _ = os.Open("data/2.jpeg")
	img2, _ = jpeg.Decode(f)
	f.Close()
	f, _ = os.Open("data/3.jpeg")
	img3, _ = jpeg.Decode(f)
	f.Close()

	res := imgutils.Concatenate(img1, img2, imgutils.ConcatHorizontalMode)
	res = imgutils.Concatenate(res, img3, imgutils.ConcatVerticalMode)

	f, _ = os.Create("result.jpeg")
	jpeg.Encode(f, res, nil)
}
