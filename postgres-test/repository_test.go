package main_test

import (
	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	"github.com/stretchr/testify/assert"
	"postgres-test"
	"testing"
)

func TestRepository(t *testing.T) {
	t.Run("should return a list of users", func(t *testing.T) {
		// setup
		dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable"

		config := embeddedpostgres.DefaultConfig().
			Username("gorm").
			Password("gorm").
			Database("gorm").
			Port(9920)
		postgres := embeddedpostgres.NewDatabase(config)

		// start postgres
		startErr := postgres.Start()
		assert.NoError(t, startErr)

		db := main.GormDB(dsn)
		migrationErr := db.AutoMigrate(&main.User{})
		assert.NoError(t, migrationErr)

		users := main.UserBuilder{}.BuildMany(10)
		addUsersErr := db.Create(&users).Error
		assert.NoError(t, addUsersErr)

		// given
		repo := main.NewRepository(db)
		// when
		response, err := repo.ListAll()

		// then
		assert.NoError(t, err)
		assert.Len(t, response, 10)

		// stop postgres
		stopErr := postgres.Stop()
		assert.NoError(t, stopErr)
	})
}
