package sleep

import (
	"math/rand"
	"time"
)

func SleepRandom() {
	n := rand.Intn(1000)
	time.Sleep(time.Duration(n) * time.Millisecond)
}
