package consumer

import (
	"douyin_service/global"
	"testing"
)

func TestConsumeEmail(t *testing.T) {
	var err error
	global.Consumer, err = NewConsumer()
	if err != nil {
		t.Fatal(err)
	}
	ConsumeEmail()
}
