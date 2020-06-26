package models

type Options struct {
	timeout string
	payload string
}

func (o Options) GetTimeout() string {
	return o.timeout
}

func (o Options) GetPayload() string {
	return o.payload
}

func NewOptions(timeout, payload string) (options *Options) {
	options = new(Options)
	options.timeout = timeout
	options.payload = payload

	return
}
