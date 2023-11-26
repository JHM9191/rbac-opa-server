package constants

type ResponseCode int
type Headers int
type General int

// Constant API
const (
	Success ResponseCode = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
	InternalError
)

func (r ResponseCode) GetResponseStatus() string {
	return [...]string{
		"SUCCESS",
		"DATA_NOT_FOUND",
		"UNKNOWN_ERROR",
		"INVALID_REQUEST",
		"UNAUTHORIZED",
		"INTERNAL_ERROR",
	}[r-1]
}

func (r ResponseCode) GetResponseMessage() string {
	return [...]string{
		"Success",
		"데이터가 존재하지 않습니다.",
		"알 수 없는 오류가 발생하였습니다.",
		"요청 실패",
		"인증 실패",
		"서버 오류가 발생하였습니다.",
	}[r-1]
}
