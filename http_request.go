package wechatpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

func PostXML(url string, req XmlMap) ([]byte, error) {
	reqBody, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	response, err := http.Post(url, "application/xml; charset=utf-8", bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return ioutil.ReadAll(response.Body)
}

func PostXMLOverTLS(url string, tlsConfig *tls.Config, req XmlMap) ([]byte, error) {
	reqBuffer, err := xml.Marshal(req)
	if err != nil {
		return nil, err
	}

	tlsClient := http.Client{}

	if tlsConfig != nil {
		tlsClient.Transport = &http.Transport{
			TLSClientConfig: tlsConfig,
		}
	}

	resp, err := tlsClient.Post(url, "application/xml; charset=utf-8", bytes.NewBuffer(reqBuffer))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
