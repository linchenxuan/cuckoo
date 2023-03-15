package valobj

type RegisterUserReq struct {
	Nickname string `json:"nickname" validate:"required,min=4,max=12"`
	Password string `json:"password" validate:"required,min=4,max=20"`
}

type UpdateUserReq struct {
	ID           int64  `json:"id"`
	Nickname     string `json:"nickname" validate:"required"`
	Avatar       string `json:"avatar"`
	Introduction string `json:"introduction"`
	Email        string `json:"email"`
}

type UserInfo struct {
	Id       int64  `json:"id"`
	NickName string `json:"nickName"`
}
