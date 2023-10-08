package database

import (
	"testing"

	"github.com/ruhancs/api-pattern/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user,_ := entity.NewUser("N1", "n.emai.com", "12345")
	userDB := NewUserRepository(db)

	err = userDB.Create(user)
	assert.Nil(t,err)
	
	var userFounded entity.User
	err = db.First(&userFounded, "id=?", user.ID).Error
	assert.Nil(t,err)
	assert.Equal(t, userFounded.ID,user.ID)
	assert.Equal(t, userFounded.Name,user.Name)
	assert.Equal(t, userFounded.Email,user.Email)
	assert.NotEmpty(t, userFounded.Password)
}

func TestFindUserByEmail(t *testing.T) {
	db,err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})
	user,_ := entity.NewUser("N1", "n.emai.com", "12345")
	userDB := NewUserRepository(db)

	err = userDB.Create(user)
	assert.Nil(t,err)
	
	userFounded,err := userDB.FindByEmail(user.Email)

	assert.Nil(t,err)
	assert.Equal(t,userFounded.ID,user.ID)
	assert.Equal(t,userFounded.Name,user.Name)
	assert.Equal(t,userFounded.Email,user.Email)
	assert.NotEmpty(t, userFounded.Password)
}