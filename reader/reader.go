package reader

import (
	"github.com/satori/go.uuid"
	"github.com/semihalev/incata/model"
)

// Reader interface for getting events
type Reader interface {
	Read(uuid.UUID) ([]model.Event, error)
}
