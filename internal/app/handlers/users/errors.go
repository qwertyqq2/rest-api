package users

import "errors"

var ErrInvalidToken = errors.New("invalid auth token")
var ErrTokenValid = errors.New("token isn`t valid")
var ErrUserAlreadyExists = errors.New("user with such credentials already exist")
