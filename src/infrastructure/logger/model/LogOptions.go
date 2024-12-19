package logModel

type Option func(opts *Options)

type Options struct {
	withStackTrace bool
	component      string
	fields         []*LogField
}

func (o *Options) WithStackTrace() bool {
	return o.withStackTrace
}

func (o *Options) GetComponent() string {
	return o.component
}

func (o *Options) GetFields() []*LogField {
	return o.fields
}
