package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
)

type Image struct {
	w, h  int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.w, img.h)
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{
		uint8(x*y),
    uint8(x+y),
		255,
		255,
	}
}

func main() {
	m := Image{48, 48}
	pic.ShowImage(m)
}

// $ go run exercise-images.go
// IMAGE:iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAIAAADYYG7QAAACAklEQVR4nOyXiW7iMBBAA8suex+BBRLI0dwQSP//62g0rl0be2I7hBKpSK/WzHhon8ZOI2aOc3YmzniYtT+TMfEqNB0Nb0KfxoEgNBsDF0Kf741C6MtdmTjz87xdRoN6Qu3G19uQ6xq6hL4NQWzZrxH63peg7wf1Qj9s8C37ZYyEfhqwMWvr5tlc6BfOqnNXSy2mFkK/JZaqoiEVUrcT+sOxEFNDcl2DtdBfwKWBOalZWx+hfzbElv3WQi58zDUjMu5saWC1E1rQCbkQdxDoGhi1mFoILbkjI0JLFVukLrNXFXveISb0X8SXKkpKfMtUaIULrSgeF2PkuoYBJtT+ljW8OtadpLqGtfmrY6MT2uhIdA0nGgwwIY9OyFMRI3VGLaZ6Ib9TyOeOzIOUJ5IqPHtV8aoJbcU7RIS2lJCLZSqkPuSlZkI7+Me4QyjwrZ1WKMCFAlwowMnxrQbW4ScU0gmFEpmqSDjRoKdQhAhF3JGFkDJSMWXUYtolFCNCMTKhWLxDRCgGEhpccJAqN7zUTOhJRYXUBxNKEKEE/kwiUkoVgt23DiKUqoRSZEIpnVACMaHgYp4G1hseWcYdGRHK4LHPJI5cPH2fO8QLFXBkFbw6DvCUHeGxb/odmYlQoRIqOCHGQUyL64VKswmV0oRKmFApMbojewg9hB5CH07oJQAA//8lkZt2rEUdUgAAAABJRU5ErkJggg==
