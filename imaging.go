package imgutils

import (
	"image"
	"image/draw"

	"github.com/disintegration/imaging"
)

type ConcatMode int

const (
	ConcatHorizontalMode = iota
	ConcatVerticalMode
)

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Crop(img image.Image, min, max image.Point) image.Image {
	return imaging.Crop(img, image.Rectangle{
		Min: min,
		Max: max,
	})
}

func Concatenate(img1, img2 image.Image, mode ConcatMode) image.Image {
	if mode == ConcatVerticalMode {
		return concatenateVert(img1, img2)
	}

	return concatenateHor(img1, img2)
}

func concatenateHor(img1, img2 image.Image) image.Image {
	maxPoint := image.Point{
		X: img1.Bounds().Dx() + img2.Bounds().Dx(),
		Y: maxInt(img1.Bounds().Dy(), img2.Bounds().Dy()),
	}

	// bounds of a second image on the final image
	img2Bounds := image.Rectangle{
		Min: image.Point{
			X: img1.Bounds().Dx(),
			Y: 0,
		},
		Max: image.Point{
			X: maxPoint.X,
			Y: img2.Bounds().Dy(),
		},
	}

	rect := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: maxPoint,
	}

	res := image.NewRGBA(rect)
	draw.Draw(res, img1.Bounds(), img1, image.Point{}, draw.Src)
	draw.Draw(res, img2Bounds, img2, image.Point{}, draw.Src)

	return res
}

func concatenateVert(img1, img2 image.Image) image.Image {
	maxPoint := image.Point{
		X: maxInt(img1.Bounds().Dx(), img2.Bounds().Dx()),
		Y: img1.Bounds().Dy() + img2.Bounds().Dy(),
	}

	// bounds of a second image on the final image
	img2Bounds := image.Rectangle{
		Min: image.Point{
			X: 0,
			Y: img1.Bounds().Dy(),
		},
		Max: image.Point{
			X: img2.Bounds().Dx(),
			Y: maxPoint.Y,
		},
	}

	rect := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: maxPoint,
	}

	res := image.NewRGBA(rect)
	draw.Draw(res, img1.Bounds(), img1, image.Point{}, draw.Src)
	draw.Draw(res, img2Bounds, img2, image.Point{}, draw.Src)

	return res
}
