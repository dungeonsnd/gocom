package thumbnail

import (
	"fmt"
	"log"
	"os"
	"path"

	"image/jpeg"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/dungeonsnd/gocom/file/fileutil"
	"github.com/nfnt/resize"
)

func ResizeImage(inputFilename string, outFilename string, width uint) bool {

	if fileutil.IsFileNotExist(inputFilename) {
		fmt.Printf("ResizeImage, inputFilename not exist, %v \n", inputFilename)
		return true
	}
	if fileutil.IsFileExist(outFilename) {
		// fmt.Printf("ResizeImage, outFilename exist, %v \n", outFilename)
		return true
	}

	ext := path.Ext(inputFilename)
	if ext == ".png" || ext == ".PNG" {
		resizePNGByBild(inputFilename, outFilename, width)
	} else if ext == ".jpg" || ext == ".JPG" {
		resizeJPGByNfnt(inputFilename, outFilename, width)
	} else {
		return false
	}
	return true
}

func resizeJPGByNfnt(inputFilename string, outFilename string, width uint) {
	// fmt.Printf("resizeJPGByNfnt, inputFilename=%v, outFilename=%v\n", inputFilename, outFilename)
	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(width, 0, img, resize.Lanczos3)

	out, err := os.Create(outFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}

func resizePNGByBild(inputFilename string, outfilename string, width uint) {
	// fmt.Printf("resizePNGByBild, inputFilename=%v, outfilename=%v\n", inputFilename, outfilename)
	img, err := imgio.Open(inputFilename)
	if err != nil {
		fmt.Println(err)
		return
	}
	w := img.Bounds().Max.X
	h := img.Bounds().Max.Y
	newH := float64(width) * (float64(h) / float64(w))
	resized := transform.Resize(img, int(width), int(newH), transform.Lanczos)
	if err := imgio.Save(outfilename, resized, imgio.PNGEncoder()); err != nil {
		fmt.Println(err)
		return
	}
}
