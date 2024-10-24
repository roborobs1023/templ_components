package components

import (
	"math/rand"
	"strconv"
)

func GenerateUniqueKey(id string) string {
	return id + "-" + strconv.Itoa(rand.Intn(100000))
}

func UniqueKey() string {
	return strconv.Itoa(rand.Intn(100000))
}
