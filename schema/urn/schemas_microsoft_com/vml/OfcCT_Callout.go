// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package vml

import (
	"encoding/xml"
	"fmt"

	"github.com/AlexeyUzhva/gooxml/schema/soo/ofc/sharedTypes"
)

type OfcCT_Callout struct {
	OnAttr              sharedTypes.ST_TrueFalse
	TypeAttr            *string
	GapAttr             *string
	AngleAttr           OfcST_Angle
	DropautoAttr        sharedTypes.ST_TrueFalse
	DropAttr            *string
	DistanceAttr        *string
	LengthspecifiedAttr sharedTypes.ST_TrueFalse
	LengthAttr          *string
	AccentbarAttr       sharedTypes.ST_TrueFalse
	TextborderAttr      sharedTypes.ST_TrueFalse
	MinusxAttr          sharedTypes.ST_TrueFalse
	MinusyAttr          sharedTypes.ST_TrueFalse
	ExtAttr             ST_Ext
}

func NewOfcCT_Callout() *OfcCT_Callout {
	ret := &OfcCT_Callout{}
	return ret
}

func (m *OfcCT_Callout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.OnAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.OnAttr.MarshalXMLAttr(xml.Name{Local: "on"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "type"},
			Value: fmt.Sprintf("%v", *m.TypeAttr)})
	}
	if m.GapAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "gap"},
			Value: fmt.Sprintf("%v", *m.GapAttr)})
	}
	if m.AngleAttr != OfcST_AngleUnset {
		attr, err := m.AngleAttr.MarshalXMLAttr(xml.Name{Local: "angle"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DropautoAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.DropautoAttr.MarshalXMLAttr(xml.Name{Local: "dropauto"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DropAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "drop"},
			Value: fmt.Sprintf("%v", *m.DropAttr)})
	}
	if m.DistanceAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "distance"},
			Value: fmt.Sprintf("%v", *m.DistanceAttr)})
	}
	if m.LengthspecifiedAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.LengthspecifiedAttr.MarshalXMLAttr(xml.Name{Local: "lengthspecified"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LengthAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "length"},
			Value: fmt.Sprintf("%v", *m.LengthAttr)})
	}
	if m.AccentbarAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AccentbarAttr.MarshalXMLAttr(xml.Name{Local: "accentbar"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextborderAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.TextborderAttr.MarshalXMLAttr(xml.Name{Local: "textborder"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.MinusxAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.MinusxAttr.MarshalXMLAttr(xml.Name{Local: "minusx"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.MinusyAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.MinusyAttr.MarshalXMLAttr(xml.Name{Local: "minusy"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_Callout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "lengthspecified" {
			m.LengthspecifiedAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "gap" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GapAttr = &parsed
			continue
		}
		if attr.Name.Local == "angle" {
			m.AngleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dropauto" {
			m.DropautoAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "drop" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DropAttr = &parsed
			continue
		}
		if attr.Name.Local == "distance" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DistanceAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TypeAttr = &parsed
			continue
		}
		if attr.Name.Local == "length" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LengthAttr = &parsed
			continue
		}
		if attr.Name.Local == "accentbar" {
			m.AccentbarAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "textborder" {
			m.TextborderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "minusx" {
			m.MinusxAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "minusy" {
			m.MinusyAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcCT_Callout: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcCT_Callout and its children
func (m *OfcCT_Callout) Validate() error {
	return m.ValidateWithPath("OfcCT_Callout")
}

// ValidateWithPath validates the OfcCT_Callout and its children, prefixing error messages with path
func (m *OfcCT_Callout) ValidateWithPath(path string) error {
	if err := m.OnAttr.ValidateWithPath(path + "/OnAttr"); err != nil {
		return err
	}
	if err := m.AngleAttr.ValidateWithPath(path + "/AngleAttr"); err != nil {
		return err
	}
	if err := m.DropautoAttr.ValidateWithPath(path + "/DropautoAttr"); err != nil {
		return err
	}
	if err := m.LengthspecifiedAttr.ValidateWithPath(path + "/LengthspecifiedAttr"); err != nil {
		return err
	}
	if err := m.AccentbarAttr.ValidateWithPath(path + "/AccentbarAttr"); err != nil {
		return err
	}
	if err := m.TextborderAttr.ValidateWithPath(path + "/TextborderAttr"); err != nil {
		return err
	}
	if err := m.MinusxAttr.ValidateWithPath(path + "/MinusxAttr"); err != nil {
		return err
	}
	if err := m.MinusyAttr.ValidateWithPath(path + "/MinusyAttr"); err != nil {
		return err
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
