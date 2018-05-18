package strategy

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

type ImageSquare struct {
	DestinationFilePath string
}

func (i *ImageSquare) Print() error {
	width := 800
	height := 600

	origin := image.Point{X: 0, Y: 0}

	bgImage := image.NewRGBA(image.Rectangle{
		Min: origin,
		Max: image.Point{X: width, Y: height},
	})
	bgColor := image.Uniform{C: color.RGBA{R: 70, G: 70, B: 70, A: 0}}
	quality := &jpeg.Options{Quality: 75}
	draw.Draw(bgImage, bgImage.Bounds(), &bgColor, origin, draw.Src)

	squareWidth := 200
	squareHeight := 200
	squareColor := image.Uniform{C: color.RGBA{R: 255, G: 0, B: 0, A: 1}}
	square := image.Rect(0, 0, squareWidth, squareHeight)
	square = square.Add(image.Point{
		X: (width / 2) - (squareWidth / 2),
		Y: (height / 2) - (squareHeight / 2),
	})
	squareImg := image.NewRGBA(square)
	draw.Draw(bgImage, squareImg.Bounds(), &squareColor, origin, draw.Src)

	w, err := os.Create(i.DestinationFilePath)
	if err != nil {
		return fmt.Errorf("error opening image: %v\n", err)
	}
	defer w.Close()

	if err = jpeg.Encode(w, bgImage, quality); err != nil {
		return fmt.Errorf("error writing image to disk: %v\n", err)
	}
	return nil
}
