// Package srv 主要是业务相关的代码
// 例如用户，订单等的处理。但因为order.go和user.go都会放在Package files内容下，所以此处的注释粒度是包维度，而不是文件级别
//
// 更多信息可参考这个链接：https://xxxx.com
package srv

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	/*
		仅大写的全局变量会被导出到文档
		这段内容想要显示，不能与下段可被导出的内容间有空行
	*/
	// 时间格式；
	TimeLayOut = time.Layout

	// salt 小写的变量不会被导出
	salt = "盐值"
)

type GenderEnum int8

const (
	// 男性用户
	Male GenderEnum = 1
	// 女性用户
	FeMale GenderEnum = 2
)

// 只有大写开头的会被导出到文档
const CountryCode = "+86"

// 不会导出到文档
const defaultName = "张三"

// User 用户相关信息
type User struct {
	// 用户名称,如果用户未设置则随机生成
	Name string
	// 年龄
	Age int
	// 地址
	Address string

	// 密码
	// 注意:存数据库时需要对用户密码进行哈希加盐
	Password string

	Gender GenderEnum

	IsVip bool // 这样写注释也会被原样解析到文档;与上下两个字段之间的空行也会被原样解析到文档

	/*
		也支持这样的注释方式
	*/
	Phone string // 用Phone作为唯一标识
}

// 对用户密码进行加密
//
// 对 盐+原始密码 这个字符串进行bcrypt操作；
// TODO：对应的校验密码的方法待开发
func EncryptPassword(password string) string {
	hash := hashAndSalt([]byte(salt + password))
	return hash

}

// 实际处理； 不会导出到文档
func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
