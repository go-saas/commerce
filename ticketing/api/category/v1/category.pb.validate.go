// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ticketing/api/category/v1/category.proto

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

// Validate checks the field values on CreateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateCategoryRequestMultiError, or nil if none found.
func (m *CreateCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateCategoryRequestValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Parent

	if len(errors) > 0 {
		return CreateCategoryRequestMultiError(errors)
	}

	return nil
}

// CreateCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by CreateCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type CreateCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateCategoryRequestMultiError) AllErrors() []error { return m }

// CreateCategoryRequestValidationError is the validation error returned by
// CreateCategoryRequest.Validate if the designated constraints aren't met.
type CreateCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateCategoryRequestValidationError) ErrorName() string {
	return "CreateCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateCategoryRequestValidationError{}

// Validate checks the field values on UpdateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateCategoryRequestMultiError, or nil if none found.
func (m *UpdateCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetCategory() == nil {
		err := UpdateCategoryRequestValidationError{
			field:  "Category",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetCategory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "Category",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCategory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCategoryRequestValidationError{
				field:  "Category",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdateMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateCategoryRequestValidationError{
					field:  "UpdateMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdateMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateCategoryRequestValidationError{
				field:  "UpdateMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateCategoryRequestMultiError(errors)
	}

	return nil
}

// UpdateCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCategoryRequestMultiError) AllErrors() []error { return m }

// UpdateCategoryRequestValidationError is the validation error returned by
// UpdateCategoryRequest.Validate if the designated constraints aren't met.
type UpdateCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCategoryRequestValidationError) ErrorName() string {
	return "UpdateCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCategoryRequestValidationError{}

// Validate checks the field values on UpdateCategory with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *UpdateCategory) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateCategory with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in UpdateCategoryMultiError,
// or nil if none found.
func (m *UpdateCategory) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateCategory) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := UpdateCategoryValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for Parent

	if len(errors) > 0 {
		return UpdateCategoryMultiError(errors)
	}

	return nil
}

// UpdateCategoryMultiError is an error wrapping multiple validation errors
// returned by UpdateCategory.ValidateAll() if the designated constraints
// aren't met.
type UpdateCategoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateCategoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateCategoryMultiError) AllErrors() []error { return m }

// UpdateCategoryValidationError is the validation error returned by
// UpdateCategory.Validate if the designated constraints aren't met.
type UpdateCategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateCategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateCategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateCategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateCategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateCategoryValidationError) ErrorName() string { return "UpdateCategoryValidationError" }

// Error satisfies the builtin error interface
func (e UpdateCategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateCategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateCategoryValidationError{}

// Validate checks the field values on DeleteCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCategoryRequestMultiError, or nil if none found.
func (m *DeleteCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return DeleteCategoryRequestMultiError(errors)
	}

	return nil
}

// DeleteCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by DeleteCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type DeleteCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCategoryRequestMultiError) AllErrors() []error { return m }

// DeleteCategoryRequestValidationError is the validation error returned by
// DeleteCategoryRequest.Validate if the designated constraints aren't met.
type DeleteCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCategoryRequestValidationError) ErrorName() string {
	return "DeleteCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCategoryRequestValidationError{}

// Validate checks the field values on DeleteCategoryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCategoryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCategoryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCategoryReplyMultiError, or nil if none found.
func (m *DeleteCategoryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCategoryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Name

	if len(errors) > 0 {
		return DeleteCategoryReplyMultiError(errors)
	}

	return nil
}

// DeleteCategoryReplyMultiError is an error wrapping multiple validation
// errors returned by DeleteCategoryReply.ValidateAll() if the designated
// constraints aren't met.
type DeleteCategoryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCategoryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCategoryReplyMultiError) AllErrors() []error { return m }

// DeleteCategoryReplyValidationError is the validation error returned by
// DeleteCategoryReply.Validate if the designated constraints aren't met.
type DeleteCategoryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCategoryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCategoryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCategoryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCategoryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCategoryReplyValidationError) ErrorName() string {
	return "DeleteCategoryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCategoryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCategoryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCategoryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCategoryReplyValidationError{}

// Validate checks the field values on GetCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCategoryRequestMultiError, or nil if none found.
func (m *GetCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetKey()) < 1 {
		err := GetCategoryRequestValidationError{
			field:  "Key",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetCategoryRequestMultiError(errors)
	}

	return nil
}

// GetCategoryRequestMultiError is an error wrapping multiple validation errors
// returned by GetCategoryRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCategoryRequestMultiError) AllErrors() []error { return m }

// GetCategoryRequestValidationError is the validation error returned by
// GetCategoryRequest.Validate if the designated constraints aren't met.
type GetCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCategoryRequestValidationError) ErrorName() string {
	return "GetCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCategoryRequestValidationError{}

// Validate checks the field values on CategoryFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CategoryFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CategoryFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CategoryFilterMultiError,
// or nil if none found.
func (m *CategoryFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *CategoryFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetKey()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Key",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Key",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetKey()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryFilterValidationError{
				field:  "Key",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetName()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Name",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Name",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetName()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryFilterValidationError{
				field:  "Name",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetParent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Parent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryFilterValidationError{
					field:  "Parent",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetParent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryFilterValidationError{
				field:  "Parent",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CategoryFilterMultiError(errors)
	}

	return nil
}

// CategoryFilterMultiError is an error wrapping multiple validation errors
// returned by CategoryFilter.ValidateAll() if the designated constraints
// aren't met.
type CategoryFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryFilterMultiError) AllErrors() []error { return m }

// CategoryFilterValidationError is the validation error returned by
// CategoryFilter.Validate if the designated constraints aren't met.
type CategoryFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryFilterValidationError) ErrorName() string { return "CategoryFilterValidationError" }

// Error satisfies the builtin error interface
func (e CategoryFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoryFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryFilterValidationError{}

// Validate checks the field values on ListCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ListCategoryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCategoryRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCategoryRequestMultiError, or nil if none found.
func (m *ListCategoryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCategoryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageOffset

	// no validation rules for PageSize

	// no validation rules for Search

	if all {
		switch v := interface{}(m.GetFields()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListCategoryRequestValidationError{
					field:  "Fields",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListCategoryRequestValidationError{
					field:  "Fields",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFields()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCategoryRequestValidationError{
				field:  "Fields",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetFilter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ListCategoryRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListCategoryRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListCategoryRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ListCategoryRequestMultiError(errors)
	}

	return nil
}

// ListCategoryRequestMultiError is an error wrapping multiple validation
// errors returned by ListCategoryRequest.ValidateAll() if the designated
// constraints aren't met.
type ListCategoryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCategoryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCategoryRequestMultiError) AllErrors() []error { return m }

// ListCategoryRequestValidationError is the validation error returned by
// ListCategoryRequest.Validate if the designated constraints aren't met.
type ListCategoryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCategoryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCategoryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCategoryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCategoryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCategoryRequestValidationError) ErrorName() string {
	return "ListCategoryRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListCategoryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCategoryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCategoryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCategoryRequestValidationError{}

// Validate checks the field values on ListCategoryReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListCategoryReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListCategoryReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListCategoryReplyMultiError, or nil if none found.
func (m *ListCategoryReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListCategoryReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for TotalSize

	// no validation rules for FilterSize

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListCategoryReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListCategoryReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListCategoryReplyValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListCategoryReplyMultiError(errors)
	}

	return nil
}

// ListCategoryReplyMultiError is an error wrapping multiple validation errors
// returned by ListCategoryReply.ValidateAll() if the designated constraints
// aren't met.
type ListCategoryReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListCategoryReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListCategoryReplyMultiError) AllErrors() []error { return m }

// ListCategoryReplyValidationError is the validation error returned by
// ListCategoryReply.Validate if the designated constraints aren't met.
type ListCategoryReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListCategoryReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListCategoryReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListCategoryReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListCategoryReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListCategoryReplyValidationError) ErrorName() string {
	return "ListCategoryReplyValidationError"
}

// Error satisfies the builtin error interface
func (e ListCategoryReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListCategoryReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListCategoryReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListCategoryReplyValidationError{}

// Validate checks the field values on Category with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Category) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Category with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CategoryMultiError, or nil
// if none found.
func (m *Category) ValidateAll() error {
	return m.validate(true)
}

func (m *Category) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	// no validation rules for Name

	// no validation rules for Path

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CategoryValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CategoryValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CategoryValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Parent

	if len(errors) > 0 {
		return CategoryMultiError(errors)
	}

	return nil
}

// CategoryMultiError is an error wrapping multiple validation errors returned
// by Category.ValidateAll() if the designated constraints aren't met.
type CategoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CategoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CategoryMultiError) AllErrors() []error { return m }

// CategoryValidationError is the validation error returned by
// Category.Validate if the designated constraints aren't met.
type CategoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoryValidationError) ErrorName() string { return "CategoryValidationError" }

// Error satisfies the builtin error interface
func (e CategoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoryValidationError{}
