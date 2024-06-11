package factory

import "design_pattern/factory_method_pattern/product"

type IFactory interface {
	NewXiaoMi() product.IPhone
	NewHuaWei() product.IPhone
}
