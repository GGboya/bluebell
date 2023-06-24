package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/GGboya/bluebell/models"
)

// 把每一步数据库操作封装成函数
// 等待logic层根据业务需求调用

const secret = "ggcoding.top"

func CheckUserExist(username string) (err error) {
	// 如果存在于mysql中，则用户重复
	sqlStr := "select count(user_id) from user where username = ?"
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("用户已经存在")
	}
	return err
}

func InsertUser(user *models.User) error {
	// 把user插入数据库
	// 原生sql语句
	sqlStr := `insert into user(user_id, username, password) values(?,?,?)`
	password := encryptPassword(user.Password)
	_, err := db.Exec(sqlStr, user.UserId, user.Username, password)
	return err
}

func encryptPassword(oPassword string) string {
	// 创建 MD5 hash 对象并写入字节数组
	h := md5.New()
	h.Write([]byte(secret)) // 加盐的操作，使得更安全
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
