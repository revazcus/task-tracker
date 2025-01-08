package jsonApiModel

import (
	"fmt"
	"infrastructure/errors"
)

func ErrRelationshipByKeyNotFound(relationshipKey, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q not found in object with id = %q and type %q", relationshipKey, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrRelationshipDataIsNotApiBaseObject(relationshipKey, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key =%q is not JsonApiBaseObject in object with id = %q and type = %q", relationshipKey, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrRelationshipDataIsNotApiBaseObjects(relationshipKey, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q is not []JsonApiBaseObject in object with id = %q and type %q", relationshipKey, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrRelationshipDataIdNotApiBaseObject(relationshipKey, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q is not JsonApiBaseObject in object with id = %q and type = %q", relationshipKey, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrRelationshipNotContainField(relationshipKey, field, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Relationship by key = %q not contain field %q in object with id = %q and type %q", relationshipKey, field, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrFailCastRelationshipField(relationshipKey, field, objectId, objectType string) error {
	errMsg := fmt.Sprintf("Fail cast relationship field value. RelationshipKey = %q. Field = %q. ObjectId = %q. ObjectType = %q", relationshipKey, field, objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrUnsupportedRelationshipDataStruct(unsupportedData interface{}) error {
	errMsg := fmt.Sprintf("Unsupported relationship data struct. Expected JsonAiBaseObject. Unsupported struct = `%s`", unsupportedData)
	return errors.NewError("SYS", errMsg)
}

func ErrIncludeForRelationshipNotFound(objectId, objectType string) error {
	errMsg := fmt.Sprintf("Include for relationship not found. Id = `%s`. Type = `%s`", objectId, objectType)
	return errors.NewError("SYS", errMsg)
}

func ErrNotUsedInclude(objectId, objectType string) error {
	errMsg := fmt.Sprintf("Include not user id relationships. Id = `%s`. Type = `%s`", objectId, objectType)
	return errors.NewError("SYS", errMsg)
}
