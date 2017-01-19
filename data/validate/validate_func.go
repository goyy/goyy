// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"regexp"
	"strconv"
	"unicode/utf8"

	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// Required returns error if the provided input is empty, nil otherwise.
func Required(input string) error {
	if strings.IsBlank(input) {
		return errors.New(i18N.Message(typRequired))
	}
	return nil
}

// Min returns error if the provided input is greater than or equal to {min},
// nil otherwise.
func Min(input string, min int) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(i18N.Message(typMin), min)
	}
	if v < min {
		return errors.Newf(i18N.Message(typMin), min)
	}
	return nil
}

// Max returns error if the provided input is less than or equal to {max},
// nil otherwise.
func Max(input string, max int) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(i18N.Message(typMax), max)
	}
	if v > max {
		return errors.Newf(i18N.Message(typMax), max)
	}
	return nil
}

// Range returns error if the provided input is between {min} and {max},
// nil otherwise.
func Range(input string, min, max int) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(i18N.Message(typRange), min, max)
	}
	if v < min || v > max {
		return errors.Newf(i18N.Message(typRange), min, max)
	}
	return nil
}

// Minf returns error if the provided input is greater than or equal to {min},
// nil otherwise.
func Minf(input string, min float64) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.Newf(i18N.Message(typMinf), min)
	}
	if v < min {
		return errors.Newf(i18N.Message(typMinf), min)
	}
	return nil
}

// Maxf returns error if the provided input is less than or equal to {max},
// nil otherwise.
func Maxf(input string, max float64) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.Newf(i18N.Message(typMaxf), max)
	}
	if v > max {
		return errors.Newf(i18N.Message(typMaxf), max)
	}
	return nil
}

// Rangef returns error if the provided input is between {min} and {max},
// nil otherwise.
func Rangef(input string, min, max float64) error {
	if strings.IsBlank(input) {
		return nil
	}
	v, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return errors.Newf(i18N.Message(typRangef), min, max)
	}
	if v < min || v > max {
		return errors.Newf(i18N.Message(typRangef), min, max)
	}
	return nil
}

// Minlen returns error if the provided input is least {min} characters,
// nil otherwise.
func Minlen(input string, min int) error {
	if strings.IsBlank(input) {
		return nil
	}
	if utf8.RuneCountInString(input) < min {
		return errors.Newf(i18N.Message(typMinlen), min)
	}
	return nil
}

// Maxlen returns error if the provided input is more than {max} characters,
// nil otherwise.
func Maxlen(input string, max int) error {
	if strings.IsBlank(input) {
		return nil
	}
	if utf8.RuneCountInString(input) > max {
		return errors.Newf(i18N.Message(typMaxlen), max)
	}
	return nil
}

// Rangelen returns error if the provided input is between {min} and {max} characters
// long, nil otherwise.
func Rangelen(input string, min, max int) error {
	if strings.IsBlank(input) {
		return nil
	}
	l := utf8.RuneCountInString(input)
	if l < min || l > max {
		return errors.Newf(i18N.Message(typRangelen), min, max)
	}
	return nil
}

// Float returns error if the provided input is not a floating point number, nil
// otherwise.
func Float(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typFloat].MatchString(input) == false {
		return errors.New(i18N.Message(typFloat))
	}
	return nil
}

// Integer returns error if the provided input is not an integer value, nil otherwise.
func Integer(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typInteger].MatchString(input) == false {
		return errors.New(i18N.Message(typInteger))
	}
	return nil
}

// Alpha returns error if the provided input is not an alphabetic (a-zA-Z) string,
// nil otherwise.
func Alpha(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlpha].MatchString(input) == false {
		return errors.New(i18N.Message(typAlpha))
	}
	return nil
}

// Alrod returns error if the provided input is not an alphabetic or rod (a-zA-Z\-_) string,
// nil otherwise.
func Alrod(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlrod].MatchString(input) == false {
		return errors.New(i18N.Message(typAlrod))
	}
	return nil
}

// Alnum returns error if the provided input is not an alphanumeric (a-zA-Z0-9)
// string, nil otherwise.
func Alnum(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlnum].MatchString(input) == false {
		return errors.New(i18N.Message(typAlnum))
	}
	return nil
}

// Alnumrod returns error if the provided input is not an alphanumeric or rod (a-zA-Z0-9\-_)
// string, nil otherwise.
func Alnumrod(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlnumrod].MatchString(input) == false {
		return errors.New(i18N.Message(typAlnumrod))
	}
	return nil
}

// Alnumhan returns error if the provided input is not an alphanumeric or chinese (a-zA-Z0-9\p{Han})
// string, nil otherwise.
func Alnumhan(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlnumhan].MatchString(input) == false {
		return errors.New(i18N.Message(typAlnumhan))
	}
	return nil
}

// Alnumhanrod returns error if the provided input is not an alphanumeric or chinese or rod (a-zA-Z0-9\p{Han}\-_)
// string, nil otherwise.
func Alnumhanrod(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlnumhanrod].MatchString(input) == false {
		return errors.New(i18N.Message(typAlnumhanrod))
	}
	return nil
}

// Alhan returns error if the provided input is not an alphabetic or chinese (a-zA-Z\p{Han})
// string, nil otherwise.
func Alhan(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlhan].MatchString(input) == false {
		return errors.New(i18N.Message(typAlhan))
	}
	return nil
}

// Alhanrod returns error if the provided input is not an alphabetic or chinese or rod (a-zA-Z\p{Han}\-_)
// string, nil otherwise.
func Alhanrod(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typAlhanrod].MatchString(input) == false {
		return errors.New(i18N.Message(typAlhanrod))
	}
	return nil
}

// Han returns error if the provided input is not an alphabetic or chinese (\p{Han})
// string, nil otherwise.
func Han(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typHan].MatchString(input) == false {
		return errors.New(i18N.Message(typHan))
	}
	return nil
}

// Hanrod returns error if the provided input is not an alphabetic or chinese or rod (\p{Han}\-_)
// string, nil otherwise.
func Hanrod(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typHanrod].MatchString(input) == false {
		return errors.New(i18N.Message(typHanrod))
	}
	return nil
}

// Match returns error if the provided input is not match {regexp} string,
// nil otherwise.
func Match(input, regexps string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if regexp.MustCompile(regexps).MatchString(input) == false {
		return errors.New(i18N.Message(typMatch))
	}
	return nil
}

// Nomatch returns error if the provided input is match {regexp} string,
// nil otherwise.
func Nomatch(input, regexps string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if regexp.MustCompile(regexps).MatchString(input) == true {
		return errors.New(i18N.Message(typNomatch))
	}
	return nil
}

// Email returns error if the provided input is not an email, nil otherwise.
func Email(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typEmail].MatchString(input) == false {
		return errors.New(i18N.Message(typEmail))
	}
	return nil
}

// URL returns error if the provided input is not an URL, nil otherwise.
func URL(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typURL].MatchString(input) == false {
		return errors.New(i18N.Message(typURL))
	}
	return nil
}

// IP returns error if the provided input is not an IP, nil otherwise.
func IP(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typIP].MatchString(input) == false {
		return errors.New(i18N.Message(typIP))
	}
	return nil
}

// Mobile returns error if the provided input is not an mobile, nil otherwise.
func Mobile(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typMobile].MatchString(input) == false {
		return errors.New(i18N.Message(typMobile))
	}
	return nil
}

// Tel returns error if the provided input is not an tel, nil otherwise.
func Tel(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typTel].MatchString(input) == false {
		return errors.New(i18N.Message(typTel))
	}
	return nil
}

// Phone returns error if the provided input is not an phone, nil otherwise.
func Phone(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typMobile].MatchString(input) == false && rules[typTel].MatchString(input) == false {
		return errors.New(i18N.Message(typPhone))
	}
	return nil
}

// Zipcode returns error if the provided input is not an zipcode, nil otherwise.
func Zipcode(input string) error {
	if strings.IsBlank(input) {
		return nil
	}
	if rules[typZipcode].MatchString(input) == false {
		return errors.New(i18N.Message(typZipcode))
	}
	return nil
}

// Chain this function takes an input an applies the given set of validation functions
// in order, each function is a link of the chain. If any validation fails,
// validate.Chain stops and returns the error.
//
// Example:
//
// err := validate.Chain(userEmail, validate.Required, validate.Email)
//
func Chain(input string, links ...func(string) error) error {
	var err error
	for _, link := range links {
		err = link(input)
		if err != nil {
			return err
		}
	}
	return nil
}

// Each this function accepts a list of error values (from values or functions) and
// returns the first error found, if any.
//
// Example:
//
// err := validate.Each(
//   validate.Email(userEmail),
//	 validate.Chain(userName, validate.Required),
// )
func Each(tests ...error) error {
	for i := range tests {
		err := tests[i]
		if err != nil {
			return err
		}
	}
	return nil
}

// All this function accepts a list of error values (from values or functions) and
// returns an array of errors values, useful for validating all user inputs at
// once.
//
// Example:
//
// err := validate.All(
//   validate.Email(userEmail),
//	 validate.Chain(userName, validate.Required),
// )
func All(tests ...error) []error {
	errs := make([]error, 0, len(tests))

	for i := range tests {
		err := tests[i]
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// Any this function accepts a list of error values (from values or functions) and
// returns nil if any of the rules is valid.
//
// Example:
//
// err := validate.Any(
//   validate.Required(userAge),
//   validate.Integer(userAge),
// )
func Any(tests ...error) error {
	var last error

	for i := range tests {
		err := tests[i]
		if err != nil {
			last = err
		}
		return nil
	}

	return last
}
