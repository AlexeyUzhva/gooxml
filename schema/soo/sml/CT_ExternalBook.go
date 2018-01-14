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

type CT_ExternalBook struct {
	IdAttr string
	// Supporting Workbook Sheet Names
	SheetNames *CT_ExternalSheetNames
	// Named Links
	DefinedNames *CT_ExternalDefinedNames
	// Cached Worksheet Data
	SheetDataSet *CT_ExternalSheetDataSet
}

func NewCT_ExternalBook() *CT_ExternalBook {
	ret := &CT_ExternalBook{}
	return ret
}

func (m *CT_ExternalBook) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	if m.SheetNames != nil {
		sesheetNames := xml.StartElement{Name: xml.Name{Local: "ma:sheetNames"}}
		e.EncodeElement(m.SheetNames, sesheetNames)
	}
	if m.DefinedNames != nil {
		sedefinedNames := xml.StartElement{Name: xml.Name{Local: "ma:definedNames"}}
		e.EncodeElement(m.DefinedNames, sedefinedNames)
	}
	if m.SheetDataSet != nil {
		sesheetDataSet := xml.StartElement{Name: xml.Name{Local: "ma:sheetDataSet"}}
		e.EncodeElement(m.SheetDataSet, sesheetDataSet)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ExternalBook) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
	}
lCT_ExternalBook:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetNames"}:
				m.SheetNames = NewCT_ExternalSheetNames()
				if err := d.DecodeElement(m.SheetNames, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "definedNames"}:
				m.DefinedNames = NewCT_ExternalDefinedNames()
				if err := d.DecodeElement(m.DefinedNames, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "sheetDataSet"}:
				m.SheetDataSet = NewCT_ExternalSheetDataSet()
				if err := d.DecodeElement(m.SheetDataSet, &el); err != nil {
					return err
				}
			default:
				gooxml.Log("skipping unsupported element on CT_ExternalBook %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ExternalBook
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ExternalBook and its children
func (m *CT_ExternalBook) Validate() error {
	return m.ValidateWithPath("CT_ExternalBook")
}

// ValidateWithPath validates the CT_ExternalBook and its children, prefixing error messages with path
func (m *CT_ExternalBook) ValidateWithPath(path string) error {
	if m.SheetNames != nil {
		if err := m.SheetNames.ValidateWithPath(path + "/SheetNames"); err != nil {
			return err
		}
	}
	if m.DefinedNames != nil {
		if err := m.DefinedNames.ValidateWithPath(path + "/DefinedNames"); err != nil {
			return err
		}
	}
	if m.SheetDataSet != nil {
		if err := m.SheetDataSet.ValidateWithPath(path + "/SheetDataSet"); err != nil {
			return err
		}
	}
	return nil
}
