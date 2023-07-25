package errors

import (
	"net/http"
	"strings"

	jujuerrors "github.com/juju/errors"
)

const (
	httpStatusKey  = "http_status"
	modelKey       = "model"
	fieldKey       = "field"
	reasonKey      = "reason"
	permissionKey  = "permission"
	formatKey      = "format"
	readableKey    = "readable"
	actionKey      = "action"
	descriptionKey = "description"
)

type Err struct {
	*jujuerrors.Err
	Code   string
	values map[string]interface{}
}

// SetValue 设置自定义数据，设置的数据会返回给客户端
func (e *Err) SetValue(key string, value interface{}) {
	if e == nil {
		return
	}
	if e.values == nil {
		e.values = make(map[string]interface{})
	}
	e.values[key] = value
}

// Value 获取自定义数据
func (e *Err) Value(key string) (interface{}, bool) {
	if e == nil || e.values == nil {
		return nil, false
	}
	v, ok := e.values[key]
	return v, ok
}

// IntValue 获取 int 类型的自定义数据
func (e *Err) IntValue(key string) (int, bool) {
	v, ok := e.Value(key)
	if !ok {
		return 0, false
	}
	iv, ok := v.(int)
	return iv, ok
}

// StringValue 获取 string 类型的自定义数据
func (e *Err) StringValue(key string) (string, bool) {
	v, ok := e.Value(key)
	if !ok {
		return "", false
	}
	sv, ok := v.(string)
	return sv, ok
}

// Values 获取所有自定义数据
func (e *Err) Values() map[string]interface{} {
	if e == nil {
		return nil
	}
	if e.values == nil {
		e.values = make(map[string]interface{})
	}
	return e.values
}

// HttpStatus 获取错误对应的 http 状态码
func (e *Err) HttpStatus() int {
	if e == nil {
		return http.StatusOK
	}
	status, ok := e.IntValue(httpStatusKey)
	if ok {
		return status
	}
	errtype := strings.Split(e.Code, ".")[0]
	status = DefaultStatusCodeBinding[errtype]
	if status == 0 {
		status = http.StatusInternalServerError
	}
	return status
}

// New 根据错误码生成一个错误
// ex. err := errors.New("InvalidParameter", "Task.Summary", "TooLong")
// ex. err := errors.New("InvalidParameter.Task.Summary.TooLong")
func New(codeparts ...string) error {
	code := mustCode(codeparts)
	jerr := jujuerrors.NewErr("<" + code + ">")
	jerr.SetLocation(1)
	return &Err{
		Code: code,
		Err:  &jerr,
	}
}

// Errorf 根据错误码生成一个错误，并加入自定义信息
func Errorf(code string, format string, args ...interface{}) error {
	jerr := jujuerrors.NewErr(prefixWithCode(format, code), args...)
	jerr.SetLocation(1)
	return &Err{
		Code: code,
		Err:  &jerr,
	}
}

// Code 拼接错误码
// ex. errors.Code("InvalidParameter", "Task", "Summary", "TooLong") == "InvalidParameter.Task.Summary.TooLong"
// ex. errors.Code("InvalidParameter", "Task.Summary", "TooLong") == "InvalidParameter.Task.Summary.TooLong"
func Code(parts ...string) string {
	return strings.Join(parts, ".")
}

// Trace 从里向外原样返回错误时，必须调用这个方法，以记录里层错误的栈信息
// ex:
//
//	if err := SomeFunc(); err != nil {
//	    return errors.Trace(err)
//	}
func Trace(other error) error {
	if other == nil {
		return nil
	}
	newerr := new(Err)
	if err, ok := other.(*Err); ok {
		newerr.Code = err.Code
		newerr.values = err.values
	} else {
		newerr.Code = UnknownError
	}
	newerr.Err = jujuerrors.Trace(other).(*jujuerrors.Err)
	newerr.SetLocation(1)
	return newerr
}

// Wrap 在现有错误的基础上包装一层错误
// ex:
// err := tx.Exec(sql, args...)
//
//	if err != nil {
//	    return errors.Wrap(err, errors.New(errors.SQLError))
//	}
func Wrap(other error, err error) error {
	var code string
	var values map[string]interface{}
	var jerr *jujuerrors.Err
	if oneserr, ok := err.(*Err); ok {
		code = oneserr.Code
		values = oneserr.values
		jerr = jujuerrors.Wrap(other, oneserr.Err).(*jujuerrors.Err)
	} else {
		code = UnknownError
		prefixed := jujuerrors.NewErr(prefixWithCode(err.Error(), code))
		jerr = jujuerrors.Wrap(other, &prefixed).(*jujuerrors.Err)
	}
	jerr.SetLocation(1)
	return &Err{
		Code:   code,
		Err:    jerr,
		values: values,
	}
}

// Wrapf 在现有错误的基础上包装一层错误，并加入自定义信息
func Wrapf(other error, code string, format string, args ...interface{}) error {
	err := Errorf(code, format, args...)
	jerr := jujuerrors.Wrap(other, err).(*jujuerrors.Err)
	jerr.SetLocation(1)
	return &Err{
		Code: code,
		Err:  jerr,
	}
}

// WithStatus 额外指定 error 对应的 http 状态码，主要用于兼容以前的 APIResult，正常情况下不应该调用这个方法，
// 而是尽可能通过 DefaultStatusCodeBinding 进行自动映射
func WithStatus(err error, statusCode int) error {
	return WithValue(err, httpStatusKey, statusCode)
}

// WithDesc 额外指定 error 的描述信息
func WithDesc(err error, format string, args ...interface{}) error {
	oneserr, ok := err.(*Err)
	if ok {
		oneserr.Err = jujuerrors.Annotatef(oneserr.Err,
			prefixWithCode(format, oneserr.Code), args...).(*jujuerrors.Err)
	} else {
		oneserr = Wrapf(err, UnknownError, format, args...).(*Err)
	}
	return oneserr
}

// WithValue 额外指定 error 的自定义字段
func WithValue(err error, key string, value interface{}) error {
	oneserr, ok := err.(*Err)
	if !ok {
		var newerr error
		if err != nil {
			newerr = Errorf(UnknownError, err.Error())
		} else {
			newerr = New(UnknownError)
		}
		oneserr = Wrap(err, newerr).(*Err)
	}
	oneserr.SetValue(key, value)
	return oneserr
}

// GetValue 获取 error 的自定义字段
func GetValue(err error, key string) (interface{}, bool) {
	oneserr, ok := err.(*Err)
	if !ok {
		return nil, false
	}
	return oneserr.Value(key)
}

// WithValues 批量额外指定 error 的自定义字段
func WithValues(err error, values map[string]interface{}) error {
	if values == nil {
		return err
	}
	oneserr, ok := err.(*Err)
	if !ok {
		var newerr error
		if err != nil {
			newerr = Errorf(UnknownError, err.Error())
		} else {
			newerr = New(UnknownError)
		}
		oneserr = Wrap(err, newerr).(*Err)
	}
	for k, v := range values {
		oneserr.SetValue(k, v)
	}
	return oneserr
}

// HttpStatus 获取错误对应的 http 状态码
func HttpStatus(err error) int {
	if err == nil {
		return http.StatusOK
	}
	if oneserr, ok := err.(*Err); ok {
		return oneserr.HttpStatus()
	} else {
		return http.StatusInternalServerError
	}
}

// SetLocation 设置错误的调用层级
func SetLocation(err error, callDepth int) error {
	oneserr, ok := err.(*Err)
	if !ok {
		var newerr error
		if err != nil {
			newerr = Errorf(UnknownError, err.Error())
		} else {
			newerr = New(UnknownError)
		}
		oneserr = newerr.(*Err)
	}
	oneserr.SetLocation(callDepth + 1)
	return oneserr
}

// Cause returns the cause of the given error.  This will be either the
// original error, or the result of a Wrap or Mask call.
//
// Cause is the usual way to diagnose errors that may have been wrapped by
// the other errors functions.
func Cause(err error) error {
	return jujuerrors.Cause(err)
}

// Details returns information about the stack of errors wrapped by err, in
// the format:
//
//	[{filename:99: error one} {otherfile:55: cause of error one}]
//
// This is a terse alternative to ErrorStack as it returns a single line.
func Details(err error) string {
	return jujuerrors.Details(err)
}

// ErrorStack returns a string representation of the annotated error. If the
// error passed as the parameter is not an annotated error, the result is
// simply the result of the Error() method on that error.
//
// If the error is an annotated error, a multi-line string is returned where
// each line represents one entry in the annotation stack. The full filename
// from the call stack is used in the output.
//
//	first error
//	github.com/juju/errors/annotation_test.go:193:
//	github.com/juju/errors/annotation_test.go:194: annotation
//	github.com/juju/errors/annotation_test.go:195:
//	github.com/juju/errors/annotation_test.go:196: more context
//	github.com/juju/errors/annotation_test.go:197:
func ErrorStack(err error) string {
	return jujuerrors.ErrorStack(err)
}

func prefixWithCode(s string, code string) string {
	return "<" + code + "> " + s
}

func mustCode(parts []string) string {
	if len(parts) == 0 {
		panic("error code parts is empty")
	}
	return strings.Join(parts, ".")
}
