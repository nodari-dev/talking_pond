package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
)

var arr = [10]byte{' ',  '`', '.', ',' ,'*', '~', '+', '&', '#', '@'}

func main() {
	file, err := os.Open("robert.png")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}


	orig_bounds := img.Bounds().Max

	scale_x := orig_bounds.X / 110
	scale_y := orig_bounds.Y / 55

	new_img_x := orig_bounds.X / scale_x
	new_img_y := orig_bounds.Y / scale_y

	for y := range new_img_y{
		for x := range new_img_x{
			r, g, b, _ := img.At(x * scale_x, y * scale_y).RGBA()
			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
			indx := lum * 10 / 256
			fmt.Printf("%c", arr[indx])
		}
		fmt.Println()
	}

}
