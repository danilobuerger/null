// Copyright (c) 2018 Danilo Bürger <info@danilobuerger.de>

package null

import "encoding/xml"

const (
	jsonNull     = "null"
	xsiNamespace = "http://www.w3.org/2001/XMLSchema-instance"
)

var xsiNilAttr = xml.Attr{
	Name:  xml.Name{Local: "xsi:nil"},
	Value: "true",
}

func isXsiNilAttr(attr xml.Attr) bool {
	return (attr.Name.Space == "xsi" || attr.Name.Space == xsiNamespace) &&
		attr.Name.Local == "nil" && attr.Value == "true"
}
