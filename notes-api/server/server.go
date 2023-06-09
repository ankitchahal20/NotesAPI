package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ankit/project/NotesAPI/notes-api/constants"
	"github.com/ankit/project/NotesAPI/notes-api/service"
	"github.com/gin-gonic/gin"
)

func registerLoginEndPoints(handler gin.IRoutes) {
	handler.POST(constants.ForwardSlash+strings.Join([]string{constants.Login}, constants.ForwardSlash), service.Login())
}

func registerSingUpEndPoints(handler gin.IRoutes) {
	handler.POST(constants.ForwardSlash+strings.Join([]string{constants.SignUp}, constants.ForwardSlash), service.SingUp())
}

func registerCreateNotesEndPoints(handler gin.IRoutes) {
	handler.POST(constants.ForwardSlash+strings.Join([]string{constants.CreateNotes}, constants.ForwardSlash), service.CreateNotes())
}

func registerGetNotesEndPoints(handler gin.IRoutes) {
	handler.GET(constants.ForwardSlash+strings.Join([]string{constants.GetNotes}, constants.ForwardSlash), service.GetNotes())
}

func registerGetNoteByIDEndPoints(handler gin.IRoutes) {
	handler.GET(constants.ForwardSlash+strings.Join([]string{constants.GetNotesByID + constants.ForwardSlash + constants.Colon + constants.NotesID}, constants.ForwardSlash), service.GetNotesByID())
}

func registerDeleteNoteByIDEndPoints(handler gin.IRoutes) {
	handler.DELETE(constants.ForwardSlash+strings.Join([]string{constants.DeleteNotesByID + constants.ForwardSlash + constants.Colon + constants.NotesID}, constants.ForwardSlash), service.DeleteNotesByID())
}

func Start() {
	plainHandler := gin.New()

	notesAPIHandler := plainHandler.Group(constants.ForwardSlash + constants.Version).Use(gin.Recovery())
	registerLoginEndPoints(notesAPIHandler)
	registerSingUpEndPoints(notesAPIHandler)
	registerCreateNotesEndPoints(notesAPIHandler)
	registerGetNotesEndPoints(notesAPIHandler)
	registerGetNoteByIDEndPoints(notesAPIHandler)
	registerDeleteNoteByIDEndPoints(notesAPIHandler)

	srv := &http.Server{
		Handler:      plainHandler,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	// Start Server
	go func() {
		log.Println("Starting Server")
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {

	/*
		if somewhere you are listening for output from a channel but in the meanwhile that channel not being given any input,
		the listening place will be blocked until it receives the output.
	*/

	interruptChan := make(chan os.Signal, 1)

	/*
		SIGINT means is Signal Interrupted, send when the user types the INTR character (e.g. Ctrl-C).
		SIGTERM signal is a generic signal used to cause program termination.
		The below statement is to send the interruot signal to the channel "interruptChan".
	*/

	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-interruptChan

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	/*
		some zombie connections may still be there and use your memory,
		in order to solve this, the safest way is to setup a timeout threshold using the Context.
	*/

	defer cancel()

	/*
		The timer context and the cancel function (the cancel function releases its resources) will be returned,
		 and you can use that context to perform Shutdown(ctx) ,
		inside the Shutdown it check if the timer context Done channel is closed and will not run indefinitely.
	*/

	srv.Shutdown(ctx)

	log.Println("Shutting down")
	os.Exit(0)
}
