package wechatpay

import (
	"encoding/xml"
	"io"
)

type xmlMap map[string]string

type xmlEntry struct {
	XMLName xml.Name
	Value   string `xml:",chardata"`
}

func (m *xmlMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	*m = xmlMap{}
	for {
		var e xmlEntry

		if err := d.Decode(&e); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		(*m)[e.XMLName.Local] = e.Value
	}
	return nil
}

func (m xmlMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if len(m) == 0 {
		return nil
	}

	start.Name.Local = "xml"

	if err := e.EncodeToken(start); err != nil {
		return err
	}

	for k, v := range m {
		if err := e.Encode(xmlEntry{XMLName: xml.Name{Local: k}, Value: v}); err != nil {
			return err
		}
	}

	return e.EncodeToken(start.End())
}
