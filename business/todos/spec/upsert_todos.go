package spec

type UpsertTodosSpec struct {
	UserId  string `validate:"uuid4"`
	Content string `validate:"required"`
	Status  string `validate:"required"`
}
