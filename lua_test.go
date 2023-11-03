package main

import (
	lua "msyavuz/peth/lua"
	"testing"
)

func TestMatch(t *testing.T) {
	arr := lua.Match("asd 1 asd", "(.+) (%d+) (.+)")
	arr2 := lua.Match("asd", "(%d)")
	if len(arr) != 3 {
		t.Fatalf("Should return 3 matches")
	}
	t.Log(arr2)
	if len(arr2) != 0 {
		t.Fatalf("Should return no matches")
	}
}
