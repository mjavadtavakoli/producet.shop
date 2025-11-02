package main

import (
	"html/template"
	"log"
	"net/http"
)

// تعریف ساختار محصول
type Product struct {
	Name  string
	Price float64
	Image string
}

// هندلر صفحه اصلی
func homeHandler(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{"تی‌شرت مردانه", 250_000, "/static/image.png"},
		{"شلوار جین", 400_000, "/static/image.png"},
		{"کفش اسپرت", 550_000, "/static/image.png"},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "خطا در بارگذاری قالب", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, products)
	if err != nil {
		http.Error(w, "خطا در رندر قالب", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("🚀 سرور در حال اجرا است روی http://localhost:4038")
	err := http.ListenAndServe(":4038", nil)
	if err != nil {
		log.Fatal(err)
	}
}
