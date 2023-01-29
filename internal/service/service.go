package service

import "context"

type Jobs interface {
	Create(ctx context.Context, form JobCreateForm) error
}

type Employers interface {
	Create(ctx context.Context, form EmployerCreateForm) error
}

type Seekers interface {
	Create(ctx context.Context, form SeekerCreateForm) error
}

type JobCreateForm struct {
	Title string `json:"title"`
}

type EmployerCreateForm struct {
	Name string `json:"name"`
}

type SeekerCreateForm struct {
	Name string `json:"name"`
}
