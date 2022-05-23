package util

import "testing"

func TestRandomSign(t *testing.T) {
	str, err := RandomSign()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(str)
}

func TestRandomAvatar(t *testing.T) {
	url := RandomAvatar("123")
	t.Log(url)
}

func TestRandomBackground(t *testing.T) {
	img, err := RandomBackground()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(img)
}
