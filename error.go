package bucketeeropenfeatureprovidergo

import "errors"

var (
	ErrUserNotFound = errors.New("user is not found in the context")
	ErrFlagNotFound = errors.New("flag is not found")
)
