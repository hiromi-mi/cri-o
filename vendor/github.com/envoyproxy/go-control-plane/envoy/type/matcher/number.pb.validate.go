// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/type/matcher/number.proto

package matcher

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on DoubleMatcher with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DoubleMatcher) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DoubleMatcher with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DoubleMatcherMultiError, or
// nil if none found.
func (m *DoubleMatcher) ValidateAll() error {
	return m.validate(true)
}

func (m *DoubleMatcher) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch m.MatchPattern.(type) {

	case *DoubleMatcher_Range:

		if all {
			switch v := interface{}(m.GetRange()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, DoubleMatcherValidationError{
						field:  "Range",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, DoubleMatcherValidationError{
						field:  "Range",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DoubleMatcherValidationError{
					field:  "Range",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *DoubleMatcher_Exact:
		// no validation rules for Exact

	default:
		err := DoubleMatcherValidationError{
			field:  "MatchPattern",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if len(errors) > 0 {
		return DoubleMatcherMultiError(errors)
	}

	return nil
}

// DoubleMatcherMultiError is an error wrapping multiple validation errors
// returned by DoubleMatcher.ValidateAll() if the designated constraints
// aren't met.
type DoubleMatcherMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DoubleMatcherMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DoubleMatcherMultiError) AllErrors() []error { return m }

// DoubleMatcherValidationError is the validation error returned by
// DoubleMatcher.Validate if the designated constraints aren't met.
type DoubleMatcherValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DoubleMatcherValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DoubleMatcherValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DoubleMatcherValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DoubleMatcherValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DoubleMatcherValidationError) ErrorName() string { return "DoubleMatcherValidationError" }

// Error satisfies the builtin error interface
func (e DoubleMatcherValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDoubleMatcher.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DoubleMatcherValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DoubleMatcherValidationError{}
