package tag

import (
	"fmt"
	"testing"
)

func TestExtractTagFromText(t *testing.T) {
	text := "腾讯云自然语言处理接口是基于腾讯云计算强大的自然语言处理能力，提供自然语言处理、语音识别、文本翻译、文本相似度、文本分类、知识图谱等功能。"

	tags, err := ExtractTagFromText(text)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(tags)
}
