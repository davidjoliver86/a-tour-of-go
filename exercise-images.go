package main

import (
    "image"
    "image/color"
    "golang.org/x/tour/pic"
)

type Image struct{
    X int
    Y int
}

func (i Image) Bounds() image.Rectangle {
    return image.Rect(0, 0, i.X, i.Y)
}

func (Image) At(x, y int) color.Color {
    // black        red
    //
    // green        yellow
    return color.RGBA{uint8(x), uint8(y), 0, 255}
}

func (Image) ColorModel() color.Model {
    return color.RGBAModel
}

func main() {
    m := Image{255, 255}
    pic.ShowImage(m)
}

