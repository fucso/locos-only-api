package event

import (
	"github.com/fucso/locos-only-api/src/domain/base"
	"github.com/fucso/locos-only-api/src/domain/venue"
)

// Model
type Event struct {
	*base.Model
	*base.Master
	Dates []EventDate
}

type EventDate struct {
	Venue venue.Venue
}

// Initializer
type EventOption struct {
	*base.ModelOption
	*base.MasterOption
}

func NewEvent(opt *EventOption) (*Event, error) {
	m := &Event{}
	m.Model = base.NewModel(opt.ModelOption)
	m.Master = base.NewMaster(opt.MasterOption)
	return m, nil
}

// Builder
type EventBuilder struct {
	*base.ModelBuilder[Event, EventBuilder]
	*base.MasterBuilder[Event, EventBuilder]
	option *EventOption
}

func NewEventBuilder() *EventBuilder {
	builder := &EventBuilder{
		option: &EventOption{},
	}
	builder.ModelBuilder = base.NewModelBuilder(builder)
	builder.MasterBuilder = base.NewMasterBuilder(builder)

	return builder
}

func (b EventBuilder) Build() (*Event, error) {
	o := b.option
	o.ModelOption = b.ModelBuilder.Option
	o.MasterOption = b.MasterBuilder.Option

	return NewEvent(o)
}
