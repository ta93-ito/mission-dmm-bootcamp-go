package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

// Request body for `POST /v1/accounts`
type AddRequest struct {
	Username string
	Password string
}

type ReadRequest struct {
	Username string
}

// Handle request for `POST /v1/accounts`
func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	account := new(object.Account)
	account.Username = req.Username
	if err := account.SetPassword(req.Password); err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	repo := h.app.Dao.Account() // domain/repository の取得
	newAccount, err := repo.CreateAccount(ctx, account)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newAccount); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := chi.URLParam(r, "username")
	if username == "" {
		err := errors.New("username is required")
		httperror.BadRequest(w, err)
		return
	}
	req := ReadRequest{Username: username}

	repo := h.app.Dao.Account()
	account, err := repo.FindByUsername(ctx, req.Username)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(account); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
