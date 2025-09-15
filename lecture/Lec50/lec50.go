package main
//Refactor The Codebase
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Tahshin. I'm Youtuber. I'm Software Engineer")
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductList []Product

func getProducts(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w,r)
	if r.Method != http.MethodGet {
		http.Error(w, "Please Give Me Get Request", 400)
		return
	}
	sendData(w,ProductList,201)
}
func createProduct(w http.ResponseWriter, r *http.Request) {
	handleCors(w)
	handlePreflightReq(w,r)
	if r.Method != "POST" {
		http.Error(w, "Please Give Me Post Request", 400)
		return
	}
	var newProduct Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Please give me valied json", 400)
		return
	}
	newProduct.ID = len(ProductList) + 1
	ProductList = append(ProductList, newProduct)
	sendData(w,newProduct,200)
}
func handleCors(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods","POST,GET,PUT,PATCH,DELETE,OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type,Habib")
	w.Header().Set("Content-Type", "application/json")

}
func handlePreflightReq(w http.ResponseWriter,r *http.Request){
	if r.Method == "OPTIONS" {
		w.WriteHeader(200)
	}
}
func sendData(w http.ResponseWriter,data interface{},statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}
func main() {
	router := http.NewServeMux()

	router.HandleFunc("/hello", helloHandler)

	router.HandleFunc("/about", aboutHandler)

	router.HandleFunc("/products", getProducts)
	router.HandleFunc("/create-product", createProduct)
	fmt.Println("Server Running On :3000")
	err := http.ListenAndServe(":3000", router)

	if err != nil {
		log.Panicln("Error Starting The Server", err)
	}
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is red, I love orange",
		Price:       100,
		ImgUrl:      "https://upload.wikimedia.org/wikipedia/commons/thumb/4/43/Ambersweet_oranges.jpg/1200px-Ambersweet_oranges.jpg",
	}
	prod2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is green, I hate Apple",
		Price:       150,
		ImgUrl:      "https://www.bbassets.com/media/uploads/p/xl/40033817_11-fresho-apple-green-regular.jpg",
	}
	prod3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is boring ,I feel bored eating bananna",
		Price:       5,
		ImgUrl:      "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcR22edYOSOtbm0fDi4u4IoeXAZVNKHqEZb7Fw&s",
	}
	ProductList = append(ProductList, prod1)
	ProductList = append(ProductList, prod2)
	ProductList = append(ProductList, prod3)
}

/*
      []=>list
	  {}=>Object
	  JSON => javascript object notation
*/
