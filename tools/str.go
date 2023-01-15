// Package tools 提供一些可公用的func
// 例如生成随机用户名，订单金额处理等
package tools

import "math/rand"

// Letters 随机字符串；
var Letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// 生成指定长度的随机字符串 (不会按格式来换行，如果想换行，可在下方加一个// 空行)
//
// 注意： func必须是大写的，才会导出到文档；小写的func不会被导出(认为仅内部使用，不需要暴露)
func GenRandomName(length int) string {
	if length <= 0 {
		return ""
	}
	return genRandomName(length)

}

// genRandomName 小写的func不会被导出(认为仅内部使用，不需要暴露)
func genRandomName(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = Letters[rand.Intn(len(Letters))]
	}
	return string(b)
}
