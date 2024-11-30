package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"
	// ffmpeg "github.com/u2takey/ffmpeg-go"
)

// var arr = []byte{' ', '`', '.', ',', '*', '~', '+', '&', '#', ' '}
var arr = []byte{' ', '`', '.', ',', '*', '~', '+', '&', '#', '@'}

func main() {

	list, err := os.ReadDir("images")
	if err != nil {
		panic("WHAATT")
	}

	for _, entry := range list {
		if !entry.IsDir() {
			file, err := os.Open("images/" + entry.Name())
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			img, _, err := image.Decode(file)
			if err != nil {
				log.Fatal(err)
			}

			render(img)
		}
	}

}

func render(img image.Image) {
	orig_bounds := img.Bounds().Max

	scale_x := orig_bounds.X / 110
	scale_y := orig_bounds.Y / 55

	new_img_x := orig_bounds.X / scale_x
	new_img_y := orig_bounds.Y / scale_y

	fmt.Print("\033[H\033[J")
	for y := range new_img_y {
		for x := range new_img_x {
			r, g, b, _ := img.At(x*scale_x, y*scale_y).RGBA()
			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
			indx := lum * uint32(len(arr)) / 256
			var cell string
			// if indx == uint32(len(arr)) - 1 {
			// 	cell = fmt.Sprintf("\033[48;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), arr[indx])
			// } else{
			// 	cell = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), arr[indx])
			// }
			
			cell = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), arr[indx])
			fmt.Print(cell)
		}
		fmt.Println()
	}
}
