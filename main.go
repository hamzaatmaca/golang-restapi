package main

import(
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Book struct{
	Id string `json:"id"`
	Price string `json:"price"`
	Title string `json:"title"`
	Author *Author `json:"author"`
}

type Author struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var books []Book

func getBooks(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
}

func deleteBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"]{
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, item := range books{
		if item.Id == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}

	}
}


func createBook(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.Id = strconv.Itoa(rand.Intn(100000000))
	books = append(books,book)
	json.NewEncoder(w).Encode(book)
}

func updateBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index, item := range books{
		if item.Id == params["id"]{
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.Id = strconv.Itoa(rand.Intn(100000000))
			books = append(books,book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r :=mux.NewRouter()

	books = append(books,Book{Id:"1",Price:"12 TL",Title:"Book1",  Author: &Author{Firstname:"Author One",Lastname:"Autohr One"} })
	books = append(books,Book{Id:"2",Price:"14 TL",Title:"Book2",  Author: &Author{Firstname:"Author Two",Lastname:"Author Two"} })
	books = append(books,Book{Id:"2",Price:"45 TL",Title:"Book3",  Author: &Author{Firstname:"Author Three",Lastname:"Author Three"} })
	books = append(books,Book{Id:"2",Price:"76 TL",Title:"Book4",  Author: &Author{Firstname:"Author Four",Lastname:"Author Four"} })

	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}",getBook).Methods("GET")
	r.HandleFunc("/books",createBook).Methods("POST")
	r.HandleFunc("/books/{id}",updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}",deleteBook).Methods("DELETE")

	fmt.Printf("Server starting on PORT ----> 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}




