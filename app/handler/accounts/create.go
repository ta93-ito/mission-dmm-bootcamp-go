package accounts

import (
	"encoding/json"
	"net/http"

	"yatter-backend-go/app/domain/dto"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/handler/httperror"
)

// Request body for `POST /v1/accounts`
type (
	AddRequest struct {
		Username string
		Password string
	}

	AddResponse struct {
		dto.Account
	}

	ReadRequest struct {
		Username string
	}
	ReadResponse struct {
		dto.Account
	}
)

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

	repo := h.app.Dao.Account()
	createdAccount, err := repo.CreateAccount(ctx, account)
	if err != nil {
		httperror.InternalServerError(w, err)
	}
	accountDTO := dto.Account{
		Username:    createdAccount.Username,
		DisplayName: createdAccount.DisplayName,
		CreateAt:    createdAccount.CreateAt,
		Avatar:      createdAccount.Avatar,
		Header:      createdAccount.Header,
		Note:        createdAccount.Note,
	}

	res := AddResponse{accountDTO}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(res); err != nil {
		httperror.InternalServerError(w, err)
		return
	}
}
