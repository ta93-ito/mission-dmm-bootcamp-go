package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

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
	accountDTO := dto.Account{
		Username:    account.Username,
		DisplayName: account.DisplayName,
		CreateAt:    account.CreateAt,
		Avatar:      account.Avatar,
		Header:      account.Header,
		Note:        account.Note,
	}

	res := ReadResponse{accountDTO}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
