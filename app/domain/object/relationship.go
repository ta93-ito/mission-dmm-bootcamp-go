package object

type Relationship struct {
	ID         int64 `json:"-"`
	FollowerID int64 `json:"follower_id"`
	FollweeID  int64 `json:"followee_id"`
}
