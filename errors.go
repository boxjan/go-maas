package maas

import "errors"

var (
	ErrEmptyClient            = errors.New("empty maas client")
	ErrRequireParametersEmpty = errors.New("require parameters is empty")
)
