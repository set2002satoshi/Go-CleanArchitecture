package temporary

import (
	"fmt"
)

const (
	INITIAL_ID = 0
	INITIAL_REVISION = 1
)


type IDENTIFICATION int

func NewIDENTIFICATION(id int) (IDENTIFICATION, error) {

	if id < 0 {
		return IDENTIFICATION(id), fmt.Errorf("Losing numbers included.")
	}

	return IDENTIFICATION(id), nil
}

type REVISION int

