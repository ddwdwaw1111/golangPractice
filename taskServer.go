package main

import (
	"log"
	"net/http"
	"os"

	"Internal/taskstore"
)

type taskServer struct {
	store *taskstore.TaskStore
}

func NewTaskServer() *taskServer {
	store := taskstore.New()
	return &taskServer{store: store}
}

func main() {
	mux := http.NewServeMux()
	server := NewTaskServer()
	mux.HandleFunc("/task/", server.taskHandler)
	mux.HandleFunc("/tag/", server.tagHandler)
	mux.HandleFunc("/due/", server.dueHandler)

	log.Fatal(http.ListenAndServe("localhost:"+os.Getenv("SERVERPORT"), mux))
}
