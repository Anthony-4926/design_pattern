package factory

import (
	"design_pattern/abstract_factory_pattern/product"
	"design_pattern/abstract_factory_pattern/product/HuaWeiProduct"
)

type HuaWeiFactory struct {
}

func (h HuaWeiFactory) NewPhone() product.IPhone {
	return HuaWeiProduct.HuaWeiPhone{}
}

func (h HuaWeiFactory) NewComputer() product.IComputer {
	return HuaWeiProduct.HuaWeiComputer{}
}

func (h HuaWeiFactory) NewRouter() product.IRouter {
	return HuaWeiProduct.HuaWeiRouter{}
}
