package http

type Metrics interface {
	Pull() map[string]interface{}
}
