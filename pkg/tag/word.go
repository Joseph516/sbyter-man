package tag

import (
	"fmt"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	nlp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/nlp/v20190408"
)

func ExtractTagFromText(text string) (tags []string, err error) {
	credential := common.NewCredential(
		"AKID1vfIftzsZEinkYEuNm177yfNw3xiP2Ws",
		"scAXmrneR2Rddr5QdptkOo3A82pSwMl1",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "nlp.tencentcloudapi.com"
	client, _ := nlp.NewClient(credential, "ap-guangzhou", cpf)

	request := nlp.NewKeywordsExtractionRequest()

	// 非空处理
	if text == "" {
		text = "视频"
	}
	request.Text = common.StringPtr(text)

	response, err := client.KeywordsExtraction(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		err = fmt.Errorf("An API error has returned: %s", err)
		return
	}
	if err != nil {
		return
	}

	// 获取关键字和分数
	tags = make([]string, 0, len(response.Response.Keywords))
	for _, value := range response.Response.Keywords {
		tags = append(tags, *value.Word)
	}

	return
}
