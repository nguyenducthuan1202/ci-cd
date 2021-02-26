package main

import "testing"

func Test_Hello(t *testing.T) {

	expect := "Hello World!"
	actual := Hello()

	if expect != actual {
		t.Error("wrong value")
	} else {
		t.Log("test Test_Hello() success")
	}
}

func Test_Hello_2(t *testing.T) {

	expect := "hello world"
	actual := Hello()

	if expect != actual {
		t.Error("wrong value")
	} else {
		t.Log("test Test_Hello_2() success")
	}

}
