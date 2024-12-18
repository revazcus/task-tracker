package jsonApiModel

type ApiObjectType string

func (t ApiObjectType) String() string {
	return string(t)
}

type ApiObjectKey struct {
	objectId   string
	objectType ApiObjectType
}

func ApiObjectKeyFrom(objectId string, objectType ApiObjectType) ApiObjectKey {
	return ApiObjectKey{
		objectId:   objectId,
		objectType: objectType,
	}
}

func (k ApiObjectKey) ObjectId() string {
	return k.objectId
}

func (k ApiObjectKey) ObjectType() ApiObjectType {
	return k.objectType
}
