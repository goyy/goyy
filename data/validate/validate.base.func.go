// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validate

import (
	"gopkg.in/goyy/goyy.v0/util/errors"
	"gopkg.in/goyy/goyy.v0/util/strings"
	"regexp"
	"strconv"
	"unicode/utf8"
)

// Returns error if the provided input is empty, nil otherwise.
func Required(input string) error {
	if strings.IsBlank(input) {
		return errors.New(Messages[typRequired])
	}
	return nil
}

// Returns error if the provided input is greater than or equal to {min},
// nil otherwise.
func Min(input string, min int) error {
	if strings.IsBlank(input) {
		return errors.Newf(Messages[typMin], min)
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(Messages[typMin], min)
	}
	if v < min {
		return errors.Newf(Messages[typMin], min)
	}
	return nil
}

// Returns error if the provided input is less than or equal to {max},
// nil otherwise.
func Max(input string, max int) error {
	if strings.IsBlank(input) {
		return errors.Newf(Messages[typMax], max)
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(Messages[typMax], max)
	}
	if v > max {
		return errors.Newf(Messages[typMax], max)
	}
	return nil
}

// Returns error if the provided input is between {min} and {max},
// nil otherwise.
func Range(input string, min, max int) error {
	if strings.IsBlank(input) {
		return errors.Newf(Messages[typRange], min, max)
	}
	v, err := strconv.Atoi(input)
	if err != nil {
		return errors.Newf(Messages[typRange], min, max)
	}
	if v < min || v > max {
		return errors.Newf(Messages[typRange], min, max)
	}
	return nil
}

// Returns error if the provided input is least {min} characters,
// nil otherwise.
func Minlen(input string, min int) error {
	if utf8.RuneCountInString(input) < min {
		return errors.Newf(Messages[typMinlen], min)
	}
	return nil
}

// Returns error if the provided input is more than {max} characters,
// nil otherwise.
func Maxlen(input string, max int) error {
	if utf8.RuneCountInString(input) > max {
		return errors.Newf(Messages[typMaxlen], max)
	}
	return nil
}

// Returns error if the provided input is between {min} and {max} characters
// long, nil otherwise.
func Rangelen(input string, min, max int) error {
	l := utf8.RuneCountInString(input)
	if l < min || l > max {
		return errors.Newf(Messages[typRangelen], min, max)
	}
	return nil
}

// Returns error if the provided input is not a floating point number, nil
// otherwise.
func Float(input string) error {
	if Rules[typFloat].MatchString(input) == false {
		return errors.New(Messages[typFloat])
	}
	return nil
}

// Returns error if the provided input is not an integer value, nil otherwise.
func Integer(input string) error {
	if Rules[typInteger].MatchString(input) == false {
		return errors.New(Messages[typInteger])
	}
	return nil
}

// Returns error if the provided input is not an alphanumeric (a-zA-Z0-9)
// string, nil otherwise.
func Alphanumeric(input string) error {
	if Rules[typAlphanumeric].MatchString(input) == false {
		return errors.New(Messages[typAlphanumeric])
	}
	return nil
}

// Returns error if the provided input is not an alphabetic (a-zA-Z) string,
// nil otherwise.
func Alphabetic(input string) error {
	if Rules[typAlphabetic].MatchString(input) == false {
		return errors.New(Messages[typAlphabetic])
	}
	return nil
}

// Returns error if the provided input is not match {regexp} string,
// nil otherwise.
func Match(input string, regexp *regexp.Regexp) error {
	if regexp.MatchString(input) == false {
		return errors.New(Messages[typMatch])
	}
	return nil
}

// Returns error if the provided input is match {regexp} string,
// nil otherwise.
func Nomatch(input string, regexp *regexp.Regexp) error {
	if regexp.MatchString(input) == true {
		return errors.New(Messages[typNomatch])
	}
	return nil
}

// Returns error if the provided input is not an email, nil otherwise.
func Email(input string) error {
	if Rules[typEmail].MatchString(input) == false {
		return errors.New(Messages[typEmail])
	}
	return nil
}

// Returns error if the provided input is not an URL, nil otherwise.
func URL(input string) error {
	if Rules[typURL].MatchString(input) == false {
		return errors.New(Messages[typURL])
	}
	return nil
}

// Returns error if the provided input is not an IP, nil otherwise.
func IP(input string) error {
	if Rules[typIP].MatchString(input) == false {
		return errors.New(Messages[typIP])
	}
	return nil
}

// Returns error if the provided input is not an mobile, nil otherwise.
func Mobile(input string) error {
	if Rules[typMobile].MatchString(input) == false {
		return errors.New(Messages[typMobile])
	}
	return nil
}

// Returns error if the provided input is not an tel, nil otherwise.
func Tel(input string) error {
	if Rules[typTel].MatchString(input) == false {
		return errors.New(Messages[typTel])
	}
	return nil
}

// Returns error if the provided input is not an phone, nil otherwise.
func Phone(input string) error {
	if Rules[typMobile].MatchString(input) == false && Rules[typTel].MatchString(input) == false {
		return errors.New(Messages[typPhone])
	}
	return nil
}

// Returns error if the provided input is not an zipcode, nil otherwise.
func Zipcode(input string) error {
	if Rules[typZipcode].MatchString(input) == false {
		return errors.New(Messages[typZipcode])
	}
	return nil
}

// This function takes an input an applies the given set of validation functions
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

// This function accepts a list of error values (from values or functions) and
// returns the first error found, if any.
//
// Example:
//
// err := validate.Each(
//   validate.Email(userEmail),
//	 validate.Chain(userName, validate.Required),
// )
func Each(tests ...error) error {
	for i, _ := range tests {
		err := tests[i]
		if err != nil {
			return err
		}
	}
	return nil
}

// This function accepts a list of error values (from values or functions) and
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

	for i, _ := range tests {
		err := tests[i]
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

// This function accepts a list of error values (from values or functions) and
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

	for i, _ := range tests {
		err := tests[i]
		if err == nil {
			return nil
		} else {
			last = err
		}
	}

	return last
}
