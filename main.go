package main

import (
	"fmt"
	_ "image/jpeg"
	_ "image/png"
)

var frame_chars = []byte{' ', '`', '.', ',', '~', '+', '*', '&', '#', '@'}

func main() {
	// list, err := os.ReadDir("images")
	// if err != nil {
	// 	panic("WHAATT")
	// }

	// list, err := os.ReadDir("tmp_frames")
	// if err != nil {
	// 	panic("WHAATT")
	// }
	//
	// for _, entry := range list {
	// 	if !entry.IsDir() {
	// 		file, err := os.Open("tmp_frames/" + entry.Name())
	// 		defer file.Close()
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 		img, _, err := image.Decode(file)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	//
	// 		render(img)
	// 	}
	// }

	decode_frame()
}

func encode_frame(){
	
}

func decode_frame(){
	//    *
 //      ***
 //     *****
 //    *******
 //   *********
 //  ***********
 // *************
 //      ***
 //      ***
	var byte_arr = []byte{
		' ', 5, 0, 0, 0, 
		'*', 1, 174, 100, 115, 
		' ', 5, 0, 0, 0, '\n',
		' ', 4, 0, 0, 0, 
		'*', 3, 174, 100, 115, 
		' ', 4, 0, 0, 0, '\n',
		' ', 3, 0, 0, 0, 
		'*', 5, 174, 100, 115, 
		' ', 3, 0, 0, 0, '\n',
	}

	// sliding window -> 5 bytes
	// 0 - char
	// 1 - repeat
	// 2 - r
	// 3 - g
	// 4 - b
	// 5 - new line
	for i := 0; i < len(byte_arr); i += 5 {
		if byte_arr[i] == '\n' {
			// or i -= 4
			i += 1
			fmt.Println()
			if i >= len(byte_arr){
				break
			}
		}

		for reps := 0; reps < int(byte_arr[i+1]); reps += 1 {
			r := byte_arr[i+2]
			g := byte_arr[i+3]
			b := byte_arr[i+4]

			var cell string = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), byte_arr[i])
			fmt.Print(cell)
		}

	}
}

// func decode_frame(){
// 	var byte_arr = []byte{'@', 3, 255, 255, 255, '\n'}
// 	// var result string
// 	for i := 0; i < len(byte_arr); i += 5 {
// 		for reps := 0; reps < int(byte_arr[i+1]); reps += 1 {
// 			r := byte_arr[i+2]
// 			g := byte_arr[i+3]
// 			b := byte_arr[i+4]
//
// 			var cell string = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), byte_arr[i])
// 			fmt.Print(cell)
// 		}
//
// 		if byte_arr[i] != 10 {
// 			fmt.Println("NURSE FOUND")
// 			return
// 		}
// 	}
// }

// func render(img image.Image) {
// 	orig_bounds := img.Bounds().Max
//
// 	scale_x := orig_bounds.X / 80
// 	scale_y := orig_bounds.Y / 40
//
// 	new_img_x := orig_bounds.X / scale_x
// 	new_img_y := orig_bounds.Y / scale_y
//
// 	fmt.Print("\033[2J\033[H")
// 	for y := range new_img_y {
// 		for x := range new_img_x {
// 			r, g, b, _ := img.At(x*scale_x, y*scale_y).RGBA()
// 			lum := (19595*r + 38470*g + 7471*b + 1<<15) >> 24
// 			indx := lum * uint32(len(frame_chars)) / 256
// 			var cell string
// 			cell = fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", uint8(r), uint8(g), uint8(b), arr[indx])
// 			fmt.Print(cell)
// 		}
// 		fmt.Println()
// 	}
// }
