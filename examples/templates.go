package examples

import (
	"html/template"
	"net/http"
	"os"
)

// go native templating system for HTML
func Templates() {
	div1()
	//div2()
}

// div with name printed in terminal
func div1() {
	html := "<div>{{.name}}</div>"
	tpl, _ := template.New("index").Parse(html)

	data := map[string]interface{}{
		"name": "Miguel",
	}

	tpl.Execute(os.Stdout, data)
}

// div with name fetched from url printed in browser
func div2() {
	html := "<div>{{.name}}</div>"
	tpl, _ := template.New("index").Parse(html)

	data := map[string]interface{}{}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		param := request.URL.Query().Get("name")
		data["name"] = param
		tpl.Execute(writer, data)
	})
	http.ListenAndServe(":8080", nil)
}
