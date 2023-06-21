package routes

import (
	"github.com/gorilla/mux"

	"simpleBlog/backend/controllers"
)

func RegisterBlogRoutes(router *mux.Router, controller *controllers.BlogController) {
	router.HandleFunc("/blogs", controller.GetBlogs).Methods("GET")
	router.HandleFunc("/blogs/{id}", controller.GetBlogByID).Methods("GET")
	router.HandleFunc("/blogs", controller.CreateBlog).Methods("POST")
	router.HandleFunc("/blogs/{id}", controller.UpdateBlog).Methods("PUT")
	router.HandleFunc("/blogs/{id}", controller.DeleteBlog).Methods("DELETE")
}
