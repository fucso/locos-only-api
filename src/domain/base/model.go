package base

import (
	"time"
)

// Model
type Model struct {
	ID        *int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// Initializer
type ModelOption struct {
	ID        *int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func NewModelOption() *ModelOption {
	return &ModelOption{}
}

func NewModel(o *ModelOption) *Model {
	m := &Model{}

	if o.ID != nil {
		m.ID = o.ID
	}

	if o.CreatedAt != nil {
		m.CreatedAt = o.CreatedAt
	}

	if o.UpdatedAt != nil {
		m.UpdatedAt = o.UpdatedAt
	}

	return m
}

// Builder
type ModelBuilder[T domain, B builder[T]] struct {
	Option  *ModelOption
	builder *B
}

func NewModelBuilder[T domain, B builder[T]](b *B) *ModelBuilder[T, B] {
	return &ModelBuilder[T, B]{
		Option:  NewModelOption(),
		builder: b,
	}
}

func (b ModelBuilder[T, B]) ID(id int) B {
	b.Option.ID = &id
	return *b.builder
}

func (b ModelBuilder[T, B]) CreatedAt(t time.Time) B {
	b.Option.CreatedAt = &t
	return *b.builder
}

func (b ModelBuilder[T, B]) UpdatedAt(t time.Time) B {
	b.Option.UpdatedAt = &t
	return *b.builder
}
