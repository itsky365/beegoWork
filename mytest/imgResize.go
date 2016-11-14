package main

import (
	"fmt"
    "image"
	"image/jpeg"
	"os"
    "github.com/BurntSushi/graphics-go/graphics"
    "time"
)

func main() {
	file, err := os.Open("IMG_20161107_132231_HDR.jpg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

    file1Name := time.Now().Unix()
	file1, err := os.Create(fmt.Sprintf("test%d.jpg", file1Name))
	if err != nil {
		fmt.Println(err)
	}
	defer file1.Close()

	img, err := jpeg.Decode(file) //解码
	if err != nil {
		fmt.Println(err)
	}

    bound := img.Bounds()
    dx := bound.Dx()
    dy := bound.Dy()
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(dx, dy)

    newdx := 800
    // 上传图片宽度大于200px
    if dx > newdx {
        // 缩略图的大小
        dst := image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
        // 产生缩略图,等比例缩放
        err = graphics.Scale(dst, img)
        if err != nil {
            fmt.Println(err)
        }
        // fmt.Println(dst)
        // header := w.Header()
        // header.Add("Content-Type", "image/jpeg")      

        jpeg.Encode(file1, dst, &jpeg.Options{80})  
    } else {
        jpeg.Encode(file1, img, &jpeg.Options{80})
    }

	// jpeg.Encode(file1, img, &jpeg.Options{5}) //编码，但是将图像质量从100改成5

}