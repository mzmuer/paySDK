package wechatpay

import (
	"log"
	"testing"
)

func TestPay_UnifiedOrder(t *testing.T) {
	p := NewPay("wx6bb31df364e230a4", "1521866491", "c36f10fd0ff59c3bcce088d7a7a6c410", false)
	p.SetSignType(SignTypeHMACSHA256)

	req := XmlMap{
		"body":             "body",
		"out_trade_no":     "111",
		"total_fee":        "1",
		"spbill_create_ip": "127.0.0.1",
		"notify_url":       "http://mzmuer.imwork.net/v1/notify/wechat/pay",
		"openid":           "ox-7H5XQGl4tds6JLBeWX1S2k1Ng",
		"trade_type":       "JSAPI",
	}
	resp, err := p.UnifiedOrder(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

func TestPay_PromotionTransfers(t *testing.T) {
	p := NewPay("wx816dc6e826dcc6b5", "1434654302", "rP88g8svXp5PLX1jZLPz1ciX38CIeETG", false)
	err := p.SetTLS("tls/apiclient_cert.online.pem", "tls/apiclient_key.online.pem")
	if err != nil {
		t.Fatal(err)
	}

	req := XmlMap{
		"partner_trade_no": "11111x1",
		"check_name":       "NO_CHECK",
		"amount":           "30",
		"spbill_create_ip": "127.0.0.1",
		"desc":             "测试付款",
		"openid":           "oyPYb5OGTvuZ-3QtTwGqk8dp1Sgo",
	}

	resp, err := p.PromotionTransfers(req)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(resp)
}

// 获取付款信息
func TestPay_Gettransferinfo(t *testing.T) {
	p := NewPay("wx816dc6e826dcc6b5", "1434654302", "rP88g8svXp5PLX1jZLPz1ciX38CIeETG", false)
	err := p.SetTLS("tls/apiclient_cert.online.pem", "tls/apiclient_key.online.pem")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := p.Gettransferinfo("11111x1")
	if err != nil {
		t.Fatal(err, resp)
	}

	log.Println("--x ", resp)
	t.Log(resp)
}
