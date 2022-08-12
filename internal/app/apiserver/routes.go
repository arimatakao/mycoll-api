package apiserver

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/arimatakao/mycoll-api/internal/app/database"
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
			Countlinks: srv.db.CountGroupLinks(),
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

		io.WriteString(w, `{ "accessToken": "`+srv.newToken(srv.db.GetUserId(req.Name))+`" }`)
		srv.logger.Info("Token issued")
	}
}

func (srv *APIServer) handleDeleteUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		id, valid := srv.parseTokenFromHeader(r)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		countDeleted := srv.db.DeleteUser(id)
		if countDeleted < 1 {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Something wrong" }`)
			return
		}

		countDeletedLinks := srv.db.DeleteAllGroupLinksByIdOwner(id)

		srv.logger.Info("One user was deleted")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":             "Success",
			"count_deleted_links": countDeletedLinks,
		})
	}
}

func (srv *APIServer) handleCreateLinks() http.HandlerFunc {
	type request struct {
		Title       string   `json:"title"`
		Tags        []string `json:"tags"`
		Description string   `json:"description"`
		Links       []string `json:"links"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		id, valid := srv.parseTokenFromHeader(r)
		if !valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}
		resultid, err := srv.db.CreateGroupLinks(id, req.Title, req.Tags, req.Description, req.Links)
		if err != nil {
			io.WriteString(w, `{ "error": "Something wrong" }`)
		}
		srv.logger.Info("Created new group links")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": resultid,
		})
	}
}

func (srv *APIServer) handleFindLinks() http.HandlerFunc {
	type response struct {
		Links []database.GroupLinks `json:"links"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		id, ok := srv.parseTokenFromHeader(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		result := srv.db.FindAllGroupLinksByIdOwner(id)
		resp := response{
			Links: result,
		}
		srv.logger.Info("Find group links")
		json.NewEncoder(w).Encode(resp)
	}
}

func (srv *APIServer) handleUpdateLinks() http.HandlerFunc {
	type request struct {
		Id          string   `json:"_id"`
		Title       string   `json:"title"`
		Tags        []string `json:"tags"`
		Description string   `json:"description"`
		Links       []string `json:"links"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		_, ok := srv.parseTokenFromHeader(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}
		srv.logger.Info("Updated links")
		countedUpdated := srv.db.UpdateGroupLinksById(req.Id, req.Title, req.Tags, req.Description, req.Links)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message":       "Update success",
			"count_updated": countedUpdated,
		})
	}
}

func (srv *APIServer) handleDeleteLinks() http.HandlerFunc {
	type request struct {
		Id string `json:"_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		_, ok := srv.parseTokenFromHeader(r)
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		var req request
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, `{ "error": "Wrong json" }`)
			return
		}
		srv.logger.Info("Deleted links")
		countDeleted := srv.db.DeleteGroupLinksById(req.Id)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"count_deleted_links": countDeleted,
		})
	}
}
