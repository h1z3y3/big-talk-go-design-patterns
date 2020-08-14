package main

type ISubject interface {
	Do() string
}

// RealSubject
type RealSubject struct {
}

func (r *RealSubject) Do() string {
	return "real"
}

// Proxy
type Proxy struct {
	subject RealSubject
}

func (p *Proxy) Do() string {
	result := "pre:"
	result += p.subject.Do()
	result += ":after"

	return result
}
