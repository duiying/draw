package draw

import (
	"fmt"
	"testing"
)

func TestDraw(t *testing.T) {
	TestRandInRange(t)
	TestWeightRand(t)
}

func TestRandInRange(t *testing.T) {
	a := 1
	b := 10
	m := map[int]int{}
	for i := 0; i < 100*10000; i++ {
		randRes := RandInRange(b, a)
		m[randRes]++
	}
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%d 出现了 %d 次", k, v))
	}
	fmt.Println("=================================================================================================")
}

func TestWeightRand(t *testing.T) {
	a := map[int]int{
		1: 1, // 1 的权重占 10%
		2: 3, // 2 的权重占 30%
		3: 6, // 3 的权重占 60%,
	}

	resA := map[int]int{}

	for i := 0; i < 100*10000; i++ {
		k := WeightRand(a)
		resA[k]++
	}

	for k, v := range resA {
		fmt.Println(fmt.Sprintf("a: %d 随机出了 %d 次", k, v))
	}

	fmt.Println("=================================================================================================")

	b := map[int32]int64{
		10: 10,
		20: 10,
		30: 10,
		40: 70,
	}

	resB := map[int32]int{}

	for i := 0; i < 100*10000; i++ {
		k := WeightRand(b)
		resB[k]++
	}

	for k, v := range resB {
		fmt.Println(fmt.Sprintf("b: %d 随机出了 %d 次", k, v))
	}

	fmt.Println("=================================================================================================")

	c := map[int64]int64{
		1:  1,
		2:  1,
		3:  1,
		4:  1,
		5:  1,
		6:  1,
		7:  1,
		8:  1,
		9:  1,
		10: 1,
	}

	resC := map[int64]int{}

	for i := 0; i < 100*10000; i++ {
		k := WeightRand(c)
		resC[k]++
	}

	for k, v := range resC {
		fmt.Println(fmt.Sprintf("c: %d 随机出了 %d 次", k, v))
	}

	fmt.Println("=================================================================================================")

	d := map[int32]int32{
		1: 5,
	}

	resD := map[int32]int{}

	for i := 0; i < 100*10000; i++ {
		k := WeightRand(d)
		resD[k]++
	}

	for k, v := range resD {
		fmt.Println(fmt.Sprintf("d: %d 随机出了 %d 次", k, v))
	}

	fmt.Println("=================================================================================================")
}
