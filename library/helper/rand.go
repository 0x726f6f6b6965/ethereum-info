package helper

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

func GeneralDuration(times, minTime, maxTime int, duration time.Duration) time.Duration {
	return time.Duration(times+RandInt(minTime, maxTime)) * duration
}
