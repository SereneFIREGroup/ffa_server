package errors

import (
	"reflect"
)

// 生成 AccessDenied 错误
func AccessDeniedError(model string, reason string) error {
	return errorWithModelFieldReason(AccessDenied, model, "", reason)
}

// 生成 AlreadyExist 错误
func AlreadyExistsError(model string, reason string) error {
	return errorWithModelFieldReason(AlreadyExists, model, "", reason)
}

// 生成 AuthFailure 错误
func AuthFailureError(reason string) error {
	return errorWithModelFieldReason(AuthFailure, "", "", reason)
}

// 生成 BadConfig 错误
func BadConfigError(key string, value interface{}) error {
	err := Errorf(BadConfig, "got invalid value '%v' for key '%s'", value, key)
	err.(*Err).SetLocation(1)
	return err
}

// 生成 BlockedDetail 错误
func BlockedDetailError(model, field, reason string) error {
	return errorWithModelFieldReason(Blocked, model, field, reason)
}

// 生成 NotFoundDetail 错误
func NotFoundDetailError(model, field, reason string) error {
	return errorWithModelFieldReason(NotFound, model, field, reason)
}

// 生成 Blocked 错误
func BlockedError(reason string) error {
	return errorWithModelFieldReason(Blocked, "", "", reason)
}

// 生成 ConstraintViolation 错误
func ConstraintViolationError(reason string) error {
	return errorWithModelFieldReason(ConstraintViolation, "", "", reason)
}

// 生成 CorruptedData 错误
func CorruptedDataError(msg string, modelAndField ...string) error {
	parts := append([]string{CorruptedData}, modelAndField...)
	err := Errorf(Code(parts...), msg)
	err.(*Err).SetLocation(1)
	return err
}

// 生成 InUse 错误
func InUseError(model string, reason string) error {
	return errorWithModelFieldReason(InUse, model, "", reason)
}

// 生成 InvalidEnum 错误
func InvalidEnumError(enum interface{}, modelAndField ...string) error {
	parts := append([]string{InvalidEnum}, modelAndField...)
	err := Errorf(Code(parts...), "got '%v'", enum)
	err.(*Err).SetLocation(1)
	return err
}

// 生成 InvalidParameter 错误
func InvalidParameterError(model string, field string, reason string) error {
	return errorWithModelFieldReason(InvalidParameter, model, field, reason)
}

func InvalidFileExtError(model, field string) error {
	return errorWithModelFieldReason(InvalidFileExt, model, field, "")
}

// 生成 LimitExceeded 错误
func LimitExceededError(model string) error {
	return errorWithModelFieldReason(LimitExceeded, model, "", "")
}

// 生成 LimitExceeded 错误
func LimitExceededReasonError(model string, reason string) error {
	return errorWithModelFieldReason(LimitExceeded, model, "", reason)
}

// 生成 Malformed 错误
func MalformedError(format string) error {
	values := make(map[string]interface{})
	values[formatKey] = format
	err := New(Malformed, format).(*Err)
	err.values = values
	err.SetLocation(1)
	return err
}

// 生成 MissingParameter 错误
func MissingParameterError(model string, field string) error {
	return errorWithModelFieldReason(MissingParameter, model, field, "")
}

// 生成 NotFound 错误
func NotFoundError(model string) error {
	return errorWithModelFieldReason(NotFound, model, "", "")
}

// 生成 Deleted 错误
func DeletedError(model string) error {
	return errorWithModelFieldReason(Deleted, model, "", "")
}

// 生成 SourceDeleted 错误
func SourceDeletedError(model string) error {
	return errorWithModelFieldReason(SourceDeleted, model, "", "")
}

// SmsError 生成 SmsSendFailure 错误, other 为外部错误
// 处理发送短信失败的情况
func SmsError(other error) error {
	return externalError(other, SmsSendFailure)
}

// 生成 OutLimit 错误
func OutLimitError(model string, field string) error {
	return errorWithModelFieldReason(OutLimit, model, field, "")
}

// 生成 TypeMismatch 错误
func TypeMismatchError(value interface{}, expectedTypes ...string) error {
	var actualType = "nil"
	if value != nil {
		actualType = reflect.ValueOf(value).Type().String()
	}
	err := Errorf(TypeMismatch, "expected %v, got [%s](%v)", expectedTypes, actualType, value)
	err.(*Err).SetLocation(1)
	return err
}

// 生成 Timeout 错误
func TimeoutError(action string) error {
	values := make(map[string]interface{})
	values[actionKey] = action
	err := New(Timeout, action).(*Err)
	err.values = values
	err.SetLocation(1)
	return err
}

// 生成 VerificationFailure 错误
func VerificationFailureError(reason string) error {
	return errorWithModelFieldReason(VerificationFailure, "", "", reason)
}

// 生成 AppStoreBotError 错误
func AppStoreBot(other error) error {
	return externalError(other, AppStoreBotError)
}

// 生成 SQLError 错误
func Sql(other error) error {
	return externalError(other, SQLError)
}

func StatusWithError(model string, reason string) error {
	return errorWithModelFieldReason(StatusError, model, "", reason)
}

func PluginLifeCycleError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceLifeCycleFailure, "", "", reason)
}

func PluginInstallError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceInstallationFailure, "", "", reason)
}

func PluginUninstallError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceUninstallationFailure, "", "", reason)
}

func PluginUploadError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceUploadFailure, "", "", reason)
}

func PluginEnableError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceEnableFailure, "", "", reason)
}

func PluginDisableError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceDisableFailure, "", "", reason)
}

func PluginUpgradeError(reason string) error {
	return errorWithModelFieldReason(PluginInstanceUpgradeFailure, "", "", reason)
}

func ConnectThirdServer(reason string) error {
	// connect third party service failure
	return errorWithModelFieldReason(ConnectThirdServerError, "", "", reason)
}

func PluginInvokeError(t string, model string, field string, reason string) error {
	return errorWithModelFieldReason(t, model, field, reason)
}

func PluginInvokeErrorWithDesc(t string, model string, field string, reason string, description string) error {
	return pluginErrorWithModelFieldReason(t, model, field, reason, description)
}

func externalError(other error, code string) error {
	if other == nil {
		return nil
	}
	err := Wrap(other, Errorf(code, other.Error()))
	err.(*Err).SetLocation(2)
	return err
}

func errorWithModelFieldReason(t string, model string, field string, reason string) error {
	parts := []string{t}
	values := make(map[string]interface{})
	if len(model) > 0 {
		parts = append(parts, model)
		values[modelKey] = model
	}
	if len(field) > 0 {
		parts = append(parts, field)
		values[fieldKey] = field
	}
	if len(reason) > 0 {
		parts = append(parts, reason)
		values[reasonKey] = reason
	}
	err := New(parts...).(*Err)
	err.values = values
	err.SetLocation(2)
	return err
}

func pluginErrorWithModelFieldReason(t string, model string, field string, reason string, description string) error {
	parts := []string{t}
	values := make(map[string]interface{})
	if len(model) > 0 {
		parts = append(parts, model)
		values[modelKey] = model
	}
	if len(field) > 0 {
		parts = append(parts, field)
		values[fieldKey] = field
	}
	if len(reason) > 0 {
		parts = append(parts, reason)
		values[reasonKey] = reason
	}
	if len(description) > 0 {
		parts = append(parts, description)
		values[descriptionKey] = description
	}
	err := New(parts...).(*Err)
	err.values = values
	err.SetLocation(2)
	return err
}
