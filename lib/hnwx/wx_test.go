package wxpay

import (
	"context"
	"testing"
	"time"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
)

func TestPay(t *testing.T) {
	ctx := context.Background()
	client, err := NewJsClinet(ctx, "aa.pem")
	if err != nil {
		t.Logf("NewJsClinet err:%s", err)
		return
	}
	response, result, err := WxPrePay(client, ctx, jsapi.PrepayRequest{
		Appid:       core.String("wx1234567890"),
		Mchid:       core.String("1234567890"),
		Description: core.String("test"),
		OutTradeNo:  core.String("1234567890"),
		NotifyUrl:   core.String("https://www.example.com"),
		Amount: &jsapi.Amount{
			Total: core.Int64(1),
		},
		Attach:     core.String("test"),
		TimeExpire: core.Time(time.Now()),
	})
	if err != nil {
		t.Logf("WxPrePay err:%s", err)
		return
	}
	t.Logf("response:%v,result:%v", response, result)
}
