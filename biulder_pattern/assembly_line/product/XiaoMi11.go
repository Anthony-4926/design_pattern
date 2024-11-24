package product

import "fmt"

type XiaoMi11 struct {
	MobilePhone
}

func (x XiaoMi11) Call() {
	fmt.Printf("屏幕:%s ", x.MobilePhone.ScreenType)
	fmt.Printf("处理器:%s ", x.MobilePhone.Processor)
	fmt.Printf("存储:%d ", x.MobilePhone.StorageCapacity)
	fmt.Println("打电话")
}
