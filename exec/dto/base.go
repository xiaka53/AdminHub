package dto

type Page struct {
	Page int `form:"page" json:"page" validate:"min=1" zh:"页码"`
	Size int `form:"size" json:"size" validate:"min=1" zh:"每页数量"`
}

type OrderBy struct {
	OrderBy string `form:"orderBy" json:"orderBy" validate:"omitempty" zh:"排序"`
}
