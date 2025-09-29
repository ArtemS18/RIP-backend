package models

import (
	"failiverCheck/internal/app/ds"
)

type ComponentsRes struct {
	Components []ds.Component `json:"components"`
}
