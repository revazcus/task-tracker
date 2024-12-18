package jsonApiModel

type JsonApiObjectRelationships map[string]JsonApiObjectRelationship

type JsonApiObjectRelationship struct {
	Links *JsonApiLinks `json:"links,omitempty"`
	Data  interface{}   `json:"data,omitempty"`
}

func NewEmptyJsonApiRelationships() JsonApiObjectRelationships {
	return make(JsonApiObjectRelationships)
}

func (r JsonApiObjectRelationships) AddWithoutDataWithSelfLink(key, selfLink string) {
	r[key] = JsonApiObjectRelationship{
		Links: &JsonApiLinks{
			Self: selfLink,
		}}
}

func (r JsonApiObjectRelationships) AddApiBaseObject(objectId, objectType string) {
	jsonApiBaseObject := NewJsonApiBaseObject(objectId, objectType)
	r[objectType] = JsonApiObjectRelationship{Data: jsonApiBaseObject}
}

func (r JsonApiObjectRelationships) AddApiBaseObjectBySpecialKey(objectId, objectType, key string) {
	jsonApiBaseObject := NewJsonApiBaseObject(objectId, objectType)
	r[key] = JsonApiObjectRelationship{Data: jsonApiBaseObject}
}

func (r JsonApiObjectRelationships) AddApiBaseObjects(key string, objects []JsonApiObject) {
	r[key] = JsonApiObjectRelationship{Data: objects}
}

func (r JsonApiObjectRelationships) AddRelationshipData(key string, data interface{}) {
	r[key] = JsonApiObjectRelationship{Data: data}
}

func (r JsonApiObjectRelationships) AddRelationshipDataWithRelatedLink(key string, data interface{}, relatedLink string) {
	r[key] = JsonApiObjectRelationship{
		Links: &JsonApiLinks{
			Related: relatedLink,
		},
		Data: data,
	}
}

func (r JsonApiObjectRelationships) AddApiBaseObjectWithRelatedLink(objectId, objectType, relatedLink string) {
	jsonApiBaseObject := NewJsonApiBaseObject(objectId, objectType)
	r[objectType] = JsonApiObjectRelationship{
		Links: &JsonApiLinks{
			Related: relatedLink,
		},
		Data: jsonApiBaseObject,
	}
}
