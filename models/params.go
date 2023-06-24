package models

// 定义请求的参数结构体
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// 定义用户的结构体
type User struct {
	UserId   int64  `db:"userid"`
	Username string `db:"username"`
	Password string `db:"password"`
}
