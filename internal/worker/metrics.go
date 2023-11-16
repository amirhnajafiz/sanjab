package worker

type Metrics interface {
	Observe(resource string)
	CallBackFailure()
}
