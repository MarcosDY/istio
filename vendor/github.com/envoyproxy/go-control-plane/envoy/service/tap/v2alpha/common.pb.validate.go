// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/tap/v2alpha/common.proto

package envoy_service_tap_v2alpha

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on TapConfig with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *TapConfig) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatchConfig() == nil {
		return TapConfigValidationError{
			field:  "MatchConfig",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetMatchConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TapConfigValidationError{
					field:  "MatchConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if m.GetOutputConfig() == nil {
		return TapConfigValidationError{
			field:  "OutputConfig",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetOutputConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TapConfigValidationError{
					field:  "OutputConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTapEnabled()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return TapConfigValidationError{
					field:  "TapEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// TapConfigValidationError is the validation error returned by
// TapConfig.Validate if the designated constraints aren't met.
type TapConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TapConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TapConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TapConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TapConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TapConfigValidationError) ErrorName() string { return "TapConfigValidationError" }

// Error satisfies the builtin error interface
func (e TapConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTapConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TapConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TapConfigValidationError{}

// Validate checks the field values on MatchPredicate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MatchPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Rule.(type) {

	case *MatchPredicate_OrMatch:

		{
			tmp := m.GetOrMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "OrMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_AndMatch:

		{
			tmp := m.GetAndMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "AndMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_NotMatch:

		{
			tmp := m.GetNotMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "NotMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_AnyMatch:

		if m.GetAnyMatch() != true {
			return MatchPredicateValidationError{
				field:  "AnyMatch",
				reason: "value must equal true",
			}
		}

	case *MatchPredicate_HttpRequestHeadersMatch:

		{
			tmp := m.GetHttpRequestHeadersMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "HttpRequestHeadersMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_HttpRequestTrailersMatch:

		{
			tmp := m.GetHttpRequestTrailersMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "HttpRequestTrailersMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_HttpResponseHeadersMatch:

		{
			tmp := m.GetHttpResponseHeadersMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "HttpResponseHeadersMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *MatchPredicate_HttpResponseTrailersMatch:

		{
			tmp := m.GetHttpResponseTrailersMatch()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicateValidationError{
						field:  "HttpResponseTrailersMatch",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return MatchPredicateValidationError{
			field:  "Rule",
			reason: "value is required",
		}

	}

	return nil
}

// MatchPredicateValidationError is the validation error returned by
// MatchPredicate.Validate if the designated constraints aren't met.
type MatchPredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicateValidationError) ErrorName() string { return "MatchPredicateValidationError" }

// Error satisfies the builtin error interface
func (e MatchPredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicateValidationError{}

// Validate checks the field values on HttpHeadersMatch with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *HttpHeadersMatch) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetHeaders() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HttpHeadersMatchValidationError{
						field:  fmt.Sprintf("Headers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// HttpHeadersMatchValidationError is the validation error returned by
// HttpHeadersMatch.Validate if the designated constraints aren't met.
type HttpHeadersMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HttpHeadersMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HttpHeadersMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HttpHeadersMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HttpHeadersMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HttpHeadersMatchValidationError) ErrorName() string { return "HttpHeadersMatchValidationError" }

// Error satisfies the builtin error interface
func (e HttpHeadersMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHttpHeadersMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HttpHeadersMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HttpHeadersMatchValidationError{}

// Validate checks the field values on OutputConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *OutputConfig) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetSinks()) != 1 {
		return OutputConfigValidationError{
			field:  "Sinks",
			reason: "value must contain exactly 1 item(s)",
		}
	}

	for idx, item := range m.GetSinks() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return OutputConfigValidationError{
						field:  fmt.Sprintf("Sinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetMaxBufferedRxBytes()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutputConfigValidationError{
					field:  "MaxBufferedRxBytes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMaxBufferedTxBytes()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return OutputConfigValidationError{
					field:  "MaxBufferedTxBytes",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for Streaming

	return nil
}

// OutputConfigValidationError is the validation error returned by
// OutputConfig.Validate if the designated constraints aren't met.
type OutputConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputConfigValidationError) ErrorName() string { return "OutputConfigValidationError" }

// Error satisfies the builtin error interface
func (e OutputConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputConfigValidationError{}

// Validate checks the field values on OutputSink with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *OutputSink) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := OutputSink_Format_name[int32(m.GetFormat())]; !ok {
		return OutputSinkValidationError{
			field:  "Format",
			reason: "value must be one of the defined enum values",
		}
	}

	switch m.OutputSinkType.(type) {

	case *OutputSink_StreamingAdmin:

		{
			tmp := m.GetStreamingAdmin()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return OutputSinkValidationError{
						field:  "StreamingAdmin",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *OutputSink_FilePerTap:

		{
			tmp := m.GetFilePerTap()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return OutputSinkValidationError{
						field:  "FilePerTap",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *OutputSink_StreamingGrpc:

		{
			tmp := m.GetStreamingGrpc()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return OutputSinkValidationError{
						field:  "StreamingGrpc",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return OutputSinkValidationError{
			field:  "OutputSinkType",
			reason: "value is required",
		}

	}

	return nil
}

// OutputSinkValidationError is the validation error returned by
// OutputSink.Validate if the designated constraints aren't met.
type OutputSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputSinkValidationError) ErrorName() string { return "OutputSinkValidationError" }

// Error satisfies the builtin error interface
func (e OutputSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputSinkValidationError{}

// Validate checks the field values on StreamingAdminSink with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *StreamingAdminSink) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// StreamingAdminSinkValidationError is the validation error returned by
// StreamingAdminSink.Validate if the designated constraints aren't met.
type StreamingAdminSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamingAdminSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamingAdminSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamingAdminSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamingAdminSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamingAdminSinkValidationError) ErrorName() string {
	return "StreamingAdminSinkValidationError"
}

// Error satisfies the builtin error interface
func (e StreamingAdminSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamingAdminSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamingAdminSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamingAdminSinkValidationError{}

// Validate checks the field values on FilePerTapSink with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilePerTapSink) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetPathPrefix()) < 1 {
		return FilePerTapSinkValidationError{
			field:  "PathPrefix",
			reason: "value length must be at least 1 bytes",
		}
	}

	return nil
}

// FilePerTapSinkValidationError is the validation error returned by
// FilePerTapSink.Validate if the designated constraints aren't met.
type FilePerTapSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilePerTapSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilePerTapSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilePerTapSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilePerTapSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilePerTapSinkValidationError) ErrorName() string { return "FilePerTapSinkValidationError" }

// Error satisfies the builtin error interface
func (e FilePerTapSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilePerTapSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilePerTapSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilePerTapSinkValidationError{}

// Validate checks the field values on StreamingGrpcSink with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *StreamingGrpcSink) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for TapId

	if m.GetGrpcService() == nil {
		return StreamingGrpcSinkValidationError{
			field:  "GrpcService",
			reason: "value is required",
		}
	}

	{
		tmp := m.GetGrpcService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return StreamingGrpcSinkValidationError{
					field:  "GrpcService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// StreamingGrpcSinkValidationError is the validation error returned by
// StreamingGrpcSink.Validate if the designated constraints aren't met.
type StreamingGrpcSinkValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StreamingGrpcSinkValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StreamingGrpcSinkValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StreamingGrpcSinkValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StreamingGrpcSinkValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StreamingGrpcSinkValidationError) ErrorName() string {
	return "StreamingGrpcSinkValidationError"
}

// Error satisfies the builtin error interface
func (e StreamingGrpcSinkValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStreamingGrpcSink.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StreamingGrpcSinkValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StreamingGrpcSinkValidationError{}

// Validate checks the field values on MatchPredicate_MatchSet with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_MatchSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 2 {
		return MatchPredicate_MatchSetValidationError{
			field:  "Rules",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return MatchPredicate_MatchSetValidationError{
						field:  fmt.Sprintf("Rules[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// MatchPredicate_MatchSetValidationError is the validation error returned by
// MatchPredicate_MatchSet.Validate if the designated constraints aren't met.
type MatchPredicate_MatchSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicate_MatchSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicate_MatchSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicate_MatchSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicate_MatchSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicate_MatchSetValidationError) ErrorName() string {
	return "MatchPredicate_MatchSetValidationError"
}

// Error satisfies the builtin error interface
func (e MatchPredicate_MatchSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_MatchSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicate_MatchSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicate_MatchSetValidationError{}
