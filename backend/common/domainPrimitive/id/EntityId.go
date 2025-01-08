package idPrimitive

import "common/domainPrimitive/generator"

type EntityId string

func NewEntityId() EntityId {
	return EntityId(generator.GenerateUUID())
}

func EntityIdFrom(strId string) (EntityId, error) {
	if strId == "" {
		return "", ErrEntityIdIsEmpty
	}

	id, err := generator.UUIDFrom(strId)
	if err != nil {
		return "", ErrEntityIdIsInvalid
	}

	return EntityId(id), nil
}

func (e EntityId) String() string {
	return string(e)
}
