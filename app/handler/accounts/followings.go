package accounts

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

type FollowingsRequest struct {
	Username string
	Limit    int64
}

type FollowingsResponse struct {
	Accounts []dto.Account `json:"-"`
}

func (h *handler) GetFollowings(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	username := chi.URLParam(r, "username")
	if username == "" {
		err := errors.New("username is required")
		httperror.BadRequest(w, err)
		return
	}
	limit, err := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 64)
	if err != nil {
		limit = 40
	}
	req := FollowingsRequest{Username: username, Limit: limit}

	repo := h.app.Dao.Account()
	accounts, err := repo.GetFollowings(ctx, req.Username, req.Limit)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	res := FollowingsResponse{Accounts: accounts}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
