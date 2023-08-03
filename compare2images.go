package compare2images

import (
	"errors"
	"image"
	"image/color"
	"math"
)

type compResult struct {
	Different uint64
	Percent   float64
	RedGreen  *image.NRGBA
	Faded     *image.NRGBA
}

func compare(img1, img2 *image.Image) (compResult, error) {

	var (
		percent      float64
		err          error = nil
		countDiffPix float64
		greenNRGBA   = color.NRGBA{34, 139, 34, 255}
		redNRGBA     = color.NRGBA{255, 0, 0, 255}
	)

	totalPixels := float64(uint64((*img1).Bounds().Max.X) * uint64((*img1).Bounds().Max.Y))
	RGimg := image.NewNRGBA((*img1).Bounds())
	Fadedimg := image.NewNRGBA((*img1).Bounds())

	for x := (*img1).Bounds().Min.X; x < (*img1).Bounds().Max.X; x++ {
		for y := (*img1).Bounds().Min.Y; y < (*img1).Bounds().Max.Y; y++ {
			if (*img1).At(x, y) != (*img2).At(x, y) {
				countDiffPix++
				RGimg.SetNRGBA(x, y, redNRGBA)
				Fadedimg.SetNRGBA(x, y, redNRGBA)
			} else {
				RGimg.SetNRGBA(x, y, greenNRGBA)
				pixel := color.NRGBAModel.Convert((*img1).At(x, y)).(color.NRGBA)
				pixel.A = (pixel.A / 2)
				Fadedimg.SetNRGBA(x, y, pixel)
			}
		}
	}

	if countDiffPix == 0 {
		percent = 0
	} else {
		percent = roundFloat(float64((countDiffPix/totalPixels)*100), 3)
	}

	if uint64((*img1).Bounds().Max.X)*uint64((*img1).Bounds().Max.Y) != uint64((*img1).Bounds().Max.X)*uint64((*RGimg).Bounds().Max.Y) {
		err = errors.New("Pix len mismatch!")
	}

	return compResult{uint64(countDiffPix), percent, RGimg, Fadedimg}, err
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Compare2images(img1, img2 *image.Image) (compResult, error) {

	var err error = nil
	var result compResult

	if (*img1).Bounds() != (*img2).Bounds() {
		err := errors.New("Images have different resolution or number of pixels")
		return compResult{0, 0, nil, nil}, err
	}

	result, err = compare(img1, img2)
	if err != nil {
		return compResult{0, 0, nil, nil}, err
	}

	return result, err

}
