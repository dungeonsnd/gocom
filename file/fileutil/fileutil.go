package fileutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var mExtsImageFile map[string]string
var mExtsVideoFile map[string]string

func init() {
	mExtsImageFile = map[string]string{".png": "", ".PNG": "", ".jpg": "", ".JPG": "", ".jpeg": "", ".JPEG": "",
		".heic": "", ".HEIC": "", ".gif": "", ".GIF": "",
	}
	mExtsVideoFile = map[string]string{".MOV": "", ".mov": "", ".MP4": "", ".mp4": ""}
}

func IsMediaFile(filename string) bool {
	ext := path.Ext(filename)
	return IsImage(ext) || IsVideo(ext)
}

func IsImage(filename string) bool {
	ext := path.Ext(filename)
	_, exist := mExtsImageFile[ext]
	return exist
}

func IsVideo(filename string) bool {
	ext := path.Ext(filename)
	_, exist := mExtsVideoFile[ext]
	return exist
}

func GetFilenameOnly(fullfilename string) string {
	return strings.TrimSuffix(fullfilename, path.Ext(fullfilename))
}

func FileNameToJpgFileName(onlyFilename string) string {
	name := GetFilenameOnly(onlyFilename)
	return name + ".JPG"
}

// 如果是单个文件则返回 true.
func IsFile(f string) (isFile bool, err error) {
	fi, err := os.Stat(f)
	if err != nil {
		return
	}
	isFile = !fi.IsDir()
	return
}

func IsFileExist(f string) bool {
	return !IsFileNotExist(f)
}
func IsFileNotExist(f string) bool {
	_, err := os.Stat(f)
	return os.IsNotExist(err)
}

func AddPathSepIfNeed(path string) (newPath string) {
	newPath = path
	if len(path) > 0 {
		if path[len(path)-1:] != "/" {
			newPath += "/"
		}
	} else {
		newPath += "/"
	}
	return
}

//递归创建文件夹
func CreateDirRecursive(filePath string) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) { // not exist
		err := os.MkdirAll(filePath, os.ModePerm)
		return err
	}
	return nil
}

func WriteToFile(filename string, content []byte, truncateIfExist bool) error {
	flag := os.O_RDWR | os.O_CREATE | os.O_EXCL
	if truncateIfExist {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}
	fileObj, err := os.OpenFile(filename, flag, 0644)
	if err != nil {
		return err
	}
	defer fileObj.Close()

	n, err := fileObj.Write(content)
	if err != nil {
		return err
	}
	if n != len(content) {
		return errors.New("written length error")
	}
	return nil
}

func ReadFromFile(filename string) (error, []byte) {
	fileObj, err := os.Open(filename)
	if err != nil {
		fmt.Printf("ReadFromFile, Open err=%v\n", err)
		return err, nil
	}
	defer fileObj.Close()

	content, err := ioutil.ReadAll(fileObj)
	if err != nil {
		return err, nil
	}

	return nil, content
}
