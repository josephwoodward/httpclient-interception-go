package httpclientinterception

type InterceptorBuilder struct {
	builder *configurationBuilder
	opts    []*registrationOptions
}

func (i *InterceptorBuilder) RegisterOptions(options ...*registrationOptions) {
	i.opts = options
	for _, o := range options {
		o.builders = append(o.builders, i.builder)
	}
}
