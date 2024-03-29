package upload

import (
	"douyin_service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

type FileType int

const TypeVideo FileType = iota

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}

// GetFileNameWithTime加入时间戳
func GetFileNameWithTime(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName + time.Now().String())
	return fileName + ext
}

func GetFileExt(name string) string {
	// fmt.Println(name)
	return path.Ext(name)
}

func GetFilenameWithoutExt(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	return fileName
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckContainExt(t FileType, name string, uploadVideoAllowExts []string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeVideo:
		for _, allowExt := range uploadVideoAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File, uploadVideoMaxSize int) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeVideo:
		if size >= uploadVideoMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CheckMaxSizeByHeader(t FileType, size int, uploadVideoMaxSize int) bool {
	switch t {
	case TypeVideo:
		if size >= uploadVideoMaxSize*1024*1024 {
			return true
		}
	}
	return false
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

// CopyFile将src位置的文件夹拷贝至dst
func CopyFile(src, dst string) error {
	out, err := os.Create(dst)
	if err != nil {
		return err
	}

	defer out.Close()

	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}

	return nil
}
