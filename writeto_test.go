package main

import (
	"testing"
)

func Test_GetJokes(t *testing.T) {
	if i := GetJokes(); i == nil {
		t.Error("GetJokes没通过")
	} else {
		t.Log("测试通过了")
	}
}
