package factory

import (
	"design_pattern/abstract_factory_pattern/product"
	"design_pattern/abstract_factory_pattern/product/XiaoMiProduct"
)

type XiaoMiFactory struct {
}

func (x XiaoMiFactory) NewPhone() product.IPhone {
	return XiaoMiProduct.XiaoMiPhone{}
}

func (x XiaoMiFactory) NewComputer() product.IComputer {
	return XiaoMiProduct.XiaoMiComputer{}
}

func (x XiaoMiFactory) NewRouter() product.IRouter {
	return XiaoMiProduct.XiaoMiRouter{}
}
