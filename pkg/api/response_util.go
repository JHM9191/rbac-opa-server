package pkg

import (
	"rbac-opa-server-mariadb/app/constants"
	"rbac-opa-server-mariadb/app/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseCode constants.ResponseCode, data T) dto.ApiResponse[T] {
	return BuildResponse_(data, responseCode, responseCode.GetResponseMessage())
}

func BuildResponse_[T any](data T, code constants.ResponseCode, message string) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		Code:    int(code),
		Message: message,
		Data:    data,
	}
}
