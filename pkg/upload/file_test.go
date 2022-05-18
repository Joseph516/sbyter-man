package upload

import (
	"fmt"
	"testing"
)

func TestGetFileNameWithTime(t *testing.T) {
	filename := "sadsaf.mp4"
	filenameEncode1 := GetFileNameWithTime(filename)
	fmt.Println("filenameEncode1:", filenameEncode1)

	// time.Sleep(1 * time.Second)

	filenameEncode2 := GetFileNameWithTime(filename)
	fmt.Println("filenameEncode2:", filenameEncode2)

	if filenameEncode1 == filenameEncode2 {
		t.Error("加密失败")
	}

}
