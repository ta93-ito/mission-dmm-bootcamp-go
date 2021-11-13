package accounts

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"yatter-backend-go/app/handler/httperror"

	"github.com/go-chi/chi"
)

type FollowRequest struct {
	FollowerName string
	FolloweeName string
}

type FollowResponse struct {
	ID         int64 `json:"id"`
	Following  bool  `json:"following"`
	FollowedBy bool  `json:"followed_by"`
}

func (h *handler) Follow(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	followerName := chi.URLParam(r, "follower_name")
	if followerName == "" {
		err := errors.New("follower name is required")
		httperror.BadRequest(w, err)
		return
	}
	followeeName := chi.URLParam(r, "followee_name")
	fmt.Println(followerName, followeeName)
	if followeeName == "" {
		err := errors.New("followee name is required")
		httperror.BadRequest(w, err)
		return
	}

	req := FollowRequest{FollowerName: followerName, FolloweeName: followeeName}

	repo := h.app.Dao.Account()
	err := repo.Follow(ctx, req.FolloweeName, req.FollowerName)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	var res FollowResponse
	res.Following = true

	isFollowed, err := repo.IsFollowed(ctx, req.FollowerName, req.FolloweeName)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}
	res.FollowedBy = isFollowed
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
