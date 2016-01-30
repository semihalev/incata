package writer

import (
	"github.com/mantzas/golinear/model"
	"github.com/twinj/uuid"
	"reflect"
	"testing"
)

func TestItemsAdded(t *testing.T) {

	item1 := model.NewEvent(uuid.NewV4(), "Test 1", "TEST", 1)
	item2 := model.NewEvent(uuid.NewV4(), "Test 2", "TEST", 1)
	expectedItems := []model.Event{}
	expectedItems = append(expectedItems, *item1)
	expectedItems = append(expectedItems, *item2)

	writer := NewMemoryWriter()
	writer.Write(*item1)
	writer.Write(*item2)

	if !reflect.DeepEqual(expectedItems, writer.Data) {
		t.Fatalf("Expected %s, got %s", expectedItems, writer.Data)
	}
}

func TestItemsEmpty(t *testing.T) {
	var expectedItems = make([]string, 0)

	writer := NewMemoryWriter()

	if len(expectedItems) != len(writer.Data) {
		t.Fatalf("Expected %s, got %s", expectedItems, writer.Data)
	}
}