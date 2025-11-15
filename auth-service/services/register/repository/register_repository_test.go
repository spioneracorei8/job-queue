package repository

import (
	"auth-service/helper"
	"auth-service/models"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DATA-DOG/go-sqlmock.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func mockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		t.Fatalf("Error initializing GORM with mock database: %v", err)
	}

	return gormDB, mock
}

func TestCreateAccount(t *testing.T) {
	var (
		id      = uuid.FromStringOrNil("62833985-5106-4f41-a0e0-2b803e7eac5e")
		user_id = uuid.FromStringOrNil("85fd55a7-ccc0-424f-a475-9f7ba7945b63")
		now     = helper.NewTimestampFromTime(time.Now())
	)
	var account = &models.Account{
		Id:             &id,
		UserId:         &user_id,
		Username:       "test",
		PasswordBcrypt: "$2a$10$vB3GHrMdQwitstd4dBfM9eeN3ct2YAKrrQy0F2QsV6bRxf9tcgBRW",
		WebAccess:      "APPLICATION",
		Status:         "ACTIVE",
		CreatedBy:      "test",
		UpdatedBy:      "test",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	t.Run("success", func(t *testing.T) {
		gormDB, mock := mockDB(t)

		// registerRepo := new(_mock_register.RegisterRepository)

		mock.ExpectBegin()
		mock.ExpectQuery(`INSERT INTO "account"`).
			WithArgs(
				account.Username,
				account.PasswordBcrypt,
				account.WebAccess,
				account.Status,
				account.CreatedBy,
				account.UpdatedBy,
				account.CreatedAt,
				account.UpdatedAt,
				account.Id,
				account.UserId,
			).
			WillReturnRows(sqlmock.NewRows([]string{"id", "user_id"}).
				AddRow(account.Id, account.UserId))
		mock.ExpectCommit()

		registerRepo := NewRegisterRepoImpl(gormDB)
		errResult := registerRepo.CreateAccount(account)

		assert.NoError(t, errResult)
	})

	t.Run("error", func(t *testing.T) {
		
	})
}
