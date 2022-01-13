package main

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strings"
)

func repetition(st string) map[string]int {

	// using strings.Field Function
	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func hello(w http.ResponseWriter, r *http.Request) {
	var getlist []int

	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "form.html")
	case "POST":
		// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		//v := []int{1,3,4,45}
		comment := r.FormValue("comment")

		for index, element := range repetition(comment) {
			//int1, _ := strconv.ParseInt(index, 6, 12)
			getlist = append(getlist, element)
			fmt.Println(index, "=", element)
		}

		sort.Sort(sort.IntSlice(getlist))
		for i, j := 0, len(getlist)-1; i < j; i, j = i+1, j-1 {
			getlist[i], getlist[j] = getlist[j], getlist[i]

		}
		fmt.Println(getlist)
		for index, element := range repetition(comment) {
			for k, _ := range getlist {

				if k == element {
					fmt.Fprintf(w, "comment = %d ,%s \n", k, index, "\n")
					continue
				}
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", hello)

	fmt.Printf("Starting server for testing HTTP POST...\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
