package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os/exec"
)

type viewData struct {
	Input  string
	Fs     string
	Output string
}

func main() {
	http.HandleFunc("/", indexHandler)
	fmt.Printf("Starting server for testing HTTP POST...\n")
	http.ListenAndServe(":8081", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintln(w, "404 page not found")
		return
	}
	
	input := r.FormValue("input")
	fs := r.FormValue("choice")
	runeInput := []rune(input)
	for _, j := range runeInput {
		if j < 32 || j > 126 {
			if j != 12 && j != 15 && j != 13 && j != 10 {
				fmt.Fprintf(w, "400 bad request")
				return
			}
		}
	}

	output, _ := exec.Command("ascii_fs/ascii_fs", input, fs).Output()

	if string(output) == "500 internal server error" {
		fmt.Fprintln(w, "500 internal server error")
		return
	}
	fmt.Println(string(output))

	data := viewData{
		Input:  input,
		Fs:     fs,
		Output: string(output)}
	tmpl, _ := template.ParseFiles("ascii.html")
	tmpl.Execute(w, data)

}
