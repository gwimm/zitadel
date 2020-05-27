package model

import es_models "github.com/caos/zitadel/internal/eventstore/models"

type UserGrant struct {
	es_models.ObjectRoot

	State     UserGrantState
	UserID    string
	ProjectID string
	RoleKeys  []string
}

type UserGrantState int32

const (
	USERGRANTSTATE_ACTIVE UserGrantState = iota
	USERGRANTSTATE_INACTIVE
	USERGRANTSTATE_REMOVED
)

func (u *UserGrant) IsValid() bool {
	return u.ProjectID != "" && u.UserID != ""
}

func (u *UserGrant) IsActive() bool {
	return u.State == USERGRANTSTATE_ACTIVE
}

func (u *UserGrant) IsInactive() bool {
	return u.State == USERGRANTSTATE_INACTIVE
}