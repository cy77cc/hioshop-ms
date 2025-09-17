package xcode

// Xcode 定义响应码类型
type Xcode uint32

// 成功类 (1000~1999)
const (
	Success       Xcode = 1000 // 请求成功
	CreateSuccess Xcode = 1001 // 创建成功
	DeleteSuccess Xcode = 1002 // 删除成功
	UpdateSuccess Xcode = 1003 // 更新成功
)

// 客户端错误 (2000~2999)
const (
	ParamError     Xcode = 2000 // 参数错误
	MissingParam   Xcode = 2001 // 缺少必要参数
	MethodNotAllow Xcode = 2002 // 请求方法不支持
	Unauthorized   Xcode = 2003 // 未认证（需要登录）
	Forbidden      Xcode = 2004 // 无权限
	NotFound       Xcode = 2005 // 资源不存在
)

// 服务端错误 (3000~3999)
const (
	ServerError     Xcode = 3000 // 服务器内部错误
	DatabaseError   Xcode = 3001 // 数据库错误
	CacheError      Xcode = 3002 // 缓存服务错误
	ExternalAPIFail Xcode = 3003 // 外部服务调用失败
	TimeoutError    Xcode = 3004 // 超时
)

// 业务错误 (4000~4999)
const (
	FileUploadFail   Xcode = 4000 // 文件上传失败
	FileTypeInvalid  Xcode = 4001 // 文件格式不支持
	UserAlreadyExist Xcode = 4002 // 用户已存在
	UserNotExist     Xcode = 4003 // 用户不存在
	PasswordError    Xcode = 4004 // 密码错误
	TokenExpired     Xcode = 4005 // Token 过期
)

// Msg 根据 Xcode 返回对应的提示信息
func (c Xcode) Msg() string {
	switch c {
	case Success:
		return "请求成功"
	case CreateSuccess:
		return "创建成功"
	case DeleteSuccess:
		return "删除成功"
	case UpdateSuccess:
		return "更新成功"

	case ParamError:
		return "参数错误"
	case MissingParam:
		return "缺少必要参数"
	case MethodNotAllow:
		return "请求方法不支持"
	case Unauthorized:
		return "未认证"
	case Forbidden:
		return "无权限"
	case NotFound:
		return "资源不存在"

	case ServerError:
		return "服务器内部错误"
	case DatabaseError:
		return "数据库错误"
	case CacheError:
		return "缓存服务错误"
	case ExternalAPIFail:
		return "外部服务调用失败"
	case TimeoutError:
		return "请求超时"

	case FileUploadFail:
		return "文件上传失败"
	case FileTypeInvalid:
		return "文件格式不支持"
	case UserAlreadyExist:
		return "用户已存在"
	case UserNotExist:
		return "用户不存在"
	case PasswordError:
		return "密码错误"
	case TokenExpired:
		return "Token 已过期"
	default:
		return "未知错误"
	}
}
