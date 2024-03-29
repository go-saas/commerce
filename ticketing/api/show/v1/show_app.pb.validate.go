// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ticketing/api/show/v1/show_app.proto

package v1

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

// Validate checks the field values on PlaceShowOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlaceShowOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlaceShowOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlaceShowOrderRequestMultiError, or nil if none found.
func (m *PlaceShowOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PlaceShowOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetShowId()) < 1 {
		err := PlaceShowOrderRequestValidationError{
			field:  "ShowId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(m.GetSalesType()) < 1 {
		err := PlaceShowOrderRequestValidationError{
			field:  "SalesType",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetSalesType() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PlaceShowOrderRequestValidationError{
						field:  fmt.Sprintf("SalesType[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PlaceShowOrderRequestValidationError{
						field:  fmt.Sprintf("SalesType[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PlaceShowOrderRequestValidationError{
					field:  fmt.Sprintf("SalesType[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PlaceShowOrderRequestMultiError(errors)
	}

	return nil
}

// PlaceShowOrderRequestMultiError is an error wrapping multiple validation
// errors returned by PlaceShowOrderRequest.ValidateAll() if the designated
// constraints aren't met.
type PlaceShowOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlaceShowOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlaceShowOrderRequestMultiError) AllErrors() []error { return m }

// PlaceShowOrderRequestValidationError is the validation error returned by
// PlaceShowOrderRequest.Validate if the designated constraints aren't met.
type PlaceShowOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaceShowOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaceShowOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaceShowOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaceShowOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaceShowOrderRequestValidationError) ErrorName() string {
	return "PlaceShowOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PlaceShowOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaceShowOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaceShowOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaceShowOrderRequestValidationError{}

// Validate checks the field values on OrderSalesType with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OrderSalesType) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OrderSalesType with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OrderSalesTypeMultiError,
// or nil if none found.
func (m *OrderSalesType) ValidateAll() error {
	return m.validate(true)
}

func (m *OrderSalesType) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetShowSalesTypeId()) < 1 {
		err := OrderSalesTypeValidationError{
			field:  "ShowSalesTypeId",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetQty() <= 0 {
		err := OrderSalesTypeValidationError{
			field:  "Qty",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.ShowSeatId != nil {
		// no validation rules for ShowSeatId
	}

	if len(errors) > 0 {
		return OrderSalesTypeMultiError(errors)
	}

	return nil
}

// OrderSalesTypeMultiError is an error wrapping multiple validation errors
// returned by OrderSalesType.ValidateAll() if the designated constraints
// aren't met.
type OrderSalesTypeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OrderSalesTypeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OrderSalesTypeMultiError) AllErrors() []error { return m }

// OrderSalesTypeValidationError is the validation error returned by
// OrderSalesType.Validate if the designated constraints aren't met.
type OrderSalesTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OrderSalesTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OrderSalesTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OrderSalesTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OrderSalesTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OrderSalesTypeValidationError) ErrorName() string { return "OrderSalesTypeValidationError" }

// Error satisfies the builtin error interface
func (e OrderSalesTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOrderSalesType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OrderSalesTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OrderSalesTypeValidationError{}

// Validate checks the field values on PlaceShowOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PlaceShowOrderReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PlaceShowOrderReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PlaceShowOrderReplyMultiError, or nil if none found.
func (m *PlaceShowOrderReply) ValidateAll() error {
	return m.validate(true)
}

func (m *PlaceShowOrderReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return PlaceShowOrderReplyMultiError(errors)
	}

	return nil
}

// PlaceShowOrderReplyMultiError is an error wrapping multiple validation
// errors returned by PlaceShowOrderReply.ValidateAll() if the designated
// constraints aren't met.
type PlaceShowOrderReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PlaceShowOrderReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PlaceShowOrderReplyMultiError) AllErrors() []error { return m }

// PlaceShowOrderReplyValidationError is the validation error returned by
// PlaceShowOrderReply.Validate if the designated constraints aren't met.
type PlaceShowOrderReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PlaceShowOrderReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PlaceShowOrderReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PlaceShowOrderReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PlaceShowOrderReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PlaceShowOrderReplyValidationError) ErrorName() string {
	return "PlaceShowOrderReplyValidationError"
}

// Error satisfies the builtin error interface
func (e PlaceShowOrderReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPlaceShowOrderReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PlaceShowOrderReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PlaceShowOrderReplyValidationError{}
