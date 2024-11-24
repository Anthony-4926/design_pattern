package builder

import (
	"design_pattern/biulder_pattern/assembly_line/def"
	"design_pattern/biulder_pattern/assembly_line/product"
)

type XiaoMi11Builder struct {
	phone product.XiaoMi11
}

func NewXiaoMi11Builder() *XiaoMi11Builder {
	return &XiaoMi11Builder{}
}
func (x *XiaoMi11Builder) SetScreenType(t def.ScreenType) IBuilder {
	x.phone.MobilePhone.ScreenType = t
	return x
}

func (x *XiaoMi11Builder) SetStorageCapacity(sc def.StorageCapacity) IBuilder {
	x.phone.MobilePhone.StorageCapacity = sc
	return x
}

func (x *XiaoMi11Builder) SetProcessor(p def.Processor) IBuilder {
	x.phone.MobilePhone.Processor = p
	return x
}

func (x *XiaoMi11Builder) Build() product.IPhone {
	return x.phone
}
