package product

import "fmt"

type XiaoMi struct {
	Factory string
}

func (x XiaoMi) Call() {
	fmt.Printf("XiaoMi call, factory is %v\n", x.Factory)
}
