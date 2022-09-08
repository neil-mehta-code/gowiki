package main

import (
	"fmt"
	"log"
	"net/http"
	"html/template"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

// saves to a .txt file, rewrite using database to save to a database
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//fmt.Fprintf(w, "<h1>Editing %s </h1>" +
	//	"<form action=\"/save/%s\" method=\"POST\">" +
	//	"<textarea name=\"body\">%s</textarea><br>" +
	//	"<input type =\"submit\" value=\"Save\">" +
	//	"</form>",
	//	p.Title, p.Title, p.Body)
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a test page")}
	//p1.save()
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/"editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
