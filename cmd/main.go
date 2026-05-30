package main

import (
	"fmt"
	"log"
	"m-faheem-khan/totp/pkg/store"
	"m-faheem-khan/totp/pkg/totp"
	"net/http"
	"strconv"
)

type Server struct {
	dbStore *store.Store
}

func (srv *Server) totpSetupRequestHandler(w http.ResponseWriter, req *http.Request) {
	key, err := totp.GenerateSecret()
	if err != nil {
		w.Write([]byte("Invalid Id - id must be int"))
		return
	}

	id := srv.dbStore.PutSecret(key)
	w.Write([]byte(fmt.Sprintf("{\"secret\": \"%s\", \"id\": \"%d\"}", key, id)))
}

func (srv *Server) totpRequestHandler(w http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(req.PathValue("id"))
	if err != nil {
		w.Write([]byte("Invalid Id - id must be int"))
		return
	}

	row := srv.dbStore.GetSecret(id)

	totp, err := totp.GenerateTOTP([]byte(row.Secret), 30, 6)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error generating TOTP: %v", err)))
		return
	}

	w.Write([]byte(fmt.Sprintf("{\"secret\": \"%s\", \"totp\": \"%s\"}", totp.Secret, totp.TOTP)))
}

func main() {

	dbStore, err := store.NewStore("app.db")
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer dbStore.Close()

	srv := &Server{dbStore: dbStore}

	srv.dbStore.Setup() // create db

	mux := http.NewServeMux()
	mux.HandleFunc("GET /setup", srv.totpSetupRequestHandler)
	mux.HandleFunc("GET /totp/{id}", srv.totpRequestHandler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	} else {
		fmt.Println("Server running on 8080")
	}
}
