package valobj

type LoginReq struct {
	NickName string `json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `json:"password" validate:"required,min=4,max=20" label:"密码"`
}
