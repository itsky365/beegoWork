package main

import (
	"fmt"
    "image"
	"image/jpeg"
	"os"
    "github.com/BurntSushi/graphics-go/graphics"
    "time"
    //"image/png"
    //"image/gif"
    "path"
    "image/png"
    "image/gif"
)


// 读取图片文件
func LoadImage(imgPath string) (image.Image, error) {
    file, err := os.Open(imgPath)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    defer file.Close()
    
    img, _, err := image.Decode(file)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    return img, err
}

// 保存图片文件
func saveImage(path string, img image.Image, fileExt string) (err error) {
    imgfile, err := os.Create(path)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer imgfile.Close()
    
    // 以PNG格式保存文件
    if fileExt == ".jpg" {
        err = jpeg.Encode(imgfile, img, &jpeg.Options{80})
    } else if fileExt == ".png" {
        err = png.Encode(imgfile, img)
    } else if fileExt == ".gif" {
        err = gif.Encode(imgfile, img, &gif.Options{NumColors: 256})
    }
    if err != nil {
        fmt.Println(err)
        return
    }
    return
}

func JpegImage() {
    
    //pathStr := "./testdata/gopher.png"
    //pathStr := "IMG_20161107_132231_HDR.jpg"
    pathStr := "20131118201009_nFLNJ.gif"
    
    img, err := LoadImage(pathStr)
    if err != nil {
        fmt.Println(err)
    }
    
    fileExt := path.Ext(path.Base(pathStr))
    fmt.Println(fileExt)
    
    file1Name := time.Now().Unix()
    file1 := fmt.Sprintf("test%d%s", file1Name, fileExt)
    
    bound := img.Bounds()
    dx := bound.Dx()
    dy := bound.Dy()
    fmt.Println(dx, dy)
    
    newdx := 500
    // 上传图片宽度大于200px
    if dx > newdx {
        // 缩略图的大小
        dst := image.NewRGBA(image.Rect(0, 0, newdx, newdx*dy/dx))
        // 产生缩略图,等比例缩放
        err = graphics.Scale(dst, img)
        if err != nil {
            fmt.Println(err)
        }
        
        //jpeg.Encode(file1, dst, &jpeg.Options{80})
        saveImage(file1, dst, fileExt)
    } else {
        //jpeg.Encode(file1, img, &jpeg.Options{80})
        saveImage(file1, img, fileExt)
    }
}



func main() {
    //fileExt := path.Ext(path.Base("IMG_20161107_132231_HDR.jpg"))
    //fmt.Println(fileExt)
    JpegImage()
    
	// jpeg.Encode(file1, img, &jpeg.Options{5}) //编码，但是将图像质量从100改成5

}