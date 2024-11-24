package product

import "design_pattern/biulder_pattern/assembly_line/def"

type IPhone interface {
	Call()
}

// MobilePhone 手机基本结构体
type MobilePhone struct {
	ScreenType      def.ScreenType
	StorageCapacity def.StorageCapacity
	Processor       def.Processor
}
