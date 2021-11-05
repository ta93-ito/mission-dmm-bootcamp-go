package object

type (
	Status struct {
		ID        int64 `json:"-"`
		AccountID `json:"account_id,"`
		Content   string   `json:"content"`
		CreatedAt DateTime `json:"created_at"`
	}
)
