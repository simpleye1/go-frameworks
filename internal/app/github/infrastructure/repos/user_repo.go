package repos

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"test/internal/app/github/domain/repos"
)

type PostgresUserRepository struct {
	logger *zap.Logger
	db     *gorm.DB
}

func (p PostgresUserRepository) FindUserById(ID uint64) *repos.UserRecord {
	var user repos.UserRecord
	result := p.db.First(&user, ID)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}
	return &user
}

func NewPostgresUserRepository(logger *zap.Logger, db *gorm.DB) *PostgresUserRepository {
	p := &PostgresUserRepository{
		logger: logger.With(zap.String("type", "PostgresUserRepository")),
		db:     db,
	}
	return p
}
