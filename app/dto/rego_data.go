package dto

type Input struct {
	UserId   string `json:"user_id"`
	Project  string `json:"project"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
