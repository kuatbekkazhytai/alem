package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type ArtistData struct {
	ID           int
	Image        string
	Name         string
	Members      []string
	CreationDate int
	FirstAlbum   string
	Relations    string
	RelationData NewRelation
}

//NewRelation is a struct of related information.
type NewRelation struct {
	ID	int
	DatesLocations map[string][]string
}

var artistInfo []ArtistData
var artistRelation []NewRelation

func main() {
	assets := http.FileServer(http.Dir("assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/artist/", artistHandler)
	http.HandleFunc("/search", searchHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	if r.URL.Path != "/" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	err = json.Unmarshal(body, &artistInfo)
	if err != nil {
		log.Fatalln(err)
		return
	}
	res, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	b = b[9:len(b)-2]
	// fmt.Println(string(b))
	defer res.Body.Close()
	if err != nil {
		errorHandler(w, r, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &artistRelation)
	if err != nil {
		log.Fatalln(err)
		return
	}
	for i := range artistInfo {
		artistInfo[i].RelationData = artistRelation[i]
	}
	tpls := template.Must(template.ParseGlob(("assets/templates/*.html")))
	err = tpls.ExecuteTemplate(w, "index.html", artistInfo)
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
}

func artistHandler(w http.ResponseWriter, r *http.Request) {
	u, err := url.Parse(r.URL.String())
	if err != nil {
		errorHandler(w, r, http.StatusInternalServerError)
		return
	}
	params := u.Query()
	idKey := params.Get("id")
	fmt.Println(idKey)
	id, err := strconv.Atoi(idKey)
	if err != nil || id > len(artistInfo) {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	for _, value := range artistInfo {
		if value.ID == id {
			resp, err := http.Get(value.Relations)
			if err != nil {
				errorHandler(w, r, http.StatusBadRequest)
				return
			}
			body, err := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
			if err != nil {
				errorHandler(w, r, http.StatusInternalServerError)
				return
			}
			defer resp.Body.Close()
			err = json.Unmarshal(body, &value.RelationData)
			if err != nil {
				log.Fatal(err)
			}
			tpls := template.Must(template.ParseGlob("assets/templates/*.html"))
			tpls.ExecuteTemplate(w, "artist.html", value)
		}
	}
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var searchResult []ArtistData
	input := r.FormValue("search-bar")
	for _, value := range artistInfo {
		if strings.ToLower(value.Name) == strings.ToLower(input) {
			searchResult = append(searchResult, value)
		} else {
			for _, v := range value.Members {
				if strings.ToLower(v) == strings.ToLower(input) {
					searchResult = append(searchResult, value)
				}
			}
		}
		if strings.ToLower(strconv.Itoa(value.CreationDate)) == strings.ToLower(input) {
			searchResult = append(searchResult, value)
		}
		if strings.ToLower(value.FirstAlbum) == strings.ToLower(input) {
			searchResult = append(searchResult, value)
		}
		for key, val := range value.RelationData.DatesLocations {
			if key == input {
				searchResult = append(searchResult, value)
			}
			for _, j := range val {
				if j == input {
					searchResult = append(searchResult, value)
				}
			}
		}
	}

	if len(searchResult) == 0 {
		errorHandler(w, r, http.StatusBadRequest)
	} else {
		tpls := template.Must(template.ParseGlob("assets/templates/*.html"))
		tpls.ExecuteTemplate(w, "search.html", searchResult)
	}
}
func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	templates := template.Must(template.ParseGlob("assets/templates/*.html"))

	switch w.WriteHeader(status); status {
	case http.StatusNotFound:
		templates.ExecuteTemplate(w, "err404.html", nil)
	case http.StatusBadRequest:
		templates.ExecuteTemplate(w, "err400.html", nil)
	case http.StatusInternalServerError:
		templates.ExecuteTemplate(w, "err500.html", nil)
	}

}
