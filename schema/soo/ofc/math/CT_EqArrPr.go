// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math

import (
	"encoding/xml"

	"github.com/AlexeyUzhva/gooxml"
)

type CT_EqArrPr struct {
	BaseJc  *CT_YAlign
	MaxDist *CT_OnOff
	ObjDist *CT_OnOff
	RSpRule *CT_SpacingRule
	RSp     *CT_UnSignedInteger
	CtrlPr  *CT_CtrlPr
}

func NewCT_EqArrPr() *CT_EqArrPr {
	ret := &CT_EqArrPr{}
	return ret
}

func (m *CT_EqArrPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.BaseJc != nil {
		sebaseJc := xml.StartElement{Name: xml.Name{Local: "m:baseJc"}}
		e.EncodeElement(m.BaseJc, sebaseJc)
	}
	if m.MaxDist != nil {
		semaxDist := xml.StartElement{Name: xml.Name{Local: "m:maxDist"}}
		e.EncodeElement(m.MaxDist, semaxDist)
	}
	if m.ObjDist != nil {
		seobjDist := xml.StartElement{Name: xml.Name{Local: "m:objDist"}}
		e.EncodeElement(m.ObjDist, seobjDist)
	}
	if m.RSpRule != nil {
		serSpRule := xml.StartElement{Name: xml.Name{Local: "m:rSpRule"}}
		e.EncodeElement(m.RSpRule, serSpRule)
	}
	if m.RSp != nil {
		serSp := xml.StartElement{Name: xml.Name{Local: "m:rSp"}}
		e.EncodeElement(m.RSp, serSp)
	}
	if m.CtrlPr != nil {
		sectrlPr := xml.StartElement{Name: xml.Name{Local: "m:ctrlPr"}}
		e.EncodeElement(m.CtrlPr, sectrlPr)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EqArrPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EqArrPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "baseJc"}:
				m.BaseJc = NewCT_YAlign()
				if err := d.DecodeElement(m.BaseJc, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "maxDist"}:
				m.MaxDist = NewCT_OnOff()
				if err := d.DecodeElement(m.MaxDist, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "objDist"}:
				m.ObjDist = NewCT_OnOff()
				if err := d.DecodeElement(m.ObjDist, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rSpRule"}:
				m.RSpRule = NewCT_SpacingRule()
				if err := d.DecodeElement(m.RSpRule, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "rSp"}:
				m.RSp = NewCT_UnSignedInteger()
				if err := d.DecodeElement(m.RSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/math", Local: "ctrlPr"}:
				m.CtrlPr = NewCT_CtrlPr()
				if err := d.DecodeElement(m.CtrlPr, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_EqArrPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EqArrPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EqArrPr and its children
func (m *CT_EqArrPr) Validate() error {
	return m.ValidateWithPath("CT_EqArrPr")
}

// ValidateWithPath validates the CT_EqArrPr and its children, prefixing error messages with path
func (m *CT_EqArrPr) ValidateWithPath(path string) error {
	if m.BaseJc != nil {
		if err := m.BaseJc.ValidateWithPath(path + "/BaseJc"); err != nil {
			return err
		}
	}
	if m.MaxDist != nil {
		if err := m.MaxDist.ValidateWithPath(path + "/MaxDist"); err != nil {
			return err
		}
	}
	if m.ObjDist != nil {
		if err := m.ObjDist.ValidateWithPath(path + "/ObjDist"); err != nil {
			return err
		}
	}
	if m.RSpRule != nil {
		if err := m.RSpRule.ValidateWithPath(path + "/RSpRule"); err != nil {
			return err
		}
	}
	if m.RSp != nil {
		if err := m.RSp.ValidateWithPath(path + "/RSp"); err != nil {
			return err
		}
	}
	if m.CtrlPr != nil {
		if err := m.CtrlPr.ValidateWithPath(path + "/CtrlPr"); err != nil {
			return err
		}
	}
	return nil
}
