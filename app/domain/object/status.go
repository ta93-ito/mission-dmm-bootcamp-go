package object

type (
	Status struct {
		ID        int64    `json:"-"`
		AccountID int64    `json:"account_id" db:"account_id"`
		Content   string   `json:"content"`
		CreateAt  DateTime `json:"create_at" db:"create_at"`
	}
)
