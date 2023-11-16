package metrics

type Metrics struct {
	observes map[string]int
	workers  int
}

func New(workers int) *Metrics {
	return &Metrics{
		workers: workers,
	}
}

func (m *Metrics) Observe(resource string) {
	if _, ok := m.observes[resource]; ok {
		m.observes[resource]++
	} else {
		m.observes[resource] = 1
	}
}

func (m *Metrics) Pull() map[string]interface{} {
	return map[string]interface{}{
		"observes": m.observes,
		"workers":  m.workers,
	}
}
