package main

import (
	"design_pattern/factory_method_pattern/factory"
	"design_pattern/factory_method_pattern/product"
	"fmt"
)

func main() {
	var p product.IPhone

	biYadi := factory.BiYaDi{}

	p = biYadi.NewXiaoMi()
	p.Call()

	p = biYadi.NewHuaWei()
	p.Call()

	fmt.Println("----------------------------")

	fuShiKang := factory.FuShiKang{}

	p = fuShiKang.NewXiaoMi()
	p.Call()
	p = fuShiKang.NewHuaWei()
	p.Call()
}
