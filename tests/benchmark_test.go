package tests_test

import (
	"testing"

	. "github.com/raphaelyancey/gorm/gorm/utils/tests"
)

func BenchmarkCreate(b *testing.B) {
	user := *GetUser("bench", Config{})

	for x := 0; x < b.N; x++ {
		user.ID = 0
		DB.Create(&user)
	}
}

func BenchmarkFind(b *testing.B) {
	user := *GetUser("find", Config{})
	DB.Create(&user)

	for x := 0; x < b.N; x++ {
		DB.Find(&User{}, "id = ?", user.ID)
	}
}

func BenchmarkUpdate(b *testing.B) {
	user := *GetUser("find", Config{})
	DB.Create(&user)

	for x := 0; x < b.N; x++ {
		DB.Model(&user).Updates(map[string]interface{}{"Age": x})
	}
}

func BenchmarkDelete(b *testing.B) {
	user := *GetUser("find", Config{})

	for x := 0; x < b.N; x++ {
		user.ID = 0
		DB.Create(&user)
		DB.Delete(&user)
	}
}
