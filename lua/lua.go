package lua

import (
	l "github.com/yuin/gopher-lua"
)

func Match(testString string, pattern string) *l.LTable {
	L := l.NewState()
	defer L.Close()

	if err := L.DoFile("lua/test.lua"); err != nil {
		panic(err)
	}
	if err := L.CallByParam(l.P{
		Fn:      L.GetGlobal("Match"), // name of Lua function
		NRet:    1,                    // number of returned values
		Protect: true,                 // return err or panic
	}, l.LString(testString), l.LString(pattern)); err != nil {
		panic(err)
	}
	// Get the returned value from the stack and cast it to a lua.LString
	if table, ok := L.Get(-1).(*l.LTable); ok {
		return table
	} // Pop the returned value from the stack
	L.Pop(1)
	return &l.LTable{}
}
