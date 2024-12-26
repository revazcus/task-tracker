package jsonApiModel

type JsonApiBaseObject struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func NewJsonApiBaseObject(objectId, objectType string) JsonApiBaseObject {
	return JsonApiBaseObject{
		Id:   objectId,
		Type: objectType,
	}
}

func (o JsonApiBaseObject) Key() ApiObjectKey {
	return ApiObjectKeyFrom(o.Id, ApiObjectType(o.Type))
}
