package wechat_pay

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"paysdk/utils"
)

type Pay struct {
	AppId     string
	MchId     string
	Key       string
	tlsConfig *tls.Config
	isSandBox bool
}

func NewPay() *Pay {
	return &Pay{}
}

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

	// 补充字段
	req["appid"] = p.AppId
	req["mch_id"] = p.MchId
	req["sign"] = utils.MapSignMD5(req, p.Key)

	resp, err := utils.PostXML(DomainApi+uri, p.fillRequestData(req))
	if err != nil {
		return nil, err
	}

	result := XmlMap{}
	if err = xml.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if err = p.verifyResponse(result); err != nil {
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

	// 请求退款
	resp, err := utils.PostXMLOverTLS(uri, p.tlsConfig, p.fillRequestData(req))
	if err != nil {
		return nil, err
	}

	result := XmlMap{}
	if err = xml.Unmarshal(resp, &result); err != nil {
		return nil, err
	}

	if err = p.verifyResponse(result); err != nil {
		return nil, err
	}

	return result, nil
}

func (p *Pay) SignVerify(m XmlMap) bool {
	sign := m["sign"]
	delete(m, "sign")
	return utils.MapSignMD5(m, p.Key) != sign
}

// --
func (p *Pay) verifyResponse(res XmlMap) error {
	if res["return_code"] != Success {
		return fmt.Errorf(res["return_msg"])
	}

	if res["result_code"] != Success {
		return fmt.Errorf(res["result_code"] + "_" + res["err_code_des"])
	}

	if p.SignVerify(res) {
		return fmt.Errorf("CreateOrder sign not match[#%+v#]", res)
	}

	return nil
}

func (p *Pay) fillRequestData(m XmlMap) XmlMap {
	m["appid"] = p.AppId
	m["mch_id"] = p.MchId
	m["sign"] = utils.MapSignMD5(m, p.Key)
	m["nonce_str"] = utils.RandomString(24)

	return m
}
