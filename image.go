package main

import (
	"image"
	"log"
)

func analyzeAlpha(img image.Image) (bool, bool) {
	skip := true
	hasAlphaPixel := false
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			_, _, _, a := img.At(x, y).RGBA()
			if a == 65535 || a == 0xff {
				skip = false
			} else {
				hasAlphaPixel = true
			}
			// XXX could return early if one non-alpha pixel and one alpha pixel has been found.
		}
	}
	log.Println(skip)
	return skip, hasAlphaPixel
}