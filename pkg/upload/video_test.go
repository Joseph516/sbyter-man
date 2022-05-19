package upload

import (
	"testing"
)

func TestExactCoverFromVideo(t *testing.T) {
	pathVideo := "/Users/joe/Desktop/developer/sbyter-man/storage/uploads/17/766a527bbe75cbe6d15f3b0218d670d6.mp4"
	pathImg := "/Users/joe/Desktop/developer/sbyter-man/storage/uploads/17/766a527bbe75cbe6d15f3b0218d670d6.png"
	err := ExactCoverFromVideo(pathVideo, pathImg)
	if err != nil {
		t.Error(err)
	}
}
