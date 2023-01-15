package srv

type Order struct {
	// 订单ID
	ID int64
	// 订单类型
	BizTyp int8
	// 订单金额
	Number float64
	// 发生订单的手机号，与用户关联
	Phone string
}

// 判断订单类型是否合法
func IsNomalOrder(bizTyp int8) bool {

	// 此处注释不会被导出
	if bizTyp != 1 && bizTyp != 2 {
		return false
	}

	return true

}
