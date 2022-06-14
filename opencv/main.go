package main

import (
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.VideoCaptureDevice(0)
	jendela := gocv.NewWindow("jendela")
	gambar := gocv.NewMat()

	for {
		webcam.Read(&gambar)
		jendela.IMShow(gambar)
		jendela.WaitKey(1)
	}
}
