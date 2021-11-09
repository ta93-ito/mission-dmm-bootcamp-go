package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

type AddRequest struct {
	Status   string
	MediaIDs []int
}

type AddResponse struct {
	ID       int64           `json:"id"`
	Account  dto.Account     `json:"account"`
	Content  string          `json:"content"`
	CreateAt object.DateTime `json:"create_at"`
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
	status.AccountID = 2
	newStatus, err := h.app.Dao.Status().CreateStatus(ctx, status)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	account, err := h.app.Dao.Account().FindByID(ctx, status.AccountID)
	if err != nil {
		httperror.InternalServerError(w, err)
		return
	}

	accountDTO := dto.Account{
		Username:    account.Username,
		DisplayName: account.DisplayName,
		Avatar:      account.Avatar,
		Header:      account.Header,
		Note:        account.Note,
		CreateAt:    account.CreateAt,
	}

	res := AddResponse{
		ID:       newStatus.ID,
		Account:  accountDTO,
		Content:  newStatus.Content,
		CreateAt: newStatus.CreateAt,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
	}
}
