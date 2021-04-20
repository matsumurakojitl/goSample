package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data1 = Data1{}
		var slice1 []*Item
		slice1 = append(slice1, itemStruct("taro", 10))
		slice1 = append(slice1, itemStruct("hanako", 18))
		// var test1 = Item{Name: "taro", Age: 10}
		// var test2 = Item{Name: "hanako", Age: 18}
		data1.Title = "Hello json!"
		data1.List = slice1
		// data1.List["test1"] = Item{Name: "taro", Age: 10}
		// data1.List["test2"] = Item{Name: "hanako", Age: 18}

		outputJson, err := json.Marshal(&data1)
		if err != nil {
			log.Print(err)
			fmt.Fprint(w, string(outputJson), err)
		}

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(outputJson))
		log.Print(r.URL.Path)
	})
	http.Handle("/", router)
	log.Print("localhost:8080")
	http.ListenAndServe(":8080", nil)
}

type Data1 struct {
	Title string `json:"title"`
	List  []*Item `json:"list"`
}

type Item struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func itemStruct(name string, age int) (r *Item) {
	r = new(Item)
	r.Name = name
	r.Age = age
	return r
}
