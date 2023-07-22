package mid

import (
	"fmt"
	"math/bits"
)

// 枚举所有[0,12]小时直接的所有以分钟为单位的时间，然后计算二进制中1的数量
func readBinaryWatch(turnedOn int) (ans []string) {
	for h := uint8(0); h < 12; h++ {
		for m := uint8(0); m < 60; m++ {
			//bits.OnesCount8计算二进制数中1的数量
			if bits.OnesCount8(h)+bits.OnesCount8(m) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return
}
