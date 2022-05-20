package spec

type UpsertNotesSpec struct {
	Status  string `validate:"required"`
	Content string `validate:"required"`
	Tags    string `validate:"required"`
	UserId  string `validate:"uuid4"`
}
