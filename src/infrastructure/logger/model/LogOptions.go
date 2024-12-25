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

func WithStackTrace() Option {
	return func(opts *Options) {
		opts.withStackTrace = true
	}
}

func WithComponent(component string) Option {
	return func(opts *Options) {
		opts.component = component
	}
}

func WithIntField(key string, value int) Option {
	return func(opts *Options) {
		opts.fields = append(opts.fields, &LogField{Key: key, Integer: value})
	}
}

func WithFloatField(key string, value float64) Option {
	return func(opts *Options) {
		opts.fields = append(opts.fields, &LogField{Key: key, Float: value})
	}
}

func WithStringField(key string, value string) Option {
	return func(opts *Options) {
		opts.fields = append(opts.fields, &LogField{Key: key, String: value})
	}
}

func WithObjectField(key string, value interface{}) Option {
	return func(opts *Options) {
		opts.fields = append(opts.fields, &LogField{Key: key, Object: value})
	}
}
