package jsonApiModel

type JsonApiObject struct {
	Id            string                     `json:"id"`
	Type          string                     `json:"type"`
	Attributes    JsonApiAttributes          `json:"attributes"`
	Relationships JsonApiObjectRelationships `json:"relationships"`
}

func (o JsonApiObject) Key() ApiObjectKey {
	return ApiObjectKeyFrom(o.Id, ApiObjectType(o.Type))
}

func (o JsonApiObject) AddAttribute(key string, data interface{}) {
	o.Attributes[key] = data
}

func (o JsonApiObject) AddRelationship(objectId, objectType string) {
	o.Relationships.AddApiBaseObject(objectId, objectType)
}

func (o JsonApiObject) AddRelationshipBySpecialKey(objectId, objectType, key string) {
	o.Relationships.AddApiBaseObjectBySpecialKey(objectId, objectType, key)
}

func (o JsonApiObject) GetRelationshipApiBaseObject(key string) (JsonApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return JsonApiBaseObject{}, ErrRelationshipByKeyNotFound(key, o.Id, o.Type)
	}

	apiBaseObject, isApiBaseObject := relationship.Data.(JsonApiBaseObject)
	if !isApiBaseObject {
		dataMap, isDataMap := relationship.Data.(map[string]interface{})
		if !isDataMap {
			return JsonApiBaseObject{}, ErrRelationshipDataIsNotApiBaseObject(key, o.Id, o.Type)
		}
		return o.convertMapToApiBaseObject(dataMap, key)
	}

	return apiBaseObject, nil
}

func (o JsonApiObject) GetRelationshipApiBaseObjects(key string) ([]JsonApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return nil, ErrRelationshipByKeyNotFound(key, o.Id, o.Type)
	}

	apiBaseObjects, isApiBaseObjects := relationship.Data.([]JsonApiBaseObject)
	if !isApiBaseObjects {
		return o.convertRelationshipToApiBaseObjects(relationship, key)
	}

	return apiBaseObjects, nil
}

func (o JsonApiObject) TryGetRelationshipApiBaseObjects(key string) ([]JsonApiBaseObject, error) {
	relationship, exists := o.Relationships[key]
	if !exists {
		return nil, ErrRelationshipByKeyNotFound(key, o.Id, o.Type)
	}

	apiBaseObjects, isApiBaseObjects := relationship.Data.([]JsonApiBaseObject)
	apiBaseObject, isApiBaseObject := relationship.Data.(JsonApiBaseObject)
	if !isApiBaseObjects && !isApiBaseObject {
		dataMap, isDataMap := relationship.Data.(map[string]interface{})
		if isDataMap {
			apiBaseObj, err := o.convertMapToApiBaseObject(dataMap, key)
			if err != nil {
				return nil, err
			}
			return []JsonApiBaseObject{apiBaseObj}, nil
		}
		return o.convertRelationshipToApiBaseObjects(relationship, key)
	}

	if isApiBaseObjects {
		return apiBaseObjects, nil
	} else {
		return []JsonApiBaseObject{apiBaseObject}, nil
	}
}

func (o JsonApiObject) convertRelationshipToApiBaseObjects(relationship JsonApiObjectRelationship, relationshipKey string) ([]JsonApiBaseObject, error) {
	apiBaseObjectMapList, isDataMap := relationship.Data.([]map[string]interface{})
	apiBaseObjectsList, isDataList := relationship.Data.([]interface{})
	if !isDataMap && !isDataList {
		return nil, ErrRelationshipDataIsNotApiBaseObjects(relationshipKey, o.Id, o.Type)
	}

	if isDataList {
		apiBaseObjectMapList = make([]map[string]interface{}, len(apiBaseObjectsList))
		for inx, apiBaseObjectItem := range apiBaseObjectsList {
			apiBaseObjectMap, ok := apiBaseObjectItem.(map[string]interface{})
			if !ok {
				return nil, ErrRelationshipDataIdNotApiBaseObject(relationshipKey, o.Id, o.Type)
			}
			apiBaseObjectMapList[inx] = apiBaseObjectMap
		}
	}

	apiBaseObjects := make([]JsonApiBaseObject, len(apiBaseObjectMapList))
	for inx, apiBaseObjectMap := range apiBaseObjectMapList {
		apiBaseObject, err := o.convertMapToApiBaseObject(apiBaseObjectMap, relationshipKey)
		if err != nil {
			return nil, err
		}
		apiBaseObjects[inx] = apiBaseObject
	}

	return apiBaseObjects, nil
}

func (o JsonApiObject) convertMapToApiBaseObject(dataMap map[string]interface{}, relationshipKey string) (JsonApiBaseObject, error) {
	objectId, existsId := dataMap["id"]
	if !existsId {
		return JsonApiBaseObject{}, ErrRelationshipNotContainField(relationshipKey, "id", o.Id, o.Type)
	}

	objectType, existsType := dataMap["type"]
	if !existsType {
		return JsonApiBaseObject{}, ErrRelationshipNotContainField(relationshipKey, "type", o.Id, o.Type)
	}

	objectIdStr, ok := objectId.(string)
	if !ok {
		return JsonApiBaseObject{}, ErrFailCastRelationshipField(relationshipKey, "id", o.Id, o.Type)
	}

	objectTypeStr, ok := objectType.(string)
	if !ok {
		return JsonApiBaseObject{}, ErrFailCastRelationshipField(relationshipKey, "type", o.Id, o.Type)
	}

	return JsonApiBaseObject{
		Id:   objectIdStr,
		Type: objectTypeStr,
	}, nil
}
