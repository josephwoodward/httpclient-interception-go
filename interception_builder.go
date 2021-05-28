package httpclientinterception

type interceptionBuilder struct {
	builder *configurationBuilder
	opts    []*interceptionOptions
}

func (b *interceptionBuilder) RegisterOptions(opts *interceptionOptions) {
	// interceptionBuilder contains the rules to match against, opts needs to be able to check against those rules,
	// either by adding the rules to opts, or by passing a pointer
	opts.builders = append(opts.builders, b.builder)

	// Does a builder need opts?
	//b.opts = append(b.opts, opts)
}
