package builder

import (
	"design_pattern/biulder_pattern/assembly_line/def"
	"design_pattern/biulder_pattern/assembly_line/product"
)

type IBuilder interface {
	SetScreenType(t def.ScreenType) IBuilder
	SetStorageCapacity(sc def.StorageCapacity) IBuilder
	SetProcessor(p def.Processor) IBuilder
	Build() product.IPhone
}
