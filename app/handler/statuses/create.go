package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type AddRequest struct {
	Status   string
	MediaIDs []int
}

type AddResponse struct {
	ID       int64           `json:"id"`
	Account  object.Account  `json:"account"`
	Content  string          `json:"content"`
	CreateAt object.DateTime `json:"created_at"`
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req AddRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httperror.BadRequest(w, err)
		return
	}

	status := new(object.Status)
	status.Content = req.Status
	// accountとstatusの紐付けが不明なためハードコーディング
	status.AccountID = 2
	repo := h.app.Dao.Status()
	newStatus, err := repo.CreateStatus(ctx, status)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	// accountとstatusの紐付けが不明なためハードコーディング
	account, _ := h.app.Dao.Account().FindByUsername(ctx, "hoge")

	res := AddResponse{
		ID:       newStatus.ID,
		Account:  *account,
		Content:  newStatus.Content,
		CreateAt: newStatus.CreatedAt,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
	}
}
