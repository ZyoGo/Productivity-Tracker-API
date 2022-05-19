package spec

type UpsertTodosSpec struct {
	UserId  string `json:"user_id" validate:"uuid4"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
