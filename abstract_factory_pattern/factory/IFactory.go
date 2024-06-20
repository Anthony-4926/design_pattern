package factory

import "design_pattern/abstract_factory_pattern/product"

type IFactory interface {
	NewPhone() product.IPhone
	NewComputer() product.IComputer
	NewRouter() product.IRouter
}
