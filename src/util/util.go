package util

import (
	"math/rand"
	"time"
)

func RandInt32(lim int32) int32 {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(lim)
}

func RandString(list []string) string {
	return list[RandInt32(int32(len(list)))]
}
