package inmemory

import (
	"github.com/google/uuid"
)

type Package struct {
	ID   uuid.UUID
	Size int
}
