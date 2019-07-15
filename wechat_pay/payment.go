package wechat_pay

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"github.com/mzmuer/paysdk"
)

type Pay struct {
	AppId     string
	MchId     string
	Key       string
	SignType  string
	tlsConfig *tls.Config
	isSandBox bool
}

func NewPay(appId, mchId, key, certFile, certKeyFile string, isSandBox bool) (*Pay, error) {
	cert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
	if err != nil {
		return nil, err
	}

	return &Pay{
		AppId:    appId,
		MchId:    mchId,
		Key:      key,
		SignType: paysdk.SignTypeMD5,
		tlsConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
		isSandBox: isSandBox,
	}, nil
}

// config
func (p *Pay) SetSignType(signType string) {
	p.SignType = signType
}

// -------------------------------------------------------------
// 创建支付订单
func (p *Pay) UnifiedOrder(req XmlMap) (XmlMap, error) {
	if req["body"] == "" ||
		req["out_trade_no"] == "" ||
		req["total_fee"] == "" ||
		req["spbill_create_ip"] == "" ||
		req["notify_url"] == "" ||
		req["trade_type"] == "" {

		return nil, fmt.Errorf("缺少必传参数")
	}

	var uri string
	if p.isSandBox {
		uri = SandboxUnifiedorderUrlSuffix
	} else {
		uri = UnifiedorderUrlSuffix
	}

	// 填充字段
	req, err := p.fillRequestData(req)
	if err != nil {
		return nil, err
	}

	// 发起请求
	resp, err := paysdk.PostXML(DomainApi+uri, req)
	if err != nil {
		return nil, err
	}

	result := XmlMap{}
	if err = xml.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if err = p.VerifyResponse(result); err != nil {
		return nil, err
	}

	return result, nil
}

// 退款请求
func (p *Pay) Refund(req XmlMap) (XmlMap, error) {
	if (req["transaction_id"] == "" && req["out_trade_no"] == "") ||
		req["total_fee"] == "" ||
		req["refund_fee	"] == "" {

		return nil, fmt.Errorf("缺少必传参数")
	}

	var uri string
	if p.isSandBox {
		uri = SandboxRefundUrlSuffix
	} else {
		uri = RefundUrlSuffix
	}

	// 填充字段
	req, err := p.fillRequestData(req)
	if err != nil {
		return nil, err
	}

	// 请求退款
	resp, err := paysdk.PostXMLOverTLS(uri, p.tlsConfig, req)
	if err != nil {
		return nil, err
	}

	result := XmlMap{}
	if err = xml.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if err = p.VerifyResponse(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pay) SignVerify(m XmlMap) (bool, error) {
	sign := m["sign"]
	delete(m, "sign")
	sign2, err := paysdk.GenerateMapSign(m, p.SignType, p.Key)
	return sign2 == sign, err
}

// --
func (p *Pay) VerifyResponse(res XmlMap) error {
	if res["return_code"] != Success {
		return fmt.Errorf(res["return_msg"])
	}

	if res["result_code"] != Success {
		return fmt.Errorf(res["result_code"] + "_" + res["err_code_des"])
	}

	match, err := p.SignVerify(res)
	if err != nil {
		return err
	}

	if !match {
		return fmt.Errorf("sign not match[#%+v#]", res)
	}

	return nil
}

// ==========================================================
func (p *Pay) fillRequestData(m XmlMap) (XmlMap, error) {
	m["appid"] = p.AppId
	m["mch_id"] = p.MchId
	m["sign_type"] = p.SignType
	m["nonce_str"] = paysdk.RandomString(24)

	sign, err := paysdk.GenerateMapSign(m, m["sign_type"], p.Key)
	if err != nil {
		return nil, err
	}

	m["sign"] = sign

	return m, nil
}
