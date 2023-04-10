package categorydto

type CategoryRequest struct {
	Name string `json:"name" form:"name" validate:"required"`
}
