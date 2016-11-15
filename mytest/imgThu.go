package main

import (
	"github.com/BurntSushi/graphics-go/graphics"
	"github.com/BurntSushi/graphics-go/graphics/graphicstest"
	"image"

	"fmt"
	"image/png"
	"os"
	"time"
)

func main() {
	dst := image.NewRGBA(image.Rect(0, 0, 80, 80))

	src, err := graphicstest.LoadImage("./testdata/gopher.png")
	if err != nil {
		fmt.Println(err)
	}
	if err := graphics.Thumbnail(dst, src); err != nil {
		fmt.Println(err)
	}

	file1Name := time.Now().Unix()
	file1, err := os.Create(fmt.Sprintf("test%d.png", file1Name))
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	png.Encode(file1, dst)

	// cmp, err := graphicstest.LoadImage("./testdata/gopher-thumb-80x80.png")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = graphicstest.ImageWithinTolerance(dst, cmp, 0)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
