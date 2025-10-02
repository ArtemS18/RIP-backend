package dto

type CreateComponentDTO struct {
	Title       string `json:"title" validate:"required"`
	Type        string `json:"type" validate:"required"`
	MTBF        uint32 `json:"mtbf" validate:"required,numeric"`
	MTTR        uint32 `json:"mttr" validate:"required,numeric"`
	Description string `json:"description" validate:"required"`
}

type UpdateComponentDTO struct {
	Title       *string  `json:"title"`
	Type        *string  `json:"type"`
	MTBF        *uint32  `json:"mtbf"`
	MTTR        *uint32  `json:"mttr"`
	Available   *float32 `json:"available"`
	Img         *string  `json:"img"`
	Description *string  `json:"description"`
}
type ComponentDTO struct {
	ID    uint32 `json:"id"`
	Title string `json:"title"`
}

type ComponentsToSystemCalcDTO struct {
	ReplicationCount uint         `json:"replication_count"`
	Component        ComponentDTO `json:"component"`
}
