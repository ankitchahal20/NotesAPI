package main

import (
	"github.com/ankit/project/NotesAPI/notes-api/server"
)

func main() {

	server.Start()

	// //fmt.Println("Server running !!!")
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// mux := chi.NewRouter()
	// mux.Post("/signup", routes.Signup)
	// mux.Post("/login", routes.Login)
	// mux.Get("/getnote", routes.GetNote)
	// mux.Post("/createnote", routes.CreateNote)
	// mux.Delete("/deletenote", routes.DeleteNote)
	// http.ListenAndServe(":8090", mux)
}
