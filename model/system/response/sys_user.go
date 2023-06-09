package response

type SysUserResponse struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
}

type LoginResponse struct {
	User      SysUserResponse `json:"user"`
	Token     string          `json:"token"`
	ExpiresAt int64           `json:"expires_at"`
}
