package SerialNumber

import "fmt"

// 各个模块的序列号
const (
	SeparatorStr             = "-"             // 分隔符
	PublicNumberStrID        = "Public"        // 公共模块
	XCheckNumberStrID        = "XCheck"        // XCheck
	NoTurningBackNumberStrID = "NoTurningBack" // NoTurningBack 无法回头
	MagicRingNumberStrID     = "MagicRing"     // MagicRing 魔戒
	ErinNumberStrID          = "Erin"          // Erin 艾琳
	EpicShelterNumberStrID   = "EpicShelter"   // EpicShelter 史诗庇护所
)

// IDCalculate id计算
func IDCalculate() {
	fmt.Println(ReflectionXSSPocPayloadType)
}
