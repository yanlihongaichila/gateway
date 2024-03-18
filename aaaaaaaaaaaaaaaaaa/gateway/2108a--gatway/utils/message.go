package utils

import (
	dysmsapi "github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/spf13/viper"
)

func sendMessage(phone, signName, code, codePrm string) error {
	client, err := dysmsapi.NewClientWithAccessKey(
		viper.GetString("aliyun.regionId"),
		viper.GetString("aliyun.regionId"),
		viper.GetString("aliyun.accessKeySecret"),
	)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = signName
	request.TemplateCode = code
	request.TemplateParam = codePrm

	_, err = client.SendSms(request)
	if err != nil {
		return err
	}
	return nil
}
