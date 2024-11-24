package facotry

import (
	"design_pattern/simple_factory_pattern/def"
	"design_pattern/simple_factory_pattern/product"
)

func NewXiaoMiPhone(phoneType int) product.IPhone {
	switch phoneType {
	case def.XiaoMi6:
		return product.XiaoMi6{}
	case def.XiaoMi7:
		return product.XiaoMi7{}
	case def.XiaoMi8:
		return product.XiaoMi8{}
	default:
		return nil
	}
}

var phoneMap = map[int]product.IPhone{
	def.XiaoMi6: product.XiaoMi6{},
	def.XiaoMi7: product.XiaoMi7{},
	def.XiaoMi8: product.XiaoMi8{},
}

func NewXiaoMiPhone2(phoneType int) product.IPhone {
	return phoneMap[phoneType]
}
