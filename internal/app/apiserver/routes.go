package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func (srv *APIServer) handleHello() http.HandlerFunc {
	type response struct {
		Apiname    string `json:"api_name"`
		Version    string `json:"version"`
		Countusers int64  `json:"count_users"`
		Countlinks int64  `json:"count_links"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		resp := response{
			Apiname:    "mycoll",
			Version:    "v1",
			Countusers: srv.db.CountUsers(),
			Countlinks: srv.db.CountLinks(),
		}
		json.NewEncoder(w).Encode(&resp)
	}
}

func (srv *APIServer) handleSignup() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil || req.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}

		if srv.db.IsUserExist(req.Name) {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "User is already exist" }`)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Something wrong" }`)
			return
		}
		srv.db.CreateUser(req.Name, string(hash))
		srv.logger.Info("Created new user ", req.Name)

		w.WriteHeader(http.StatusAccepted)
		io.WriteString(w, `{ "message": "Created new user" }`)
	}
}

func (srv *APIServer) handleSignin() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}

		if !srv.db.IsUserExist(req.Name) {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Something wrong" }`)
			return
		}

		dbname, dbpassword := srv.db.GetUserNamePassword(req.Name)
		if dbname != req.Name {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong username or password" }`)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(req.Password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong username or password" }`)
			return
		}

		io.WriteString(w, `{ "accessToken": "`+srv.newToken(req.Name)+`" }`)
		srv.logger.Info("Token issued")
	}
}

func (srv *APIServer) handleDeleteUser() http.HandlerFunc {
	type request struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}
		dbname, dbpassword := srv.db.GetUserNamePassword(req.Name)
		if dbname != req.Name {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong username or password" }`)
			return
		}
		err = bcrypt.CompareHashAndPassword([]byte(dbpassword), []byte(req.Password))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong username or password" }`)
			return
		}

		countDeleted := srv.db.DeleteUser(req.Name)
		if countDeleted < 1 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Something wrong" }`)
			return
		}
		io.WriteString(w, `{ "message": "Deleted is success" }`)
		srv.logger.Info("One user was deleted")
	}
}

func (srv *APIServer) handleCreateLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Create link")
	}
}

func (srv *APIServer) handleFindLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Find links")
	}
}

func (srv *APIServer) handleUpdateLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Update links")
	}
}

func (srv *APIServer) handleDeleteLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Delete links")
	}
}
