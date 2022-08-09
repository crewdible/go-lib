package stringlib

import (
	"bytes"
	"fmt"
	"image/color"

	b64 "encoding/base64"

	svg "github.com/ajstarks/svgo"
	"github.com/boombuler/barcode"
	c128 "github.com/boombuler/barcode/code128"
)

func GenerateSvgBarcode128(content string, width, height int) (string, error) {
	buf := new(bytes.Buffer)
	canvas := svg.New(buf)

	// Create the barcode
	b, err := c128.Encode(content)
	if err != nil {
		return "", err
	}

	// Scale the barcode to 200x200 pixels
	imgWidth := b.Bounds().Dx() * width
	imgHeight := b.Bounds().Dy() * height
	barScale, err := barcode.Scale(b, imgWidth, imgHeight)
	if err != nil {
		return "", err
	}

	bounds := barScale.Bounds()
	canvas.Start(bounds.Dx(), imgHeight)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		if barScale.At(x, bounds.Min.Y) == color.Black {
			start := x
			x++

			for x < bounds.Max.X && barScale.At(x, bounds.Min.Y) == color.Black {
				x++
			}

			canvas.Rect(start, 0, x-start, imgHeight, "fill:black")
		}
	}

	canvas.End()

	sEnc := b64.StdEncoding.EncodeToString(buf.Bytes())
	svgBc := fmt.Sprintf("data:image/svg+xml;base64,%s", sEnc)

	return svgBc, err

}
