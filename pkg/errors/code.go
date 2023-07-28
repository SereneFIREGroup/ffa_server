package errors

import "net/http"

// 默认映射到非 500 状态码的错误
const (
	// AccessDenied 用户没有权限访问某项资源，或者某项资源不存在
	// 格式：AccessDenied.MODEL[.REASON]
	// ex. AccessDenied.Project
	AccessDenied = "AccessDenied"

	// AlreadyExists 某项资源已经存在，因此无法添加
	// 格式: AlreadyExists.MODEL[.REASON]
	// ex. AlreadyExists.Task
	// ex. AlreadyExists.Field.NameConflict
	AlreadyExists = "AlreadyExists"

	// AuthFailure 用户授权失败，比如登录失败或者 token 过期
	// 格式：AuthFailure.REASON
	// ex. AuthFailure.InvalidPassword
	// ex. AuthFailure.InvalidToken
	AuthFailure = "AuthFailure"

	// Blocked 用户被封禁
	// 格式：Blocked[.REASON]
	// ex. Blocked.IPBanned
	Blocked = "Blocked"

	// ConstraintViolation 未满足指定的约束条件，因此无法执行操作，正常情况下应该使用更具体的错误，比如 PermissionDenied、
	// AccessDenied、AlreadyExists、InUse，只有那些难以归类的约束条件才能使用这个错误码
	// 格式：ConstraintViolation.REASON
	// ex. ConstraintViolation.JoinOtherOrganization
	ConstraintViolation = "ConstraintViolation"

	// Deleted 某项资源被删除
	// 格式：Deleted.MODEL[.REASON]
	// ex. Deleted.Task
	Deleted = "Deleted"

	// SourceDeleted 某项资源的源资源被删除
	// 格式：SourceDeleted.MODEL[.REASON]
	// ex. SourceDeleted.Task
	SourceDeleted = "SourceDeleted"

	// InUse 某项资源正在使用中，因此无法被删除
	// 格式：InUse.MODEL[.REASON]
	// ex. InUse.TaskStatus.UsedInTransitions
	InUse = "InUse"

	// InvalidParameter 客户端传入的参数不合法，主要是指格式上的错误，非格式错误请尽量使用其它返回码
	// 格式：InvalidParameter[.MODEL].FIELD[.REASON]
	// 格式注：只有确实没有对应 MODEL 的情况才能省略 MODEL，其它任何情况都应该加上 MODEL
	// ex. InvalidParameter.Task.Summary
	// ex. InvalidParameter.Task.Summary.TooLong
	// ex. InvalidParameter.Limit
	// ex. InvalidParameter.Limit.OutOfRange
	InvalidParameter = "InvalidParameter"

	// InvalidFileExt 用户上传的文件后缀不合法
	// 格式：InvalidFileExt.[.MODEL].FIELD
	// ex. InvalidFileExt.Resouce.Name
	InvalidFileExt = "InvalidFileExt"

	// LimitExceeded 资源的使用超出了限额
	// 格式：LimitExceeded.MODEL
	// ex. LimitExceeded.TeamMember
	LimitExceeded = "LimitExceeded"

	// Malformed 数据格式不正确，解析失败
	// 格式：Malformed.FORMAT
	// ex. Malformed.JSON
	// ex. Malformed.XML
	Malformed = "Malformed"

	// MissingParameter 客户端没有传入某个必填参数
	// 格式：MissingParameter[.MODEL].FIELD
	// 格式注：只有确实没有对应 MODEL 的情况才能省略 MODEL，其它任何情况都应该加上 MODEL
	// ex. MissingParameter.Task.Summary
	// ex. MissingParameter.Limit
	MissingParameter = "MissingParameter"

	// NotFound 某项资源不存在
	// 格式：NotFound.MODEL[.REASON]
	// ex. NotFound.Task
	NotFound = "NotFound"

	// OK 只能用于 ErrPayload 返回给客户端表示请求成功，其它任何时候都应该返回 nil，而不应该使用这个值
	// 格式：OK
	OK = "OK"

	// OutLimit 超出限制
	OutLimit = "OutLimit"

	// PermissionDenied 用户没有某项权限，因此无法执行操作
	// 格式：PermissionDenied.PERMISSION
	// 格式注：PERMISSION 就是把权限标签从下划线分隔改成驼峰式
	// ex. PermissionDenied.AddTask
	PermissionDenied = "PermissionDenied"

	// SmsSendFailure 短信发送失败
	SmsSendFailure = "SmsSendFailure"

	// Timeout 操作超时
	// 格式：Timeout.ACTION
	// ex. Timeout.GlobalSearch
	// ex. Timeout.CopyProject
	Timeout = "Timeout"

	// VerificationFailure 校验失败，比如校验码无效或者过期
	// 格式：VerificationFailure.REASON
	// ex. VerificationFailure.InvalidCode
	// ex. VerificationFailure.CodeExpired
	VerificationFailure = "VerificationFailure"

	// PluginInstanceLifeCycleFailure 实例安装失败
	// InstanceInstallationFailure.REASON
	// ex. InstanceInstallationFailure.FileTooLarge
	PluginInstanceLifeCycleFailure = "PluginInstanceLifeCycleFailure"

	// PluginInstanceInstallationFailure 实例安装失败
	// InstanceInstallationFailure.REASON
	// ex. InstanceInstallationFailure.FileTooLarge
	PluginInstanceInstallationFailure = "PluginInstanceInstallationFailure"

	// PluginInstanceUninstallationFailure 实例卸载安装失败
	// InstanceInstallationFailure.REASON
	// ex. InstanceInstallationFailure.FileTooLarge
	PluginInstanceUninstallationFailure = "PluginInstanceUninstallationFailure"

	// PluginInstanceUploadFailure 实例上传失败
	// PluginInstanceUploadFailure.REASON
	// ex. InstanceInstallationFailure.FileTooLarge
	PluginInstanceUploadFailure = "PluginInstanceUploadFailure"

	// PluginInstanceEnableFailure 实例启动失败
	// PluginUninstanceStartError.REASON
	// ex. PluginUninstanceStartError.FileTooLarge
	PluginInstanceEnableFailure = "PluginInstanceEnableFailure"

	// PluginInstanceDisableFailure 实例停用失败
	// PluginUninstanceStopFailure.REASON
	// ex. PluginUninstanceStopFailure.FileTooLarge
	PluginInstanceDisableFailure = "PluginInstanceDisableFailure"

	// PluginInstanceUpgradeFailure 实例升级失败
	// PluginUninstanceUpgradeFailure.REASON
	// ex. PluginUninstanceUpgradeFailure.FileTooLarge
	PluginInstanceUpgradeFailure = "PluginInstanceUpgradeFailure"

	// ConnectThirdServerError ConnectThirdError
	// tmpl: ConnectThirdPartyError.REASON
	// ex. ConnectThirdPartyError.DNSError
	ConnectThirdServerError = "ConnectThirdServerError"
)

// 默认映射到 500 状态码的错误
const (
	// AppStoreBotError 由 AppStore 爬虫产生的错误，用 AppStoreBot 方法生成，不要直接使用
	// 格式：AppStoreBotError
	// ex. return errors.AppStoreBot(err)
	AppStoreBotError = "AppStoreBotError"

	// BadConfig 配置文件错误，用 BadConfigError 方法生成，不要直接使用
	// 格式：BadConfig
	// ex. return errors.BadConfigError(key, value)
	BadConfig = "BadConfig"

	// CorruptedData 脏数据，即数据理应满足某项要求，但实际却没有满足
	// 格式：CorruptedData[.MODEL][.FIELD][.REASON]
	// ex. CorruptedData.Field.DefaultValue
	CorruptedData = "CorruptedData"

	// InvalidEnum 无效的枚举值。可能为客户端错误或者服务端错误，默认为服务端错误，作为客户端错误时需要用 InvalidParameter 包起来
	// 格式：InvalidEnum[.MODEL].FIELD
	// 格式注：只有确实没有对应 MODEL 的情况才能省略 MODEL，其它任何情况都应该加上 MODEL
	// ex. return errors.InvalidEnumError(enum, errors.PermissionRule, errors.Permission)
	InvalidEnum = "InvalidEnum"

	// KeyConflict 主键冲突，类似于 AlreadyExist，但专门用于服务端错误
	// 格式：KeyConflict[.MODEL][.FIELD]
	// 格式注：MODEL 和 FIELD 至少应该指定一个
	// ex. KeyConflict.Context.Type
	KeyConflict = "KeyConflict"

	// ServerError 服务器内部错误，只能用于包裹其它不合适直接返回给客户端的服务器错误，其它任何时候都不应该使用
	// 格式：ServerError
	ServerError = "ServerError"

	// SQLError 由 MySQL 产生的错误，用 Sql 方法生成，不要直接使用
	// 格式：SQLError
	// ex. return errors.Sql(err)
	SQLError = "SQLError"

	// StatusError 状态错误
	// 格式 StatusError
	// ex. return errors.StatusError(Migration, reason)
	StatusError = "StatusError"

	// UnknownError 未知错误，用来表示那些没有标注错误码的错误，任何时候都不应该主动使用
	// 格式：UnknownError
	UnknownError = "UnknownError"

	// TypeMismatch 变量类型不符合预期要求，用 TypeMismatchError 方法生成，不要直接使用
	// 格式：TypeMismatch
	// ex:
	// s, ok := v.(string)
	// if !ok {
	//     return errors.TypeMismatchError(v, "string")
	// }
	TypeMismatch = "TypeMismatch"
)

const (
	JSON = "JSON"
)

// model
const (
	Family      = "Family"
	FourMoney   = "FourMoney"
	PocketMoney = "PocketMoney"
	User        = "User"
	Earning     = "Earning"
)

// field
const (
	Amount     = "Amount"
	Category   = "Category"
	Desc       = "Desc"
	EndDate    = "EndDate"
	Email      = "Email"
	FIREGold   = "FIREGold"
	ID         = "ID"
	Owner      = "Owner"
	Password   = "Password"
	Phone      = "Phone"
	StartDate  = "StartDate"
	Token      = "Token"
	Type       = "Type"
	UserName   = "UserName"
	ULID       = "ULID"
	VerifyCode = "VerifyCode"
)

// reason
const (
	AlreadyExist      = "AlreadyExist"
	Empty             = "Empty"
	IncorrectPassword = "IncorrectPassword"
	InvalidFormat     = "InvalidFormat"
	InvalidToken      = "InvalidToken"
	InvalidOwner      = "InvalidOwner"
)

var (
	DefaultStatusCodeBinding = map[string]int{
		AccessDenied:                        http.StatusForbidden,
		AlreadyExists:                       http.StatusConflict,
		StatusError:                         http.StatusConflict,
		AppStoreBotError:                    http.StatusInternalServerError,
		AuthFailure:                         http.StatusUnauthorized,
		BadConfig:                           http.StatusInternalServerError,
		Blocked:                             http.StatusForbidden,
		ConstraintViolation:                 http.StatusForbidden,
		CorruptedData:                       http.StatusInternalServerError,
		InUse:                               http.StatusBadRequest,
		InvalidEnum:                         http.StatusInternalServerError,
		InvalidParameter:                    http.StatusBadRequest,
		KeyConflict:                         http.StatusInternalServerError,
		LimitExceeded:                       http.StatusForbidden,
		Malformed:                           http.StatusBadRequest,
		MissingParameter:                    http.StatusBadRequest,
		NotFound:                            http.StatusNotFound,
		Deleted:                             http.StatusInternalServerError,
		SourceDeleted:                       http.StatusInternalServerError,
		OK:                                  http.StatusOK,
		OutLimit:                            http.StatusInternalServerError,
		PermissionDenied:                    http.StatusForbidden,
		ServerError:                         http.StatusInternalServerError,
		SQLError:                            http.StatusInternalServerError,
		Timeout:                             http.StatusBadRequest,
		UnknownError:                        http.StatusInternalServerError,
		TypeMismatch:                        http.StatusInternalServerError,
		VerificationFailure:                 http.StatusBadRequest,
		InvalidFileExt:                      http.StatusBadRequest,
		PluginInstanceInstallationFailure:   http.StatusBadRequest,
		PluginInstanceUninstallationFailure: http.StatusBadRequest,
		PluginInstanceUploadFailure:         http.StatusBadRequest,
		PluginInstanceEnableFailure:         http.StatusBadRequest,
		PluginInstanceDisableFailure:        http.StatusBadRequest,
		ConnectThirdServerError:             http.StatusBadRequest,
	}
)
