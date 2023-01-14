package user

type CreateCommand struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
}

type UpdateCommand struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
}
