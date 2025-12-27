package uuidpb

import (
	"github.com/google/uuid"
)

func (u *UUID) AsUUID() uuid.UUID {
	if u == nil {
		return uuid.Nil
	}
	var val [16]byte
	copy(val[:], u.Value)
	return val
}

func New(val uuid.UUID) UUID {
	return UUID{
		Value: val[:],
	}
}
