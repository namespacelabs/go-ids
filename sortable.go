package ids

import (
	"github.com/segmentio/ksuid"
)

func NewSortableID() string {
	return ksuid.New().String()
}
