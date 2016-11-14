package main

import (
    "github.com/rwcarlsen/goexif/exif"
    "fmt"
    "os"
)

func main()  {
    fname := "aa-1280x800.jpg"

    f, err := os.Open(fname)
    if err != nil {
        fmt.Println(err)
    }

    // Optionally register camera makenote data parsing - currently Nikon and
    // Canon are supported.
    exif.RegisterParsers(mknote.All...)

    x, err := exif.Decode(f)
    if err != nil {
        fmt.Println(err)
    }

    camModel, _ := x.Get(exif.Model) // normally, don't ignore errors!
    fmt.Println(camModel.StringVal())

    focal, _ := x.Get(exif.FocalLength)
    numer, denom, _ := focal.Rat2(0) // retrieve first (only) rat. value
    fmt.Printf("%v/%v", numer, denom)

    // Two convenience functions exist for date/time taken and GPS coords:
    tm, _ := x.DateTime()
    fmt.Println("Taken: ", tm)

    lat, long, _ := x.LatLong()
    fmt.Println("lat, long: ", lat, ", ", long)
}