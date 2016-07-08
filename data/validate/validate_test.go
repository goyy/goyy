// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate_test

import (
	"testing"

	"gopkg.in/goyy/goyy.v0/data/validate"
)

func TestValidateRequired(t *testing.T) {
	in := "ABc123"
	if err := validate.Required(in); err != nil {
		t.Errorf(`validate.Required(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "    "
	if err := validate.Required(in); err == nil {
		t.Errorf(`validate.Required(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateMin(t *testing.T) {
	min := 20
	in := "30"
	if err := validate.Min(in, min); err != nil {
		t.Errorf(`validate.Min(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "20"
	if err := validate.Min(in, min); err != nil {
		t.Errorf(`validate.Min(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "10"
	if err := validate.Min(in, min); err == nil {
		t.Errorf(`validate.Min(%s, %d) = %v, want %s`, in, min, "nil", "not nil")
	}
}

func TestValidateMax(t *testing.T) {
	max := 20
	in := "10"
	if err := validate.Max(in, max); err != nil {
		t.Errorf(`validate.Max(%s, %d) = %v, want %s`, in, max, err.Error(), "nil")
	}
	in = "20"
	if err := validate.Max(in, max); err != nil {
		t.Errorf(`validate.Max(%s, %d) = %v, want %s`, in, max, err.Error(), "nil")
	}
	in = "30"
	if err := validate.Max(in, max); err == nil {
		t.Errorf(`validate.Max(%s, %d) = %v, want %s`, in, max, "nil", "not nil")
	}
}

func TestValidateRange(t *testing.T) {
	min := 20
	max := 30
	in := "25"
	if err := validate.Range(in, min, max); err != nil {
		t.Errorf(`validate.Range(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "20"
	if err := validate.Range(in, min, max); err != nil {
		t.Errorf(`validate.Range(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "30"
	if err := validate.Range(in, min, max); err != nil {
		t.Errorf(`validate.Range(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "35"
	if err := validate.Range(in, min, max); err == nil {
		t.Errorf(`validate.Range(%s, %d, %d) = %v, want %s`, in, min, max, "nil", "not nil")
	}
}

func TestValidateMinlen(t *testing.T) {
	min := 5
	in := "ABc123"
	if err := validate.Minlen(in, min); err != nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "ABc12"
	if err := validate.Minlen(in, min); err != nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "ABc一二三"
	if err := validate.Minlen(in, min); err != nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "ABc一二"
	if err := validate.Minlen(in, min); err != nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, err.Error(), "nil")
	}
	in = "ABc"
	if err := validate.Minlen(in, min); err == nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, "nil", "not nil")
	}
	in = "ABc1"
	if err := validate.Minlen(in, min); err == nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, "nil", "not nil")
	}
	in = "ABc一"
	if err := validate.Minlen(in, min); err == nil {
		t.Errorf(`validate.Minlen(%s, %d) = %v, want %s`, in, min, "nil", "not nil")
	}
}

func TestValidateMaxlen(t *testing.T) {
	max := 5
	in := "ABc"
	if err := validate.Maxlen(in, max); err != nil {
		t.Errorf(`validate.Maxlen(%s, %d) = %v, want %s`, in, max, err.Error(), "nil")
	}
	in = "ABc12"
	if err := validate.Maxlen(in, max); err != nil {
		t.Errorf(`validate.Maxlen(%s, %d) = %v, want %s`, in, max, err.Error(), "nil")
	}
	in = "ABc一二"
	if err := validate.Maxlen(in, max); err != nil {
		t.Errorf(`validate.Maxlen(%s, %d) = %v, want %s`, in, max, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Maxlen(in, max); err == nil {
		t.Errorf(`validate.Maxlen(%s, %d) = %v, want %s`, in, max, "nil", "not nil")
	}
	in = "ABc一二三"
	if err := validate.Maxlen(in, max); err == nil {
		t.Errorf(`validate.Maxlen(%s, %d) = %v, want %s`, in, max, "nil", "not nil")
	}
}

func TestValidateRangelen(t *testing.T) {
	min := 2
	max := 5
	in := "ABc"
	if err := validate.Rangelen(in, min, max); err != nil {
		t.Errorf(`validate.Rangelen(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "AB"
	if err := validate.Rangelen(in, min, max); err != nil {
		t.Errorf(`validate.Rangelen(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "ABc12"
	if err := validate.Rangelen(in, min, max); err != nil {
		t.Errorf(`validate.Rangelen(%s, %d, %d) = %v, want %s`, in, min, max, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Rangelen(in, min, max); err == nil {
		t.Errorf(`validate.Rangelen(%s, %d, %d) = %v, want %s`, in, min, max, "nil", "not nil")
	}
}

func TestValidateFloat(t *testing.T) {
	in := "1.23"
	if err := validate.Float(in); err != nil {
		t.Errorf(`validate.Float(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "one"
	if err := validate.Float(in); err == nil {
		t.Errorf(`validate.Float(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateInteger(t *testing.T) {
	in := "123"
	if err := validate.Integer(in); err != nil {
		t.Errorf(`validate.Integer(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "1.23"
	if err := validate.Integer(in); err == nil {
		t.Errorf(`validate.Integer(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlpha(t *testing.T) {
	in := "ABc"
	if err := validate.Alpha(in); err != nil {
		t.Errorf(`validate.Alpha(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Alpha(in); err == nil {
		t.Errorf(`validate.Alpha(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlrod(t *testing.T) {
	in := "ABc"
	if err := validate.Alrod(in); err != nil {
		t.Errorf(`validate.Alrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "AB-c"
	if err := validate.Alrod(in); err != nil {
		t.Errorf(`validate.Alrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "AB_c"
	if err := validate.Alrod(in); err != nil {
		t.Errorf(`validate.Alrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Alrod(in); err == nil {
		t.Errorf(`validate.Alrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlnum(t *testing.T) {
	in := "ABc123"
	if err := validate.Alnum(in); err != nil {
		t.Errorf(`validate.Alnum(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc_123"
	if err := validate.Alnum(in); err == nil {
		t.Errorf(`validate.Alnum(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlnumrod(t *testing.T) {
	in := "ABc123"
	if err := validate.Alnumrod(in); err != nil {
		t.Errorf(`validate.Alnumrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc-123"
	if err := validate.Alnumrod(in); err != nil {
		t.Errorf(`validate.Alnumrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc_123"
	if err := validate.Alnumrod(in); err != nil {
		t.Errorf(`validate.Alnumrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc&123"
	if err := validate.Alnumrod(in); err == nil {
		t.Errorf(`validate.Alnumrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlnumhan(t *testing.T) {
	in := "ABc"
	if err := validate.Alnumhan(in); err != nil {
		t.Errorf(`validate.Alnumhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Alnumhan(in); err != nil {
		t.Errorf(`validate.Alnumhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc中文"
	if err := validate.Alnumhan(in); err != nil {
		t.Errorf(`validate.Alnumhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123中文"
	if err := validate.Alnumhan(in); err != nil {
		t.Errorf(`validate.Alnumhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc-123"
	if err := validate.Alnumhan(in); err == nil {
		t.Errorf(`validate.Alnumhan(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlnumhanrod(t *testing.T) {
	in := "ABc"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc-123"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc中文"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc_中文"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc123中文"
	if err := validate.Alnumhanrod(in); err != nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc&123"
	if err := validate.Alnumhanrod(in); err == nil {
		t.Errorf(`validate.Alnumhanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlhan(t *testing.T) {
	in := "ABc"
	if err := validate.Alhan(in); err != nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc中文"
	if err := validate.Alhan(in); err != nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "中文"
	if err := validate.Alhan(in); err != nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "123"
	if err := validate.Alhan(in); err == nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "ABc123"
	if err := validate.Alhan(in); err == nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "ABc_123"
	if err := validate.Alhan(in); err == nil {
		t.Errorf(`validate.Alhan(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateAlhanrod(t *testing.T) {
	in := "ABc"
	if err := validate.Alhanrod(in); err != nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc中文"
	if err := validate.Alhanrod(in); err != nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc-中文"
	if err := validate.Alhanrod(in); err != nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "中文"
	if err := validate.Alhanrod(in); err != nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "123"
	if err := validate.Alhanrod(in); err == nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "ABc123"
	if err := validate.Alhanrod(in); err == nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "ABc_123"
	if err := validate.Alhanrod(in); err == nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "ABc&123"
	if err := validate.Alhanrod(in); err == nil {
		t.Errorf(`validate.Alhanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateHan(t *testing.T) {
	in := "中文"
	if err := validate.Han(in); err != nil {
		t.Errorf(`validate.Han(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "ABc"
	if err := validate.Han(in); err == nil {
		t.Errorf(`validate.Han(%s) = %v, want %s`, in, "nil", "not nil")
	}
	in = "123"
	if err := validate.Han(in); err == nil {
		t.Errorf(`validate.Han(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateHanrod(t *testing.T) {
	in := "中文"
	if err := validate.Hanrod(in); err != nil {
		t.Errorf(`validate.Hanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "中-文"
	if err := validate.Hanrod(in); err != nil {
		t.Errorf(`validate.Hanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "中_文"
	if err := validate.Hanrod(in); err != nil {
		t.Errorf(`validate.Hanrod(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "中&文"
	if err := validate.Hanrod(in); err == nil {
		t.Errorf(`validate.Hanrod(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateEmail(t *testing.T) {
	in := "admin@goyy.org"
	if err := validate.Email(in); err != nil {
		t.Errorf(`validate.Email(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "admin"
	if err := validate.Email(in); err == nil {
		t.Errorf(`validate.Email(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateURL(t *testing.T) {
	in := "ftp://admin:123456@goyy.org/"
	if err := validate.URL(in); err != nil {
		t.Errorf(`validate.URL(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "admin@goyy.org"
	if err := validate.URL(in); err == nil {
		t.Errorf(`validate.URL(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateIP(t *testing.T) {
	in := "198.168.1.1"
	if err := validate.IP(in); err != nil {
		t.Errorf(`validate.IP(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "fe80::f8b4:6832:8fe4:fb98"
	if err := validate.IP(in); err != nil {
		t.Errorf(`validate.IP(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "198.168.1.256"
	if err := validate.IP(in); err == nil {
		t.Errorf(`validate.IP(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateMobile(t *testing.T) {
	in := "18611112222"
	if err := validate.Mobile(in); err != nil {
		t.Errorf(`validate.Mobile(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "13511112222"
	if err := validate.Mobile(in); err != nil {
		t.Errorf(`validate.Mobile(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "17011112222"
	if err := validate.Mobile(in); err != nil {
		t.Errorf(`validate.Mobile(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "+8618611112222"
	if err := validate.Mobile(in); err != nil {
		t.Errorf(`validate.Mobile(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "1111-2222"
	if err := validate.Mobile(in); err == nil {
		t.Errorf(`validate.Mobile(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateTel(t *testing.T) {
	in := "01011112222"
	if err := validate.Tel(in); err != nil {
		t.Errorf(`validate.Tel(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "010-11112222"
	if err := validate.Tel(in); err != nil {
		t.Errorf(`validate.Tel(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "11112222"
	if err := validate.Tel(in); err != nil {
		t.Errorf(`validate.Tel(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "1112222"
	if err := validate.Tel(in); err != nil {
		t.Errorf(`validate.Tel(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "1111-2222"
	if err := validate.Tel(in); err == nil {
		t.Errorf(`validate.Tel(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateZipcode(t *testing.T) {
	in := "623000"
	if err := validate.Zipcode(in); err != nil {
		t.Errorf(`validate.Zipcode(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "102629"
	if err := validate.Zipcode(in); err != nil {
		t.Errorf(`validate.Zipcode(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "335501"
	if err := validate.Zipcode(in); err != nil {
		t.Errorf(`validate.Zipcode(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "065201"
	if err := validate.Zipcode(in); err != nil {
		t.Errorf(`validate.Zipcode(%s) = %v, want %s`, in, err.Error(), "nil")
	}
	in = "010"
	if err := validate.Zipcode(in); err == nil {
		t.Errorf(`validate.Zipcode(%s) = %v, want %s`, in, "nil", "not nil")
	}
}

func TestValidateChain(t *testing.T) {
	fnOk := func(s string) error {
		return nil
	}
	err := validate.Chain("admin@goyy.org", validate.Email, fnOk)
	if err != nil {
		t.Errorf(`validate.Chain(...) = %v, want %s`, err.Error(), "nil")
	}
	fnErr := func(s string) error {
		return nil
	}
	err = validate.Chain("admin@goyy.org", validate.Required, validate.Email, fnErr, fnOk)
	if err != nil {
		t.Errorf(`validate.Chain(...) = %v, want %s`, err.Error(), "nil")
	}
}

func TestValidateEach(t *testing.T) {
	err := validate.Each(
		validate.Email("admin@goyy.org"),
		validate.Email("goyy.org"),
		validate.Email("admin"),
	)
	if err == nil {
		t.Errorf(`validate.Each(...) = %v, want %s`, "nil", "not nil")
	}
	err = validate.Each(
		validate.Email("admin@goyy.org"),
		validate.Required("goyy"),
		validate.Float("1.23"),
	)
	if err != nil {
		t.Errorf(`validate.Each(...) = %v, want %s`, err.Error(), "nil")
	}
}

func TestValidateAll(t *testing.T) {
	var err []error
	err = validate.All(
		validate.Email("admin@goyy.org"),
		validate.Email("goyy.org"),
		validate.Email("admin"),
	)
	if len(err) == 0 {
		t.Errorf(`validate.All(...) = %v, want %s`, "len(err) == 0", "len(err) > 0")
	}
	err = validate.All(
		validate.Email("admin@goyy.org"),
		validate.Required("goyy"),
		validate.Float("1.23"),
	)
	if len(err) != 0 {
		t.Errorf(`validate.All(...) = %v, want %s`, "len(err) > 0", "len(err) == 0")
	}
}

func TestValidateAny(t *testing.T) {
	err := validate.Any(
		validate.Email("admin@goyy.org"),
		validate.Email("goyy.org"),
		validate.Email("admin"),
	)
	if err != nil {
		t.Errorf(`validate.Any(...) = %v, want %s`, err.Error(), "nil")
	}
	err = validate.Any(
		validate.Email("admin@goyy.org"),
		validate.Required("goyy"),
		validate.Float("1.23"),
	)
	if err != nil {
		t.Errorf(`validate.Any(...) = %v, want %s`, err.Error(), "nil")
	}
	err = validate.Any(
		validate.Email("goyy.org"),
		validate.Email("123"),
		validate.Float("a"),
	)
	if err == nil {
		t.Errorf(`validate.Any(...) = %v, want %s`, "nil", "not nil")
	}
}
