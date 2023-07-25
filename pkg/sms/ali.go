package sms

import (
	"encoding/json"
	"strings"

	"github.com/denverdino/aliyungo/sms"
	"github.com/serenefiregroup/ffa_server/pkg/config"
	"github.com/serenefiregroup/ffa_server/pkg/errors"
	"github.com/serenefiregroup/ffa_server/pkg/log"
)

type Context struct {
	Phone      string
	TemplateID string
	Params     map[string]interface{}
}

var ali *AliyunSMS

func LoadConfig() error {
	ali = new(AliyunSMS)
	ali.dysmsAccessKey = config.String("aliyun_sms_access_key", "")
	ali.dysmsSecret = config.String("aliyun_sms_secret", "")
	ali.productName = config.ProductNameZH
	return nil
}

type AliyunSMS struct {
	productName    string
	dysmsAccessKey string
	dysmsSecret    string
}

func (a *AliyunSMS) postSMS(ctx Context) (bool, error) {
	if a.dysmsAccessKey == "" || a.dysmsSecret == "" {
		return false, nil
	}

	phone := strings.TrimPrefix(ctx.Phone, "+86")
	params := ctx.Params
	paramJSON, _ := json.Marshal(params)

	args := &sms.SendSmsArgs{
		TemplateCode:  ctx.TemplateID,
		TemplateParam: string(paramJSON),
		PhoneNumbers:  phone,
		SignName:      a.productName,
	}

	cli := sms.NewDYSmsClient(a.dysmsAccessKey, a.dysmsSecret)
	resp, err := cli.SendSms(args)
	if err != nil {
		values := map[string]interface{}{
			"code":    resp.Code,
			"message": resp.Message,
		}
		return false, errors.WithValues(errors.SmsError(err), values)
	}
	if resp.Code != "OK" {
		values := map[string]interface{}{
			"code":         resp.Code,
			"message":      resp.Message,
			"TemplateCode": ctx.TemplateID,
		}
		err = errors.WithValues(errors.New("SMS"), values)
		log.Warn("send sms error: %+v", err)
		return false, err
	}
	return true, nil
}

func PostSignUpVerifySMS(phone string, code string) (bool, error) {
	ctx := Context{
		Phone:      phone,
		TemplateID: "SMS_461890174",
		Params: map[string]interface{}{
			"code": code,
		},
	}
	return ali.postSMS(ctx)
}
