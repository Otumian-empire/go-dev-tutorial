package main

import (
	"log"
	"strconv"
)

func IsNil(obj interface{}) bool {
	return obj == nil
}

func IsNotNil(obj interface{}) bool {
	return !IsNil(obj)
}

func Log(obj ...interface{}) {
	log.Println(obj...)
}

func FLog(obj ...interface{}) {
	log.Fatal(obj...)
}

func Recover() {
	if err := recover(); IsNotNil(err) {
		Log("Recover from an error")
		Log(err)
	}
}

func ToInt(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
