package database

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

var ProductList []Product

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
