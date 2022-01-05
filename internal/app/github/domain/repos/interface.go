package repos

import "time"

type DetailRepository interface {
	FindDetailById(ID uint64) *DetailRecord
}

type DetailRecord struct {
	ID        uint64
	Name      string
	Price     float32
	CreatedAt time.Time
}

type UserRepository interface {
	FindUserById(ID uint64) *UserRecord
}

type UserRecord struct {
	ID       uint64
	UserName string
	Password string
	Email    string
}
