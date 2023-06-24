package logic

import (
	"github.com/GGboya/bluebell/dao/mysql"
	"github.com/GGboya/bluebell/models"
	"github.com/GGboya/bluebell/pkg/snowflake"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个user实例，也就是结构体
	user := &models.User{
		userID, p.Username, p.Password,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}
