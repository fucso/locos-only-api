package venue

import (
	"github.com/fucso/locos-only-api/src/domain/base"
)

// Model
type Venue struct {
	*base.Model
	*base.Master
}

// Initializer
type VenueOption struct {
	*base.ModelOption
	*base.MasterOption
}

func NewVenue(opt *VenueOption) (*Venue, error) {
	m := &Venue{}
	m.Model = base.NewModel(opt.ModelOption)
	m.Master = base.NewMaster(opt.MasterOption)
	return m, nil
}

// Builder
type VenueBuilder struct {
	*base.ModelBuilder[Venue, VenueBuilder]
	*base.MasterBuilder[Venue, VenueBuilder]
	option *VenueOption
}

func NewVenueBuilder() *VenueBuilder {
	builder := &VenueBuilder{
		option: &VenueOption{},
	}
	builder.ModelBuilder = base.NewModelBuilder(builder)
	builder.MasterBuilder = base.NewMasterBuilder(builder)

	return builder
}

func (b VenueBuilder) Build() (*Venue, error) {
	o := b.option
	o.ModelOption = b.ModelBuilder.Option
	o.MasterOption = b.MasterBuilder.Option

	return NewVenue(o)
}
