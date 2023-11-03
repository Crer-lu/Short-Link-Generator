package dto

type UserRegisterReq struct {
	Email    string `json:"email"`    // 用户邮箱
	Name     string `json:"name"`     // 用户名
	Password string `json:"password"` // 用户密码
}

type UserRegisterResp struct{}

type UserCaptchaReq struct{}

type UserCaptchaResp struct {
	CAPTCHAID  string `json:"captcha_id"`  // 验证码 ID
	CAPTCHAURL string `json:"captcha_url"` // 验证码图片地址
}

type UserLogoutReq struct{}

type UserLogoutResp struct{}

type UserLoginReq struct {
	CAPTCHAID    string `json:"captcha_id"`
	CAPTCHAValue string `json:"captcha_value"`
	Email        string `json:"email"`    // 用户邮箱
	Password     string `json:"password"` // 用户密码
}

type UserLoginResp struct{}

type UserGetInfoReq struct{}

type UserGetInfoResp struct {
	Email string `json:"email"` // 用户邮箱
	ID    int64  `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

type UserUpdateInfoReq struct {
	Email string `json:"email"` // 用户邮箱
	ID    int64  `json:"id"`    // 用户ID
	Name  string `json:"name"`  // 用户名
}

type UserUpdateInfoResp struct{}

type UserUpdatePasswdReq struct {
	NewPwd string `json:"new_pwd"` // 新密码
	OldPwd string `json:"old_pwd"` // 旧密码
}

type UserUpdatePasswdResp struct{}
