package statuses

import (
	"encoding/json"
	"net/http"
	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
	"yatter-backend-go/app/handler/request"
)

type (
	ReadRequest struct {
		ID int64
	}
	ReadResponse struct {
		ID       int64           `json:"id"`
		Account  dto.Account     `json:"account"`
		Content  string          `json:"content"`
		CreateAt object.DateTime `json:"create_at"`
	}
)

func (h *handler) Read(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := request.IDOf(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	req := ReadRequest{ID: id}

	status, err := h.app.Dao.Status().FindStatus(ctx, req.ID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	account, err := h.app.Dao.Account().FindByID(ctx, status.AccountID)
	if err != nil {
		httperror.InternalServerError(w, err)
	}

	accountDTO := dto.Account{
		Username:    account.Username,
		DisplayName: account.DisplayName,
		Avatar:      account.Avatar,
		Header:      account.Header,
		Note:        account.Note,
		CreateAt:    account.CreateAt,
	}

	res := ReadResponse{
		ID:       status.ID,
		Account:  accountDTO,
		Content:  status.Content,
		CreateAt: status.CreateAt,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
