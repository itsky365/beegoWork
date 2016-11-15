package main

import (
	"os"
	// "image"
	// "image/jpeg"
	"github.com/rwcarlsen/goexif/exif"
	"log"
	// "time"
	// "github.com/BurntSushi/graphics-go/graphics"
	"fmt"
)

func main() {
	f1, err := os.Open("aa-1280x800.jpg") //file to scale
	if err != nil {
		panic(err)
	}
	defer f1.Close()
	f3, err := os.Create("aa3.jpg")
	if err != nil {
		panic(err)
	}
	defer f3.Close()
	exif, err := exif.Decode(f1)
	if err != nil {
		log.Println(err)
	}
	//TODO:Rotate image to correct based on exif flag "Orientation"
	a, e := exif.Get("Orientation")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(a)
	}
	// f1.Seek(0, 0)
	// m1, err := jpeg.Decode(f1)
	// if err != nil {
	//     panic(err)
	// }
	// bounds := image.Rect(0, 0, 1200, 800)//1200x800 tumbnail
	// t := time.Now()
	// m := image.NewRGBA(bounds)
	// graphics.Scale(m, m1)
	// tt := time.Now().Sub(t)
	// err = jpeg.Encode(f3, m, &jpeg.Options{90})
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Printf("ok\n", tt)
}
