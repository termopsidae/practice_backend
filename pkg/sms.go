package pkg

import (
	"errors"
	"fmt"
	"github.com/vonage/vonage-go-sdk"
	client "github.com/yunpian/yunpian-go-sdk/sdk"
	"paractice/routing/types"
)

func SendVonageCodeMsg(msgParam types.SengMsg) error {
	auth := vonage.CreateAuthFromKeySecret("0f32143c", "LF9C0YzY1kgXHVQY")
	smsClient := vonage.NewSMSClient(auth)
	opts := vonage.SMSOpts{
		Type: "unicode",
	}
	response, errResp, err := smsClient.Send("【Pooluo】", msgParam.Area+msgParam.Phone, "验证码: "+msgParam.Msg, opts)
	if err != nil {
		return err
	}
	if response.Messages[0].Status == "0" {
		fmt.Println("Account Balance: " + response.Messages[0].RemainingBalance)
	} else {
		fmt.Println("SendMessage Error code " + errResp.Messages[0].Status + ": " + errResp.Messages[0].ErrorText)
		return errors.New(errResp.Messages[0].ErrorText)
	}
	return nil
}

func SendCheckCodeMessage(msgParam types.SengMsg) error {
	// 发送短信
	ypClient := client.New("ec557d72a53ef29f0aa0c39e79d59814")
	param := client.NewParam(2)
	param[client.MOBILE] = msgParam.Phone
	param[client.TEXT] = "【 TreasureBox】您的手机验证码是" + msgParam.Msg + "。本条信息无需回复"
	r := ypClient.Sms().SingleSend(param)
	if r.Code != 0 {
		return errors.New(r.Msg)
	}
	return nil
}
