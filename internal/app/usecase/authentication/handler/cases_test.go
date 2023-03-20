package handler

import (
	passHelper "test-dans/internal/app/common/helper/password"
	"test-dans/model"
)

var dummyHashedPass, _ = passHelper.HashPassword("test")

var dummyUser = model.UserLogin{
	Username: "test",
	Password: dummyHashedPass,
}
