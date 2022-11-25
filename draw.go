package draw

import (
	"math/rand"
	"time"
)

type MyInt interface {
	int | int32 | int64
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandInt() int {
	return rand.Int()
}

// RandInRange [a, b]
func RandInRange[T MyInt](a, b T) (c T) {
	if a > b {
		tmp := a
		a = b
		b = tmp
	}
	return a + T(RandInt())%Max(b-a+1, 1)
}

func Max[T MyInt](a, b T) (c T) {
	if a > b {
		return a
	}
	return b
}

// WeightRand id => 权重
func WeightRand[TK, TV MyInt](weightedValues map[TK]TV) TK {
	totalVal := TV(0)
	for _, val := range weightedValues {
		totalVal += val
	}
	totalVal = Max(1, totalVal)
	var randNum TV = 1
	if totalVal > 1 {
		randNum = RandInRange(1, totalVal)
	}

	for key, value := range weightedValues {
		randNum -= value
		if randNum <= 0 {
			return key
		}
	}

	return 1
}
