package idPrimitive

import (
	"strconv"
)

type EntityId int

func EntityIdFrom(strId string) (EntityId, error) {
	if strId == "" {
		return 0, ErrEntityIdIsEmpty
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return 0, ErrEntityIdWrongFormat
	}

	if id <= 0 {
		return 0, ErrEntityIdIsInvalid
	}

	return EntityId(id), nil
}
