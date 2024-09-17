package provider

import (
	"fmt"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type Provider struct {
	AccessKeyId  string
	AccessSecret string
	regionId     string
}

func NewProvider() *Provider {
	return &Provider{regionId: "ap-southeast-1", AccessKeyId: "123", AccessSecret: "abc"}
}

func (a *Provider) SendSms(from, to, message string) error {
	if client, err := sdk.NewClientWithAccessKey(a.regionId, a.AccessKeyId, a.AccessSecret); err != nil {
		return fmt.Errorf("alibaba init " + err.Error())
	} else {
		request := requests.NewCommonRequest()
		request.Method = "POST"
		request.Scheme = "https" // https | http
		request.Domain = "dysmsapi.ap-southeast-1.aliyuncs.com"
		request.Version = "2018-05-01"
		request.ApiName = "SendMessageToGlobe"
		request.QueryParams["To"] = to
		request.QueryParams["RegionId"] = a.regionId
		request.QueryParams["Message"] = message
		_, err := client.ProcessCommonRequest(request)

		return err
	}
}
