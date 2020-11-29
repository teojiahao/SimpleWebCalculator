package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func index(res http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		num1, _ := strconv.Atoi(req.FormValue("number1"))
		num2, _ := strconv.Atoi(req.FormValue("number2"))
		buttons := req.FormValue("buttons")
		strconv.Atoi(req.FormValue("number1"))
		switch buttons {
		case "+":
			fmt.Fprintln(res, num1, buttons, num2, "=", Add(num1, num2))
		case "-":
			fmt.Fprintln(res, num1, buttons, num2, "=", Subtract(num1, num2))
		case "x":
			fmt.Fprintln(res, num1, buttons, num2, "=", Multiply(num1, num2))
		case "/":
			fmt.Fprintln(res, num1, buttons, num2, "=", Divide(num1, num2))
		}
		return
	}

	tpl.ExecuteTemplate(res, "index.gohtml", nil)
}
