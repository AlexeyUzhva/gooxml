// Copyright 2017 Baliance. All rights reserved.
//
// DO NOT EDIT: generated by gooxml ECMA-376 generator
//
// Use of this source code is governed by the terms of the Affero GNU General
// Public License version 3.0 as published by the Free Software Foundation and
// appearing in the file LICENSE included in the packaging of this file. A
// commercial license can be purchased by contacting sales@baliance.com.

package math_test

import (
	"encoding/xml"
	"testing"

	"github.com/AlexeyUzhva/gooxml/schema/soo/ofc/math"
)

func TestEG_OMathElementsConstructor(t *testing.T) {
	v := math.NewEG_OMathElements()
	if v == nil {
		t.Errorf("math.NewEG_OMathElements must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.EG_OMathElements should validate: %s", err)
	}
}

func TestEG_OMathElementsMarshalUnmarshal(t *testing.T) {
	v := math.NewEG_OMathElements()
	buf, _ := xml.Marshal(v)
	v2 := math.NewEG_OMathElements()
	xml.Unmarshal(buf, v2)
}
