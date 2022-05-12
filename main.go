package main

import (
	// "belajar-API/book"
	"belajar-API/config"
	"belajar-API/entity"
	"belajar-API/handler"
	"belajar-API/repository"
	"belajar-API/service"

	// "fmt"
	// "log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	// "gorm.io/driver/mysql"
)

var (
	db             *gorm.DB              = config.SetupDatabaseConnection()
	bookRepository repository.Repository = repository.NewRepository(db)
	bookService    service.Service       = service.NewService(bookRepository)
	bookHandler                          = handler.NewBookHandler(bookService)
)

func main() {
	defer config.CloseConnectionDatabase(db)

	db.AutoMigrate(&entity.Book{})

	router := gin.Default()

	router.GET("/", bookHandler.RootHandler)
	bookRoutes := router.Group("/api/v1/book")

	// end point
	{
		bookRoutes.GET("/", bookHandler.GetBooksHandler)
		bookRoutes.GET("/:id", bookHandler.GetBookHandler)
		bookRoutes.PUT("/:id", bookHandler.UpdateBookHandler)
		bookRoutes.DELETE("/:id", bookHandler.DeleteBookHandler)
		bookRoutes.POST("/", bookHandler.PostBookHandler)
	}

	router.Run(":9000")
}

// db.AutoMigrate(&book.Book{})

// create
// book := &book.Book{}
// book.Title = "Mendaki mingguan"
// book.Description = "petualangan pemuda"
// book.Rating = 4
// book.Discount = 10
// book.Price = 70000

// err = db.Create(&book).Error
// if err != nil {
// 	fmt.Println("error create book")
// }

// read
// var book book.Book
// err = db.First(&book).Error
// if err != nil {
// 	fmt.Println("error create book")
// }

// fmt.Println("title :", book.Title)
// fmt.Printf("book object : %v", book)

// var books []book.Book
// err = db.Debug().Where("title = ?", "Penari tunggal").Find(&books).Error
// if err != nil {
// 	fmt.Println("error create book")
// }

// for _, b := range books {
// 	fmt.Println("title :", b.Title)
// 	fmt.Printf("book object : %v", b)
// }

// update
// var book book.Book
// err = db.Debug().Where("title = ?", "Penari tunggal").Find(&book).Error
// if err != nil {
// 	fmt.Println("error find book")
// }

// book.Title = "Penari Tunggal (diubah)"
// err = db.Save(&book).Error
// if err != nil {
// 	fmt.Println("error update book")
// }

// fmt.Println("title :", book.Title)
// fmt.Printf("book object : %v", book)

// delete
// var book book.Book
// err = db.Debug().Where("id = ?", 1).Find(&book).Error
// if err != nil {
// 	fmt.Println("error find book")
// }

// err = db.Delete(&book).Error
// if err != nil {
// 	fmt.Println("error update book")
// }

// bookRepository := book.NewRepository(db)

// Find All
// books, err := bookRepository.FindAll()
// if err != nil {
// 	fmt.Println("error find book")
// }

// for _, b := range books {
// 	fmt.Println("title :", b.Title)
// 	fmt.Printf("book object : %v", b)
// }

// Find by Id
// book, err := bookRepository.FindById(2)
// if err != nil {
// 	fmt.Println("error find book")
// }
// fmt.Println("title :", book.Title)

// create
// book := book.Book{
// 	Title:       "Penerjun lembah",
// 	Description: "petualangan nyali",
// 	Rating:      45,
// 	Discount:    0,
// 	Price:       90000,
// }
// bookRepository.Create(book)
