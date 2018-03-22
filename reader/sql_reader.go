package reader

import (
	"github.com/satori/go.uuid"
	"github.com/semihalev/incata/marshal"
	"github.com/semihalev/incata/model"
	"github.com/semihalev/incata/storage"
)

// SQLReader for reading events
type SQLReader struct {
	Storage    *storage.Storage
	Serializer marshal.Serializer
}

// NewSQLReader creates a new sql reader
func NewSQLReader(storage *storage.Storage, ser marshal.Serializer) *SQLReader {

	return &SQLReader{Storage: storage, Serializer: ser}
}

// Read from db with source id
func (r *SQLReader) Read(sourceID uuid.UUID) ([]model.Event, error) {

	rows, err := r.Storage.Query(r.Storage.SelectBySourceIDStatement, sourceID.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events = make([]model.Event, 0)

	for rows.Next() {
		var event = new(model.Event)

		if err := rows.Scan(&event.ID, &event.SourceID, &event.Created, &event.EventType, &event.Version, &event.Payload); err != nil {
			return nil, err
		}

		events = append(events, *event)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
