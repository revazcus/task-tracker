package jsonApiModel

import "github.com/revazcus/task-tracker/backend/infrastructure/errors"

const (
	LimitMetaKey      = "limit"
	OffsetMetaKey     = "offset"
	TotalItemsMetaKey = "totalItems"
)

type JsonApiPayloadBuilder struct {
	metaMap         map[string]interface{}
	dataObjectsList []*JsonApiObject
	dataObjectsMap  apiObjectsMap
	includeMap      map[ApiObjectKey]*JsonApiObject
	defaultIncludes map[ApiObjectType]*JsonApiObject
	errors          *errors.Errors
}

func NewJsonApiPayloadBuilder() *JsonApiPayloadBuilder {
	return &JsonApiPayloadBuilder{
		metaMap:         make(map[string]interface{}),
		dataObjectsList: make([]*JsonApiObject, 0, 2),
		includeMap:      make(map[ApiObjectKey]*JsonApiObject),
		defaultIncludes: make(map[ApiObjectType]*JsonApiObject),
		errors:          errors.NewErrors(),
	}
}

func (b *JsonApiPayloadBuilder) AddPaginationMeta(meta *PaginationMeta) *JsonApiPayloadBuilder {
	if meta == nil {
		return b
	}
	b.AddMeta(LimitMetaKey, meta.Limit)
	b.AddMeta(OffsetMetaKey, meta.Offset)
	b.AddMeta(TotalItemsMetaKey, meta.TotalItems)
	return b
}

func (b *JsonApiPayloadBuilder) AddMeta(key string, value interface{}) *JsonApiPayloadBuilder {
	b.metaMap[key] = value
	return b
}

func (b *JsonApiPayloadBuilder) AddData(data ...*JsonApiObject) *JsonApiPayloadBuilder {
	for _, dataItem := range data {
		if dataItem == nil {
			continue
		}
		b.dataObjectsList = append(b.dataObjectsList, dataItem)
	}
	return b
}

func (b *JsonApiPayloadBuilder) AddInclude(includeData ...*JsonApiObject) *JsonApiPayloadBuilder {
	for _, includeDataItem := range includeData {
		if includeDataItem == nil {
			continue
		}
		objectKey := includeDataItem.Key()
		b.includeMap[objectKey] = includeDataItem
	}
	return b
}

func (b *JsonApiPayloadBuilder) AddDefaultsIncludes(objectTeype string, defaults *JsonApiObject) *JsonApiPayloadBuilder {
	b.defaultIncludes[ApiObjectType(objectTeype)] = defaults
	return b
}

func (b *JsonApiPayloadBuilder) Build() (JsonApiPayload, error) {
	b.fillApiObjectsMap()
	b.fillNilRelationships()
	b.checkExistsIncludes()
	if b.errors.IsPresent() {
		return JsonApiPayload{}, b.errors
	}
	b.checkUsedIncludes()
	if b.errors.IsPresent() {
		return JsonApiPayload{}, b.errors
	}

	responsePayload := NewEmptyJsonApiPayload()
	b.setMetaToPayload(responsePayload)
	b.setDataToPayload(responsePayload)
	b.setIncludeToPayload(responsePayload)

	return *responsePayload, nil
}

func (b *JsonApiPayloadBuilder) fillApiObjectsMap() {
	b.dataObjectsMap = make(apiObjectsMap, 0, len(b.dataObjectsList))

	for _, dataItem := range b.dataObjectsList {
		objectKey := dataItem.Key()
		b.dataObjectsMap = b.dataObjectsMap.putObject(objectKey, dataItem)
	}
}

func (b *JsonApiPayloadBuilder) fillNilRelationships() {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		if data.Relationships == nil {
			data.Relationships = NewEmptyJsonApiRelationships()
		}
	}

	for _, includeData := range b.includeMap {
		if includeData.Relationships == nil {
			includeData.Relationships = NewEmptyJsonApiRelationships()
		}
	}
}

func (b *JsonApiPayloadBuilder) checkExistsIncludes() {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		b.checkExistsIncludesForRelationships(data.Relationships)
	}

	for _, include := range b.includeMap {
		b.checkExistsIncludesForRelationships(include.Relationships)
	}
}

func (b *JsonApiPayloadBuilder) checkExistsIncludesForRelationships(relationships JsonApiObjectRelationships) {
	relationshipApiObjectWithRequireInclude, err := getRelationshipApiObjectsWithRequireInclude(relationships)
	if err != nil {
		b.errors.AddError(err)
		return
	}

	for _, relationshipItem := range relationshipApiObjectWithRequireInclude {
		relationshipKey := relationshipItem.Key()
		if _, exists := b.includeMap[relationshipKey]; !exists {
			if _, defaultsExist := b.defaultIncludes[relationshipKey.ObjectType()]; defaultsExist {
				defaults := *b.defaultIncludes[relationshipKey.ObjectType()]
				defaults.Id = relationshipItem.Id
				b.includeMap[relationshipKey] = &defaults
			} else {
				b.errors.AddError(ErrIncludeForRelationshipNotFound(relationshipItem.Id, relationshipItem.Type))
				continue
			}
		}
	}
}

func (b *JsonApiPayloadBuilder) checkUsedIncludes() {
	relationshipApiObjectSet, err := b.getRelationshipApiObjectsSet()
	if err != nil {
		b.errors.AddError(err)
		return
	}

	for includeKey := range b.includeMap {
		if _, exists := relationshipApiObjectSet[includeKey]; !exists {
			b.errors.AddError(ErrNotUsedInclude(includeKey.ObjectId(), string(includeKey.objectType)))
		}
	}
}

func (b *JsonApiPayloadBuilder) getRelationshipApiObjectsSet() (map[ApiObjectKey]struct{}, error) {
	relationshipApiObjectsSet := make(map[ApiObjectKey]struct{})
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		if err := fillRelationshipApiObjectsSet(relationshipApiObjectsSet, data.Relationships); err != nil {
			return nil, err
		}
	}

	for _, include := range b.includeMap {
		if err := fillRelationshipApiObjectsSet(relationshipApiObjectsSet, include.Relationships); err != nil {
			return nil, err
		}
	}

	return relationshipApiObjectsSet, nil
}

func (b *JsonApiPayloadBuilder) setMetaToPayload(responsePayload *JsonApiPayload) {
	for key, value := range b.metaMap {
		responsePayload.AddMeta(key, value)
	}
}

func (b *JsonApiPayloadBuilder) setDataToPayload(responsePayload *JsonApiPayload) {
	for _, dataPair := range b.dataObjectsMap {
		data := dataPair.object
		responsePayload.AddData(data)
	}
}

func (b *JsonApiPayloadBuilder) setIncludeToPayload(responsePayload *JsonApiPayload) {
	for _, include := range b.includeMap {
		responsePayload.AddInclude(include)
	}
}

func getRelationshipApiObjectsWithRequireInclude(relationships JsonApiObjectRelationships) ([]JsonApiBaseObject, error) {
	relationshipsDataList := make([]JsonApiBaseObject, 0)
	for _, relationshipItem := range relationships {
		if err := validateRelationship(relationshipItem); err != nil {
			return nil, err
		}
		if relationshipItem.Data == nil {
			continue
		}
		if relationshipItem.Links != nil && relationshipItem.Links.Related != "" {
			continue
		}

		relationshipData, ok := relationshipItem.Data.(JsonApiBaseObject)
		if !ok {
			relationshipArrayData, ok := relationshipItem.Data.([]JsonApiBaseObject)
			if !ok {
				return nil, ErrUnsupportedRelationshipDataStruct(relationshipItem.Data)
			}
			relationshipsDataList = append(relationshipsDataList, relationshipArrayData...)
			continue
		}

		relationshipsDataList = append(relationshipsDataList, relationshipData)
	}

	return relationshipsDataList, nil
}

func fillRelationshipApiObjectsSet(relationshipObjectsSet map[ApiObjectKey]struct{}, relationships JsonApiObjectRelationships) error {
	relationshipApiObjects, err := getAllRelationshipApiObjects(relationships)
	if err != nil {
		return err
	}

	for _, relationshipApiObject := range relationshipApiObjects {
		relationshipKey := relationshipApiObject.Key()
		relationshipObjectsSet[relationshipKey] = struct{}{}
	}
	return nil
}

func getAllRelationshipApiObjects(relationships JsonApiObjectRelationships) ([]JsonApiBaseObject, error) {
	relationshipsDataList := make([]JsonApiBaseObject, 0)
	for _, relationshipItem := range relationships {
		if err := validateRelationship(relationshipItem); err != nil {
			return nil, err
		}
		if relationshipItem.Data == nil {
			continue
		}

		relationshipData, ok := relationshipItem.Data.(JsonApiBaseObject)
		if !ok {
			relationshipArrayData, ok := relationshipItem.Data.([]JsonApiBaseObject)
			if !ok {
				return nil, ErrUnsupportedRelationshipDataStruct(relationshipItem.Data)
			}
			relationshipsDataList = append(relationshipsDataList, relationshipArrayData...)
			continue
		}

		relationshipsDataList = append(relationshipsDataList, relationshipData)
	}
	return relationshipsDataList, nil
}

func validateRelationship(relationshipItem JsonApiObjectRelationship) error {
	if relationshipItem.Data != nil {
		return nil
	}
	if relationshipItem.Links == nil {
		return errors.NewError("SYS", "Relationship data is nil and links is nil")
	}
	if relationshipItem.Links.Self == "" {
		return errors.NewError("SYS", "Relationship data is nil and links.self is empty")
	}
	return nil
}

type apiObjectPair struct {
	key    ApiObjectKey
	object *JsonApiObject
}

type apiObjectsMap []apiObjectPair

func (m apiObjectsMap) putObject(key ApiObjectKey, object *JsonApiObject) apiObjectsMap {
	pair := apiObjectPair{
		key:    key,
		object: object,
	}

	if index, exists := m.findObjectIndex(key); exists {
		m[index] = pair
		return m
	} else {
		return append(m, pair)
	}

}

func (m apiObjectsMap) findObjectIndex(key ApiObjectKey) (index int, exists bool) {
	for inx, apiObjectPairItem := range m {
		if apiObjectPairItem.key == key {
			return inx, true
		}
	}
	return -1, false
}
