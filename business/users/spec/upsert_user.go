package spec

type UpsertUserSpec struct {
	Username     string `json:"username" validate:"required,email"`
	Password     string `json:"password" validate:"min=5"`
	PhoneNumber int64 `json:"phone_number" validate:"gt=11,number"`
}