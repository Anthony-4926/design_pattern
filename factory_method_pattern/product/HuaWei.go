package product

import "fmt"

type HuaWei struct {
	Factory string
}

func (h HuaWei) Call() {
	fmt.Printf("HuaWei call, factory is %v\n", h.Factory)
}
