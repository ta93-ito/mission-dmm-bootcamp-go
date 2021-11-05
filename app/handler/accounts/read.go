package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

type ReadRequest struct {
	Username string
}

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// requestパッケージを使用するようにする
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
