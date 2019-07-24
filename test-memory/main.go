package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//testSliceRW(1000000000)
	//testSliceRand(1000000000)
	testSliceStackRW(10000000)
	//testMapRW(10000000, 1.4)
	testMapStackRW(10000000, 1)
}

func testSliceRW(leng int)  {
	rst := make([]int, leng)
	start := time.Now()
	for ind := range rst {
		rst[ind] = ind
	}
	fmt.Printf("write time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
	start = time.Now()
	var cnt int64 = 0
	for ind := range rst {
		cnt += int64(ind)
	}
	fmt.Printf("read time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
}

func testSliceRand(leng int)  {
	rst := make([]int, leng)
	rd := make([]int, leng)
	for ind := range rd {
		rd[ind] = rand.Intn(leng)
		rst[ind] = ind
	}
	start := time.Now()
	var cnt int64 = 0
	for ind := 0; ind < leng; ind++ {
		cnt += int64(rst[rd[ind]])
	}
	fmt.Printf("read time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))

}

func testSliceStackRW(leng int)  {
	rst := make([]Profile, leng)
	start := time.Now()
	for ind := range rst {
		rst[ind] = Profile{ProfileID: "a", Active: ind, FaceCount: ind}
	}
	fmt.Printf("write time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
	start = time.Now()
	var cnt int64 = 0
	var cs string
	for _, v := range rst {
		cnt += int64(v.Credit)
		cnt += int64(v.Credit1)
		cnt += int64(v.Credit2)
		cnt += int64(v.Credit3)
		cnt += int64(v.FaceCount)
		cnt += int64(v.Active)
		cs = v.ProfileID
	}
	fmt.Println(cs)
	fmt.Printf("read time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
}

func testMapRW(leng int, factor float64)  {
	rst := make(map[int]int, int(factor * float64(leng)))
	start := time.Now()
	for ind := 0; ind < leng; ind++ {
		rst[ind] = ind
	}
	fmt.Printf("write time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
	start = time.Now()
	var cnt int64 = 0
	for k, v := range rst {
		cnt += int64(v)
		cnt += int64(k)
	}
	fmt.Printf("read time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
}



type Profile struct {
	ProfileID string
	Credit    int
	Credit1   int
	Credit2   int
	Credit3   int
	FaceCount int
	Active    int
}

func testMapStackRW(leng int, factor float64)  {
	rst := make(map[int]Profile, int(float64(leng) * factor))
	start := time.Now()
	for ind := 0; ind < leng; ind++ {
		rst[ind] = Profile{ProfileID: "a", Active:ind, FaceCount:ind}
	}
	fmt.Printf("write time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
	var cnt int64 = 0
	start = time.Now()
	var cs string
	for k, v := range rst {
		cnt += int64(k)
		cnt += int64(v.Credit)
		cnt += int64(v.Credit1)
		cnt += int64(v.Credit2)
		cnt += int64(v.Credit3)
		cnt += int64(v.FaceCount)
		cnt += int64(v.Active)
		cs = v.ProfileID
	}
	fmt.Println(cs)
	fmt.Printf("read time cost is %f ns\n", float64(time.Now().Sub(start)) / float64(leng))
}