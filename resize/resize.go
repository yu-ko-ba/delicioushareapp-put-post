package resize

import (
	"bytes"
	"image"
	"image/jpeg"
	_ "image/jpeg"

	"github.com/tacg0909/meshitero-put-post/calctargetsize"
	"golang.org/x/image/draw"
)

func Resize(imageBuf *bytes.Buffer, maxLength int) (err error) {
    decordedImage, _, err := image.Decode(imageBuf)
    if err != nil {
        return
    }
    rectangle := decordedImage.Bounds()
    width := rectangle.Dx()
    height := rectangle.Dy()
    targetWidth, targetHeight := calctargetsize.CalcTargetSize(width, height, maxLength)
    dst := image.NewRGBA(image.Rect(0, 0, targetWidth, targetHeight))
    draw.CatmullRom.Scale(
        dst,
        dst.Bounds(),
        decordedImage,
        rectangle,
        draw.Over,
        nil,
    )
    err = jpeg.Encode(imageBuf, dst, &jpeg.Options{Quality: 100})
    return
}
