package request

type Register struct {
	Phone    string `json:"phone" example:"手机号"`
	Password string `json:"password" example:"密码"`
	Code     string `json:"code" example:"验证码"`
}
