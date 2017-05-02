package main

import (
	"testing"
)

// func Test_GetJokes(t *testing.T) {
// 	if i := GetJokes(); i == nil {
// 		t.Error("GetJokes没通过")
// 	} else {
// 		t.Log("第一个测试通过了")
// 	}
// }

func Test_Writetofile(t *testing.T) {
	if i := Writetofile(GetJokes()); i != "yes" {
		t.Error("Writetofile没通过")
	} else {
		t.Log("第二个测试通过了")
	}
}
