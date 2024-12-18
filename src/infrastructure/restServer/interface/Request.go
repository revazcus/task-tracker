package restServerInterface

type RequestModel interface {
	FillFromBytes(req []byte) error
}
