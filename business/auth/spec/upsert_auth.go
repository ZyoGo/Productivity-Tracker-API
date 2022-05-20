package spec

type UpsertAuthSpec struct {
	Username string `validate:"required,email"`
	Password string `validate:"min=5"`
}
