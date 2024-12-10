package idPrimitive

import (
	"strconv"
)

type EntityId string

func EntityIdFrom(strId string) (EntityId, error) {
	if strId == "" {
		return "", ErrEntityIdIsEmpty
	}

	id, err := strconv.Atoi(strId)
	if err != nil {
		return "", ErrEntityIdWrongFormat
	}

	if id <= 0 {
		return "", ErrEntityIdIsInvalid
	}

	return EntityId(strId), nil
}
