package taskTag

type Tag string

type TagEnum map[string]Tag

func (t Tag) String() string {
	return string(t)
}

const (
	bug     = "Bug"
	feature = "Feature"
	quest   = "Quest"
)

var Tags = TagEnum{
	bug:     bug,
	feature: feature,
	quest:   quest,
}

func (e TagEnum) Bug() Tag {
	return e[bug]
}

func (e TagEnum) Feature() Tag {
	return e[feature]
}

func (e TagEnum) Quest() Tag {
	return e[quest]
}

func (e TagEnum) Of(code string) (Tag, error) {
	tag, ok := e[code]
	if !ok {
		return "", ErrUnsupportedTag(code)
	}

	return tag, nil
}
