// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package sml

import (
	"encoding/xml"
	"fmt"

	"github.com/AlexeyUzhva/gooxml"
)

type CT_ColorScale struct {
	// Conditional Format Value Object
	Cfvo []*CT_Cfvo
	// Color Gradiant Interpolation
	Color []*CT_Color
}

func NewCT_ColorScale() *CT_ColorScale {
	ret := &CT_ColorScale{}
	return ret
}

func (m *CT_ColorScale) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secfvo := xml.StartElement{Name: xml.Name{Local: "ma:cfvo"}}
	for _, c := range m.Cfvo {
		e.EncodeElement(c, secfvo)
	}
	secolor := xml.StartElement{Name: xml.Name{Local: "ma:color"}}
	for _, c := range m.Color {
		e.EncodeElement(c, secolor)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ColorScale) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ColorScale:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cfvo"}:
				tmp := NewCT_Cfvo()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Cfvo = append(m.Cfvo, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "color"}:
				tmp := NewCT_Color()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Color = append(m.Color, tmp)
			default:
				gooxml.Log("skipping unsupported element on CT_ColorScale %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ColorScale
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ColorScale and its children
func (m *CT_ColorScale) Validate() error {
	return m.ValidateWithPath("CT_ColorScale")
}

// ValidateWithPath validates the CT_ColorScale and its children, prefixing error messages with path
func (m *CT_ColorScale) ValidateWithPath(path string) error {
	for i, v := range m.Cfvo {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Cfvo[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Color {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Color[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
