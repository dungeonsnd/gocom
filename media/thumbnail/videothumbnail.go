package thumbnail

import (
	"fmt"
	"path"

	"github.com/dungeonsnd/gocom/file/fileutil"
	"github.com/dungeonsnd/gocom/sys/run"
)

// CaptureVideoSnapshot("../../tools/ffmpeg-4.2.2-win64-static/bin/ffmpeg", "IMG_6056.MOV", "IMG_6056.jpg", 200)
func CaptureVideoSnapshot(ffmpegBinFile string, inputFilename string, outFilename string, outputImgWidth uint) bool {

	if fileutil.IsFileNotExist(inputFilename) {
		fmt.Printf("CaptureVideoSnapshot, inputFilename not exist, %v \n", inputFilename)
		return true
	}
	if fileutil.IsFileExist(outFilename) {
		// fmt.Printf("CaptureVideoSnapshot, outFilename exist, %v \n", outFilename)
		return true
	}

	ext := path.Ext(inputFilename)
	if ext != ".mp4" && ext != ".MP4" && ext != ".mov" && ext != ".MOV" {
		return false
	}

	parms := []string{"-ss", "00:00:00", "-i", inputFilename, "-vframes", "1", "-q:v", "2", "-vf", fmt.Sprintf("scale=%d:-1", outputImgWidth), outFilename}
	run.RunExe(ffmpegBinFile, parms)
	return true
}
