package main

import (
	"strconv"
)

func StrToInt64(s string) int64 {
	v, _ := strconv.ParseInt(s, 10, 16)
	return v
}

func Int64ToStr(i int64) string {
	return strconv.FormatInt(i, 10)
}
