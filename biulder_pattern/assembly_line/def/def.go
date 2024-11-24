package def

// 手机屏幕类型
type ScreenType string

const (
	OLED   ScreenType = "OLED"
	LCD    ScreenType = "LCD"
	AMOLED ScreenType = "AMOLED"
)

// 手机存储容量类型
type StorageCapacity int

const (
	_ StorageCapacity = iota
	GB32
	GB64
	GB128
	GB256
)

// 手机处理器类型
type Processor string

const (
	A16             Processor = "A16"
	Snapdragon8Gen2 Processor = "Snapdragon 8 Gen 2"
	Kirin9000       Processor = "Kirin 9000"
)
