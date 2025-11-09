package twocrystalballs

import "math"

func TwoCrystalBalls(building []int) int {

	jump := int(math.Sqrt(float64(len(building))))

	i := jump
	for ; i < len(building); i += jump {
		if building[i] == 1 {
			break
		}
	}

	i -= jump

	for j := 0; j <= jump && i < len(building); j++ {
		if building[i] == 1 {
			return i + 1
		} else {
			i++
			continue
		}
	}

	return -1
}
