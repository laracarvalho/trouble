package cmd

import (
	"fmt"
	"testing"
)

func TestSet(t *testing.T) {
	s := make(map[string]string)
	
	r := SetFunc(s, "test", "testvalue")

	if (r["test"] != "testvalue") {
		t.Fatalf("Item not set. Got: %v", r)
	} else {
		fmt.Print(r)
	}
}

func TestGet(t *testing.T) {
	s := make(map[string]string)

	s["test2"] = "mytest"
	
	r := GetFunc(s, "test2")

	if (*r != "mytest") {
		t.Fatalf("Item not set. Got: %v", r)
	} else {
		fmt.Print(r)
	}
}