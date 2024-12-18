package jsonApiModel

type JsonApiPayload struct {
	Meta     JsonApiMeta      `json:"meta"`
	Data     []*JsonApiObject `json:"data"`
	Included []*JsonApiObject `json:"included"`
}

func NewEmptyJsonApiPayload() *JsonApiPayload {
	return &JsonApiPayload{
		Meta:     make(JsonApiMeta),
		Data:     make([]*JsonApiObject, 0),
		Included: make([]*JsonApiObject, 0),
	}
}

func (p *JsonApiPayload) AddMeta(key string, data interface{}) {
	p.Meta[key] = data
}

func (p *JsonApiPayload) AddData(data ...*JsonApiObject) {
	p.Data = append(p.Data, data...)
}

func (p *JsonApiPayload) AddInclude(includes ...*JsonApiObject) {
	p.Included = append(p.Included, includes...)
}
