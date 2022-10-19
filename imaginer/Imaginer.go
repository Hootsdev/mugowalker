package imaginer

import (
	"fmt"
	"image"
	"image/draw"

	"github.com/keritos/tesseract"
	// "github.com/otiai10/gosseract/v2"
	log "github.com/sirupsen/logrus"
	"github.com/vitali-fedulov/images/v2"
)

type Cutter interface {
	Concat(image.Image, int, int)
}

type TextScanner interface {
	ImageText(image.Image) (string, error)
}

type Similizer interface {
	Similarity(image.Image, image.Image) (similar bool, percent int)
}

func Concat(img1 image.Image, x1, y1, x2, y2 int) image.Image {
	sr := image.Rect(x1, y1, x2, y2)
	rect := image.Rectangle{image.Point{}, image.Point{}.Add(sr.Size())}
	dst := image.NewRGBA(rect)
	draw.Draw(dst, rect, img1, sr.Min, draw.Src)
	return dst
}

func Similarity(imgA, imgB image.Image) (similar bool) {
	// Open photos.
	// imgA, err := images.Open("photoA.jpg")
	// if err != nil {
	// 	panic(err)
	// }
	// imgB, err := images.Open("photoB.jpg")
	// if err != nil {
	// 	panic(err)
	// }

	// Calculate hashes and image sizes.
	hashA, imgSizeA := images.Hash(imgA)
	hashB, imgSizeB := images.Hash(imgB)

	// Image comparison.
	// if images.Similar(hashA, hashB, imgSizeA, imgSizeB) {
	// 	log.Debugf("Images are similar.")
	// 	similar
	// } else {
	// 	log.Debugf("Images are distinct.")
	// }
	similar = images.Similar(hashA, hashB, imgSizeA, imgSizeB)
	log.Debugf("Are Images similar? --> %v", similar)

	return
}

func OpenImg(fname string) image.Image {
	imgA, err := images.Open(fname)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return nil
	}
	return imgA
}

// func Text(img string) {
// 	client := gosseract.NewClient()
// 	defer client.Close()
// 	client.SetImage(img)
// 	text, _ := client.Text()
// 	fmt.Println(text)
// 	// Hello, World!
// }

func Txt(img string) string {
	t, err := tesseract.ReadTextFromFile(img)
	if err != nil {
		fmt.Printf("OCR Error: %v", err.Error())
	}
	return t
}
