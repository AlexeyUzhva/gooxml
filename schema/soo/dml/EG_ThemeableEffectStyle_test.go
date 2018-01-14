// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package dml_test

import (
	"encoding/xml"
	"testing"

	"github.com/AlexeyUzhva/gooxml/schema/soo/dml"
)

func TestEG_ThemeableEffectStyleConstructor(t *testing.T) {
	v := dml.NewEG_ThemeableEffectStyle()
	if v == nil {
		t.Errorf("dml.NewEG_ThemeableEffectStyle must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_ThemeableEffectStyle should validate: %s", err)
	}
}

func TestEG_ThemeableEffectStyleMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_ThemeableEffectStyle()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_ThemeableEffectStyle()
	xml.Unmarshal(buf, v2)
}
