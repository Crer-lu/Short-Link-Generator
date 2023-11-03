package dto

//error code

const (
	NoErr           = 0 // 正常
	ErrEmailExist   = 1 // 邮箱已被使用
	ErrCaptcha      = 2 // 验证码错误
	ErrPassword     = 3 // 密码错误
	ErrUserNotFound = 4 // 用户不存在

	ErrShortLinkExist = 5 // 短链已经被使用
	ErrNoShortLink    = 6 // 短链不存在
	ErrPrivilege      = 7 // 权限不足

	ErrShortLinkActive = 8 // 短链被中断使用
	ErrShortLinkTime   = 9 // 不在可用时间范围

	BadReqeust    = 400
	InternalError = 500
)
