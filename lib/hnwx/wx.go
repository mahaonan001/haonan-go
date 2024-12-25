package wxpay

import (
	"context"
	"log"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/certificates"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

var (
	mchCertificateSerialNumber string
	mchId                      string
	mchAPIv3Key                string
)

// path/to/merchant/apiclient_key.pem
func initClientWx(ctx context.Context, pempath string) (*core.Client, error) {
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(pempath)
	if err != nil {
		log.Fatal("load merchant private key error")
		return nil, err
	}

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchId, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat pay client err:%s", err)
		return nil, err
	}

	// 发送请求，以下载微信支付平台证书为例
	// https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay5_1.shtml
	svc := certificates.CertificatesApiService{Client: client}
	resp, result, err := svc.DownloadCertificates(ctx)
	if err != nil {
		log.Printf("无法下载微信平台支付证书,status:%d,err:%s\n", result.Response.StatusCode, err)
		return nil, err
	}
	log.Printf("status=%d resp=%s\n", result.Response.StatusCode, resp)
	return client, nil
}

// NewJsClinet 初始化微信支付客户端
func NewJsClinet(ctx context.Context, pempath string) (*jsapi.JsapiApiService, error) {
	client, err := initClientWx(ctx, pempath)
	if err != nil {
		return nil, err
	}
	return &jsapi.JsapiApiService{Client: client}, nil
}

// WxPrePay 微信预支付,返回预支付信息
func WxPrePay(svc *jsapi.JsapiApiService, ctx context.Context, payRequest jsapi.PrepayRequest) (*jsapi.PrepayWithRequestPaymentResponse, *core.APIResult, error) {
	resp, result, err := svc.PrepayWithRequestPayment(ctx, payRequest)
	if err != nil {
		return nil, nil, err
	}
	return resp, result, nil
}

// QueryOrder 查询订单,transactionId为微信支付订单号
func QueryOrder(svc *jsapi.JsapiApiService, ctx context.Context, transactionId string) (*payments.Transaction, *core.APIResult, error) {
	resp, result, err := svc.QueryOrderById(ctx,
		jsapi.QueryOrderByIdRequest{
			TransactionId: core.String(transactionId),
			Mchid:         core.String(mchId),
		},
	)
	if err != nil {
		log.Printf("query order err:%s\n", err)
		return nil, nil, err
	}
	return resp, result, nil
}
