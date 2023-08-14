package email

import (
	"fmt"

	"github.com/TuringCup/TuringBackend/config"
	"github.com/gin-gonic/gin"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
)

func SendValidCode(email string, code string) (err error) {
	credential := common.NewCredential(
		config.Conf.SES.SecretID,
		config.Conf.SES.SecretKey,
	)
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ses.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := ses.NewClient(credential, "ap-guangzhou", cpf)

	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := ses.NewSendEmailRequest()

	request.FromEmailAddress = common.StringPtr("no-reply@mail.turingcup.tech")
	request.Destination = common.StringPtrs([]string{email})
	request.Template = &ses.Template{
		TemplateID:   common.Uint64Ptr(26519),
		TemplateData: common.StringPtr("{\"valid-code\":\"" + code + "\"}"),
	}
	request.Subject = common.StringPtr("图灵杯验证码")
	response, err := client.SendEmail(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Fprintf(gin.DefaultErrorWriter, "An API error has returned: %s\n", err)
		return
	}
	if err != nil {
		fmt.Fprintln(gin.DefaultErrorWriter, err)
		return
	}
	// 输出json格式的字符串回包
	fmt.Fprintf(gin.DefaultWriter, "%s\n", response.ToJsonString())
	return nil
}
