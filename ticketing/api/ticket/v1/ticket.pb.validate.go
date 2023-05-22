// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: ticketing/api/ticket/v1/ticket.proto

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

// Validate checks the field values on GetTicketRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetTicketRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetTicketRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetTicketRequestMultiError, or nil if none found.
func (m *GetTicketRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetTicketRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetId()) < 1 {
		err := GetTicketRequestValidationError{
			field:  "Id",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetTicketRequestMultiError(errors)
	}

	return nil
}

// GetTicketRequestMultiError is an error wrapping multiple validation errors
// returned by GetTicketRequest.ValidateAll() if the designated constraints
// aren't met.
type GetTicketRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetTicketRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetTicketRequestMultiError) AllErrors() []error { return m }

// GetTicketRequestValidationError is the validation error returned by
// GetTicketRequest.Validate if the designated constraints aren't met.
type GetTicketRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetTicketRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetTicketRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetTicketRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetTicketRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetTicketRequestValidationError) ErrorName() string { return "GetTicketRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetTicketRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetTicketRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetTicketRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetTicketRequestValidationError{}

// Validate checks the field values on TicketFilter with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *TicketFilter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on TicketFilter with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in TicketFilterMultiError, or
// nil if none found.
func (m *TicketFilter) ValidateAll() error {
	return m.validate(true)
}

func (m *TicketFilter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "Id",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketFilterValidationError{
				field:  "Id",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetLocationId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "LocationId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "LocationId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocationId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketFilterValidationError{
				field:  "LocationId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHallId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "HallId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "HallId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHallId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketFilterValidationError{
				field:  "HallId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetActivityId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "ActivityId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "ActivityId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActivityId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketFilterValidationError{
				field:  "ActivityId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUserId()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketFilterValidationError{
					field:  "UserId",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUserId()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketFilterValidationError{
				field:  "UserId",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return TicketFilterMultiError(errors)
	}

	return nil
}

// TicketFilterMultiError is an error wrapping multiple validation errors
// returned by TicketFilter.ValidateAll() if the designated constraints aren't met.
type TicketFilterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TicketFilterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TicketFilterMultiError) AllErrors() []error { return m }

// TicketFilterValidationError is the validation error returned by
// TicketFilter.Validate if the designated constraints aren't met.
type TicketFilterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TicketFilterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TicketFilterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TicketFilterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TicketFilterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TicketFilterValidationError) ErrorName() string { return "TicketFilterValidationError" }

// Error satisfies the builtin error interface
func (e TicketFilterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTicketFilter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TicketFilterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TicketFilterValidationError{}

// Validate checks the field values on ListTicketRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTicketRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTicketRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTicketRequestMultiError, or nil if none found.
func (m *ListTicketRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTicketRequest) validate(all bool) error {
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
				errors = append(errors, ListTicketRequestValidationError{
					field:  "Fields",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTicketRequestValidationError{
					field:  "Fields",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFields()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTicketRequestValidationError{
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
				errors = append(errors, ListTicketRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ListTicketRequestValidationError{
					field:  "Filter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFilter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ListTicketRequestValidationError{
				field:  "Filter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for AfterPageToken

	// no validation rules for BeforePageToken

	if len(errors) > 0 {
		return ListTicketRequestMultiError(errors)
	}

	return nil
}

// ListTicketRequestMultiError is an error wrapping multiple validation errors
// returned by ListTicketRequest.ValidateAll() if the designated constraints
// aren't met.
type ListTicketRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTicketRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTicketRequestMultiError) AllErrors() []error { return m }

// ListTicketRequestValidationError is the validation error returned by
// ListTicketRequest.Validate if the designated constraints aren't met.
type ListTicketRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTicketRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTicketRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTicketRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTicketRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTicketRequestValidationError) ErrorName() string {
	return "ListTicketRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListTicketRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTicketRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTicketRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTicketRequestValidationError{}

// Validate checks the field values on ListTicketReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListTicketReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListTicketReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListTicketReplyMultiError, or nil if none found.
func (m *ListTicketReply) ValidateAll() error {
	return m.validate(true)
}

func (m *ListTicketReply) validate(all bool) error {
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
					errors = append(errors, ListTicketReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListTicketReplyValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListTicketReplyValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.NextAfterPageToken != nil {
		// no validation rules for NextAfterPageToken
	}

	if m.NextBeforePageToken != nil {
		// no validation rules for NextBeforePageToken
	}

	if len(errors) > 0 {
		return ListTicketReplyMultiError(errors)
	}

	return nil
}

// ListTicketReplyMultiError is an error wrapping multiple validation errors
// returned by ListTicketReply.ValidateAll() if the designated constraints
// aren't met.
type ListTicketReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListTicketReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListTicketReplyMultiError) AllErrors() []error { return m }

// ListTicketReplyValidationError is the validation error returned by
// ListTicketReply.Validate if the designated constraints aren't met.
type ListTicketReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListTicketReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListTicketReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListTicketReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListTicketReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListTicketReplyValidationError) ErrorName() string { return "ListTicketReplyValidationError" }

// Error satisfies the builtin error interface
func (e ListTicketReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListTicketReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListTicketReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListTicketReplyValidationError{}

// Validate checks the field values on Ticket with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Ticket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ticket with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TicketMultiError, or nil if none found.
func (m *Ticket) ValidateAll() error {
	return m.validate(true)
}

func (m *Ticket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetLocation()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Location",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocation()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "Location",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetHall()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Hall",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Hall",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetHall()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "Hall",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetActivity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Activity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Activity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetActivity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "Activity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShow()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Show",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "Show",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShow()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "Show",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetShowSalesType()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "ShowSalesType",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "ShowSalesType",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetShowSalesType()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "ShowSalesType",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, TicketValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return TicketValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Status

	// no validation rules for OrderId

	// no validation rules for UserId

	// no validation rules for Notice

	if len(errors) > 0 {
		return TicketMultiError(errors)
	}

	return nil
}

// TicketMultiError is an error wrapping multiple validation errors returned by
// Ticket.ValidateAll() if the designated constraints aren't met.
type TicketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TicketMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TicketMultiError) AllErrors() []error { return m }

// TicketValidationError is the validation error returned by Ticket.Validate if
// the designated constraints aren't met.
type TicketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TicketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TicketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TicketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TicketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TicketValidationError) ErrorName() string { return "TicketValidationError" }

// Error satisfies the builtin error interface
func (e TicketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTicket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TicketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TicketValidationError{}
