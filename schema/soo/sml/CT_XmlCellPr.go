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
	"strconv"

	"github.com/AlexeyUzhva/gooxml"
)

type CT_XmlCellPr struct {
	// Table Field Id
	IdAttr uint32
	// Unique Table Name
	UniqueNameAttr *string
	// Column XML Properties
	XmlPr *CT_XmlPr
	// Future Feature Data Storage Area
	ExtLst *CT_ExtensionList
}

func NewCT_XmlCellPr() *CT_XmlCellPr {
	ret := &CT_XmlCellPr{}
	ret.XmlPr = NewCT_XmlPr()
	return ret
}

func (m *CT_XmlCellPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	if m.UniqueNameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueName"},
			Value: fmt.Sprintf("%v", *m.UniqueNameAttr)})
	}
	e.EncodeToken(start)
	sexmlPr := xml.StartElement{Name: xml.Name{Local: "ma:xmlPr"}}
	e.EncodeElement(m.XmlPr, sexmlPr)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_XmlCellPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.XmlPr = NewCT_XmlPr()
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "uniqueName" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueNameAttr = &parsed
			continue
		}
	}
lCT_XmlCellPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "xmlPr"}:
				if err := d.DecodeElement(m.XmlPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_XmlCellPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_XmlCellPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_XmlCellPr and its children
func (m *CT_XmlCellPr) Validate() error {
	return m.ValidateWithPath("CT_XmlCellPr")
}

// ValidateWithPath validates the CT_XmlCellPr and its children, prefixing error messages with path
func (m *CT_XmlCellPr) ValidateWithPath(path string) error {
	if err := m.XmlPr.ValidateWithPath(path + "/XmlPr"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
