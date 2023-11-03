package lua

import (
	l "github.com/yuin/gopher-lua"
)

func Match(testString string, pattern string) []string {
	L := l.NewState()
	defer L.Close()

	if err := L.DoFile("lua/test.lua"); err != nil {
		return []string{}
	}
	if err := L.CallByParam(l.P{
		Fn:      L.GetGlobal("Match"), // name of Lua function
		NRet:    1,                    // number of returned values
		Protect: true,                 // return err or panic
	}, l.LString(testString), l.LString(pattern)); err != nil {
		return []string{}
	}
	// Get the returned value from the stack and cast it to a lua.LString
	if table, ok := L.Get(-1).(*l.LTable); ok {
		// Put the results in an array
		strArr := []string{}
		table.ForEach(func(l1, l2 l.LValue) {
			strArr = append(strArr, l2.String())
		})
		return strArr
	} // Pop the returned value from the stack
	L.Pop(1)

	return []string{}
}
