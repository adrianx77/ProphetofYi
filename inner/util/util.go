package util

import (
	"math/rand"
	"time"
)

func Initialize()  {
	rand.Seed(time.Now().UnixNano())
}

func Random(s int ,e int) int  {
	r := e-s
	o := rand.Intn(r)
	return o + s
}

