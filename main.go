package main

import (
	"fmt"
	"msyavuz/peth/views"
	"net/http"

	"github.com/a-h/templ"

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

			strArr := lua.Match(testString, pattern)

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
