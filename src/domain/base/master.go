package base

// Model
type Master struct {
	Name string
}

// Initializer
type MasterOption struct {
	Name *string
}

func NewMasterOption() *MasterOption {
	return &MasterOption{}
}

func NewMaster(o *MasterOption) *Master {
	m := &Master{}
	if o.Name != nil {
		m.Name = *o.Name
	}

	return m
}

// Builder
type MasterBuilder[T domain, B builder[T]] struct {
	Option  *MasterOption
	builder *B
}

func NewMasterBuilder[T domain, B builder[T]](b *B) *MasterBuilder[T, B] {
	return &MasterBuilder[T, B]{
		Option:  NewMasterOption(),
		builder: b,
	}
}

func (b MasterBuilder[T, B]) Name(n string) B {
	b.Option.Name = &n
	return *b.builder
}
