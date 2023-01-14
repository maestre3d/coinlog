package contact

type CreateCommand struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
	UserID      string `json:"user_id" validate:"required"`
	LinkedToID  string `json:"linked_to_user"`
	ImageURL    string `json:"image_url" validate:"omitempty,url"`
}

type UpdateCommand struct {
	ID          string `json:"id" validate:"required"`
	DisplayName string `json:"display_name" validate:"required,lte=96"`
	LinkedToID  string `json:"linked_to_user"`
	ImageURL    string `json:"image_url" validate:"omitempty,url"`
}
