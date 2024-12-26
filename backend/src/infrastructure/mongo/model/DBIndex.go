package mongoModel

type DBIndexType int

const (
	DBIndexAsc DBIndexType = 1
)

type DBIndex struct {
	Collection string
	Name       string
	Keys       []string
	Type       DBIndexType
	Uniq       bool
}

type DBTextIndex struct {
	Collection string
	Name       string
	Keys       []string
}
