package httpclientinterception

type interceptionBuilder struct {
	builder *configurationBuilder
	opts    []*interceptionOptions
}

func (i *interceptionBuilder) RegisterOptions(options ...*interceptionOptions) {
	i.opts = options
	for _, o := range options {
		o.builders = append(o.builders, i.builder)
	}
}
