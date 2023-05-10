// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: order/api/order/v1/order_internal.proto

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

// Validate checks the field values on CreateInternalOrderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateInternalOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateInternalOrderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateInternalOrderRequestMultiError, or nil if none found.
func (m *CreateInternalOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateInternalOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetShippingAddr()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "ShippingAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "ShippingAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShippingAddr()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderRequestValidationError{
				field:  "ShippingAddr",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetBillingAddr()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "BillingAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "BillingAddr",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBillingAddr()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderRequestValidationError{
				field:  "BillingAddr",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for CustomerId

	if all {
		switch v := interface{}(m.GetExtra()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderRequestValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtra()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderRequestValidationError{
				field:  "Extra",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetItems()) < 1 {
		err := CreateInternalOrderRequestValidationError{
			field:  "Items",
			reason: "value must contain at least 1 item(s)",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateInternalOrderRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateInternalOrderRequestValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateInternalOrderRequestValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.PayBefore != nil {

		if all {
			switch v := interface{}(m.GetPayBefore()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CreateInternalOrderRequestValidationError{
						field:  "PayBefore",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CreateInternalOrderRequestValidationError{
						field:  "PayBefore",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetPayBefore()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CreateInternalOrderRequestValidationError{
					field:  "PayBefore",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CreateInternalOrderRequestMultiError(errors)
	}

	return nil
}

// CreateInternalOrderRequestMultiError is an error wrapping multiple
// validation errors returned by CreateInternalOrderRequest.ValidateAll() if
// the designated constraints aren't met.
type CreateInternalOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateInternalOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateInternalOrderRequestMultiError) AllErrors() []error { return m }

// CreateInternalOrderRequestValidationError is the validation error returned
// by CreateInternalOrderRequest.Validate if the designated constraints aren't met.
type CreateInternalOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateInternalOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateInternalOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateInternalOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateInternalOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateInternalOrderRequestValidationError) ErrorName() string {
	return "CreateInternalOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateInternalOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateInternalOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateInternalOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateInternalOrderRequestValidationError{}

// Validate checks the field values on CreateInternalOrderItem with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateInternalOrderItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateInternalOrderItem with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateInternalOrderItemMultiError, or nil if none found.
func (m *CreateInternalOrderItem) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateInternalOrderItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Qty

	if all {
		switch v := interface{}(m.GetPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "Price",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderItemValidationError{
				field:  "Price",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOriginalPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "OriginalPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "OriginalPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOriginalPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderItemValidationError{
				field:  "OriginalPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsGiveaway

	if m.GetProduct() == nil {
		err := CreateInternalOrderItemValidationError{
			field:  "Product",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetProduct()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateInternalOrderItemValidationError{
					field:  "Product",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProduct()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateInternalOrderItemValidationError{
				field:  "Product",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateInternalOrderItemMultiError(errors)
	}

	return nil
}

// CreateInternalOrderItemMultiError is an error wrapping multiple validation
// errors returned by CreateInternalOrderItem.ValidateAll() if the designated
// constraints aren't met.
type CreateInternalOrderItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateInternalOrderItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateInternalOrderItemMultiError) AllErrors() []error { return m }

// CreateInternalOrderItemValidationError is the validation error returned by
// CreateInternalOrderItem.Validate if the designated constraints aren't met.
type CreateInternalOrderItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateInternalOrderItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateInternalOrderItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateInternalOrderItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateInternalOrderItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateInternalOrderItemValidationError) ErrorName() string {
	return "CreateInternalOrderItemValidationError"
}

// Error satisfies the builtin error interface
func (e CreateInternalOrderItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateInternalOrderItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateInternalOrderItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateInternalOrderItemValidationError{}

// Validate checks the field values on GetInternalOrderRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetInternalOrderRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetInternalOrderRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetInternalOrderRequestMultiError, or nil if none found.
func (m *GetInternalOrderRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetInternalOrderRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetInternalOrderRequestMultiError(errors)
	}

	return nil
}

// GetInternalOrderRequestMultiError is an error wrapping multiple validation
// errors returned by GetInternalOrderRequest.ValidateAll() if the designated
// constraints aren't met.
type GetInternalOrderRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetInternalOrderRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetInternalOrderRequestMultiError) AllErrors() []error { return m }

// GetInternalOrderRequestValidationError is the validation error returned by
// GetInternalOrderRequest.Validate if the designated constraints aren't met.
type GetInternalOrderRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetInternalOrderRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetInternalOrderRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetInternalOrderRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetInternalOrderRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetInternalOrderRequestValidationError) ErrorName() string {
	return "GetInternalOrderRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetInternalOrderRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetInternalOrderRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetInternalOrderRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetInternalOrderRequestValidationError{}

// Validate checks the field values on InternalOrderPaySuccessRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InternalOrderPaySuccessRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalOrderPaySuccessRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// InternalOrderPaySuccessRequestMultiError, or nil if none found.
func (m *InternalOrderPaySuccessRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalOrderPaySuccessRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := InternalOrderPaySuccessRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetPayExtra()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PayExtra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PayExtra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayExtra()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalOrderPaySuccessRequestValidationError{
				field:  "PayExtra",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPaidTime()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PaidTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PaidTime",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaidTime()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalOrderPaySuccessRequestValidationError{
				field:  "PaidTime",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetPaidPrice()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PaidPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, InternalOrderPaySuccessRequestValidationError{
					field:  "PaidPrice",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPaidPrice()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return InternalOrderPaySuccessRequestValidationError{
				field:  "PaidPrice",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for PayWay

	if len(errors) > 0 {
		return InternalOrderPaySuccessRequestMultiError(errors)
	}

	return nil
}

// InternalOrderPaySuccessRequestMultiError is an error wrapping multiple
// validation errors returned by InternalOrderPaySuccessRequest.ValidateAll()
// if the designated constraints aren't met.
type InternalOrderPaySuccessRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalOrderPaySuccessRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalOrderPaySuccessRequestMultiError) AllErrors() []error { return m }

// InternalOrderPaySuccessRequestValidationError is the validation error
// returned by InternalOrderPaySuccessRequest.Validate if the designated
// constraints aren't met.
type InternalOrderPaySuccessRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalOrderPaySuccessRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalOrderPaySuccessRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalOrderPaySuccessRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalOrderPaySuccessRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalOrderPaySuccessRequestValidationError) ErrorName() string {
	return "InternalOrderPaySuccessRequestValidationError"
}

// Error satisfies the builtin error interface
func (e InternalOrderPaySuccessRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalOrderPaySuccessRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalOrderPaySuccessRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalOrderPaySuccessRequestValidationError{}

// Validate checks the field values on InternalOrderPaySuccessReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *InternalOrderPaySuccessReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InternalOrderPaySuccessReply with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// InternalOrderPaySuccessReplyMultiError, or nil if none found.
func (m *InternalOrderPaySuccessReply) ValidateAll() error {
	return m.validate(true)
}

func (m *InternalOrderPaySuccessReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return InternalOrderPaySuccessReplyMultiError(errors)
	}

	return nil
}

// InternalOrderPaySuccessReplyMultiError is an error wrapping multiple
// validation errors returned by InternalOrderPaySuccessReply.ValidateAll() if
// the designated constraints aren't met.
type InternalOrderPaySuccessReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InternalOrderPaySuccessReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InternalOrderPaySuccessReplyMultiError) AllErrors() []error { return m }

// InternalOrderPaySuccessReplyValidationError is the validation error returned
// by InternalOrderPaySuccessReply.Validate if the designated constraints
// aren't met.
type InternalOrderPaySuccessReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InternalOrderPaySuccessReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InternalOrderPaySuccessReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InternalOrderPaySuccessReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InternalOrderPaySuccessReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InternalOrderPaySuccessReplyValidationError) ErrorName() string {
	return "InternalOrderPaySuccessReplyValidationError"
}

// Error satisfies the builtin error interface
func (e InternalOrderPaySuccessReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInternalOrderPaySuccessReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InternalOrderPaySuccessReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InternalOrderPaySuccessReplyValidationError{}
