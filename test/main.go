package main

type Test struct {
}

type TestOption interface {
	DaggerObject

	Do(m *Test) *Test
}

func New(
	opts ...TestOption,
) *Test {
	m := &Test{}

	for _, opt := range opts {
		m = opt.Do(m)
	}

	return m
}

func (m *Test) Container() *Container {
	return dag.Container().From("busybox")
}
