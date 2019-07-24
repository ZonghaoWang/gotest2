package main

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)


const (
	OneE = 100000000
	TenM = 10000000
	TenE = 1000000000
	OneK = 1000
	OneM = 1000000
)

type Profile struct {
	ProfileID string
	Credit    int
	Credit1   int
	Credit2   int
	Credit3   int
	FaceCount int
	Active    int
}

func BenchmarkSliceInt(b *testing.B)  {
	rst := make([]int, TenE)
	for ind := range rst {
		rst[ind] = ind
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for ind := range rst {
			cnt += int64(ind)
		}
	}
}

func BenchmarkSliceInt8(b *testing.B)  {
	rst := make([]int8, TenE)
	for ind := range rst {
		rst[ind] = int8(ind % 128)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for ind := range rst {
			cnt += int64(ind)
		}
	}
}

func BenchmarkSliceRand(b *testing.B)  {
	leng := OneE
	rst := make([]int, leng)
	rd := make([]int, leng)
	for ind := range rd {
		rd[ind] = rand.Intn(leng)
		rst[ind] = ind
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for _, ind := range rd {
			cnt += int64(rst[ind])
		}
	}
}

func BenchmarkSliceStack(b *testing.B)  {
	leng := OneE
	rst := make([]Profile, leng)
	for ind := range rst {
		//rst[ind] = Profile{ProfileID: strconv.Itoa(ind), Active:ind, FaceCount:ind}
		rst[ind] = Profile{Active: ind}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for _, ind := range rst {
			cnt += int64(ind.Active)
		}
	}
}

func BenchmarkSliceStack2(b *testing.B)  {
	leng := OneE
	rst := make([]*Profile, leng)
	for ind := range rst {
		//rst[ind] = Profile{ProfileID: strconv.Itoa(ind), Active:ind, FaceCount:ind}
		rst[ind] = &Profile{Active: ind}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for _, ind := range rst {
			cnt += int64(ind.Active)
		}
	}
}

func BenchmarkMapIntKey(b *testing.B)  {
	leng := OneE
	rst := make(map[int]int, leng)
	for ind := 0; ind < leng; ind++ {
		rst[ind] = ind
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for _, v := range rst {
			cnt += int64(v)
		}
	}
}

func BenchmarkMapStringKey(b *testing.B)  {
	leng := OneE
	rst := make(map[string]int, leng)
	for ind := 0; ind < leng; ind++ {
		rst[strconv.Itoa(ind)] = ind
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for _, v := range rst {
			cnt += int64(v)
		}
	}
}

func BenchmarkMapStructValue(b *testing.B)  {
	leng := OneE
	rst := make(map[int]Profile, leng)
	for ind := 0; ind < leng; ind++ {
		rst[ind] = Profile{Active:ind}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var cnt int64 = 0
		for k := range rst {
			cnt += int64(k)
		}
	}
}


func BenchmarkTimer(b *testing.B)  {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		time.Now().Unix()
	}
}