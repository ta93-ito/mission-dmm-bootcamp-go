package object

type (
	Status struct {
		ID        int64    `json:"id" db:"id"`
		AccountID int64    `json:"account_id" db:"account_id"`
		Content   string   `json:"content" db:"content"`
		CreateAt  DateTime `json:"create_at" db:"create_at"`
	}
)
