package main

import (
	"fmt"
	"msyavuz/peth/views"
	"net/http"

	"github.com/a-h/templ"

	// TODO: I need name these better
	l "github.com/yuin/gopher-lua"
	lua "msyavuz/peth/lua"
)

func main() {
	index := view.Index()

	http.Handle("/", templ.Handler(index))

	http.HandleFunc("/lua", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case "POST":
			r.ParseForm()

			// Get the pattern and test string from form data
			pattern := r.FormValue("pattern")
			testString := r.FormValue("test")

			// Put the results in an array
			strArr := []string{}
			lTable := lua.Match(testString, pattern)
			lTable.ForEach(func(l1, l2 l.LValue) {
				strArr = append(strArr, l2.String())
			})
			if len(strArr) == 0 {
				listComponent := view.Error()
				listComponent.Render(r.Context(), w)
			} else {
				listComponent := view.CapturedList(strArr)
				listComponent.Render(r.Context(), w)
			}

		case "GET":
			fmt.Fprint(w, "There is no get endpoint.")
		}
	})

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
