package pkg

import (
	"errors"
	"fmt"
	"rbac-opa-server-mariadb/app/constants"
)

func PanicException_(key, message string) {
	err := errors.New(message)
	err = fmt.Errorf("%s: %w", key, err)
	if err != nil {
		panic(err)
	}
}

func PanicException(responseStatus constants.ResponseCode) {
	PanicException_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage())
}
