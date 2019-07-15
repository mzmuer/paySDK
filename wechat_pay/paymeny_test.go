package wechat_pay

import (
	"testing"
)

func TestPay_UnifiedOrder(t *testing.T) {
	p, err := NewPay(
		"x",
		"x",
		"x",
		"x",
		"",
		false)
	if err != nil {
		t.Fatal(err)
	}

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
