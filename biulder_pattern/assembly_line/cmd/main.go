package main

import (
	"design_pattern/biulder_pattern/assembly_line/builder"
	"design_pattern/biulder_pattern/assembly_line/def"
)

func main() {
	xiaoMi11 := builder.NewXiaoMi11Builder().
		SetProcessor(def.Snapdragon8Gen2).
		SetScreenType(def.OLED).
		SetStorageCapacity(def.GB128).
		Build()
	xiaoMi11.Call()
}
