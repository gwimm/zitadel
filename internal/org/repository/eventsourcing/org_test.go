package eventsourcing

import (
	"context"
	"testing"

	"github.com/caos/zitadel/internal/api/auth"
	"github.com/caos/zitadel/internal/errors"
	es_models "github.com/caos/zitadel/internal/eventstore/models"
	org_model "github.com/caos/zitadel/internal/org/model"
)

func Test_isReservedValidation(t *testing.T) {
	type res struct {
		isErr              func(error) bool
		agggregateSequence uint64
	}
	type args struct {
		aggregate *es_models.Aggregate
		eventType es_models.EventType
		Events    []*es_models.Event
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no events success",
			args: args{
				aggregate: &es_models.Aggregate{},
				eventType: "object.reserved",
				Events:    []*es_models.Event{},
			},
			res: res{
				isErr:              nil,
				agggregateSequence: 0,
			},
		},
		{
			name: "not reseved success",
			args: args{
				aggregate: &es_models.Aggregate{},
				eventType: "object.reserved",
				Events: []*es_models.Event{
					{
						AggregateID:   "asdf",
						AggregateType: "org",
						Sequence:      45,
						Type:          "object.released",
					},
				},
			},
			res: res{
				isErr:              nil,
				agggregateSequence: 45,
			},
		},
		{
			name: "reseved error",
			args: args{
				aggregate: &es_models.Aggregate{},
				eventType: "object.reserved",
				Events: []*es_models.Event{
					{
						AggregateID:   "asdf",
						AggregateType: "org",
						Sequence:      45,
						Type:          "object.reserved",
					},
				},
			},
			res: res{
				isErr:              errors.IsPreconditionFailed,
				agggregateSequence: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validate := isReservedValidation(tt.args.aggregate, tt.args.eventType)

			err := validate(tt.args.Events...)

			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got: %v", err)
			}
			if err == nil && tt.args.aggregate.PreviousSequence != tt.res.agggregateSequence {
				t.Errorf("expected sequence %d got %d", tt.res.agggregateSequence, tt.args.aggregate.PreviousSequence)
			}
		})
	}
}

func aggregateWithPrecondition() *es_models.Aggregate {
	return nil
}

func Test_uniqueNameAggregate(t *testing.T) {
	type res struct {
		expected *es_models.Aggregate
		isErr    func(error) bool
	}
	type args struct {
		ctx        context.Context
		aggCreator *es_models.AggregateCreator
		orgName    string
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no org name error",
			args: args{
				ctx:        auth.NewMockContext("orgID", "userID"),
				aggCreator: es_models.NewAggregateCreator("test"),
				orgName:    "",
			},
			res: res{
				expected: nil,
				isErr:    errors.IsPreconditionFailed,
			},
		},
		{
			name: "aggregate created",
			args: args{
				ctx:        auth.NewMockContext("orgID", "userID"),
				aggCreator: es_models.NewAggregateCreator("test"),
				orgName:    "asdf",
			},
			res: res{
				expected: aggregateWithPrecondition(),
				isErr:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uniqueNameAggregate(tt.args.ctx, tt.args.aggCreator, "", tt.args.orgName)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && (got.Precondition == nil || got.Precondition.Query == nil || got.Precondition.Validation == nil) {
				t.Errorf("precondition is not set correctly")
			}
		})
	}
}

func Test_uniqueDomainAggregate(t *testing.T) {
	type res struct {
		expected *es_models.Aggregate
		isErr    func(error) bool
	}
	type args struct {
		ctx        context.Context
		aggCreator *es_models.AggregateCreator
		orgDomain  string
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no org domain error",
			args: args{
				ctx:        auth.NewMockContext("orgID", "userID"),
				aggCreator: es_models.NewAggregateCreator("test"),
				orgDomain:  "",
			},
			res: res{
				expected: nil,
				isErr:    errors.IsPreconditionFailed,
			},
		},
		{
			name: "aggregate created",
			args: args{
				ctx:        auth.NewMockContext("orgID", "userID"),
				aggCreator: es_models.NewAggregateCreator("test"),
				orgDomain:  "asdf",
			},
			res: res{
				expected: aggregateWithPrecondition(),
				isErr:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := uniqueDomainAggregate(tt.args.ctx, tt.args.aggCreator, "", tt.args.orgDomain)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && (got.Precondition == nil || got.Precondition.Query == nil || got.Precondition.Validation == nil) {
				t.Errorf("precondition is not set correctly")
			}
		})
	}
}

func TestOrgReactivateAggregate(t *testing.T) {
	type res struct {
		isErr func(error) bool
	}
	type args struct {
		aggCreator *es_models.AggregateCreator
		org        *Org
		ctx        context.Context
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "correct",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "orgID",
						Sequence:    2,
					},
					State: int32(org_model.ORGSTATE_INACTIVE),
				},
			},
		},
		{
			name: "already active error",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "orgID",
						Sequence:    2,
					},
					State: int32(org_model.ORGSTATE_ACTIVE),
				},
			},
			res: res{
				isErr: errors.IsErrorInvalidArgument,
			},
		},
		{
			name: "org nil error",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org:        nil,
			},
			res: res{
				isErr: errors.IsPreconditionFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aggregateCreator := orgReactivateAggregate(tt.args.aggCreator, tt.args.org)
			aggregate, err := aggregateCreator(tt.args.ctx)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && aggregate == nil {
				t.Error("aggregate must not be nil")
			}
		})
	}
}

func TestOrgDeactivateAggregate(t *testing.T) {
	type res struct {
		isErr func(error) bool
	}
	type args struct {
		aggCreator *es_models.AggregateCreator
		org        *Org
		ctx        context.Context
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "correct",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "orgID",
						Sequence:    2,
					},
					State: int32(org_model.ORGSTATE_ACTIVE),
				},
			},
		},
		{
			name: "already inactive error",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "orgID",
						Sequence:    2,
					},
					State: int32(org_model.ORGSTATE_INACTIVE),
				},
			},
			res: res{
				isErr: errors.IsErrorInvalidArgument,
			},
		},
		{
			name: "org nil error",
			args: args{
				aggCreator: es_models.NewAggregateCreator("test"),
				ctx:        auth.NewMockContext("org", "user"),
				org:        nil,
			},
			res: res{
				isErr: errors.IsPreconditionFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aggregateCreator := orgDeactivateAggregate(tt.args.aggCreator, tt.args.org)
			aggregate, err := aggregateCreator(tt.args.ctx)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && aggregate == nil {
				t.Error("aggregate must not be nil")
			}
		})
	}
}

func TestOrgUpdateAggregates(t *testing.T) {
	type res struct {
		aggregateCount int
		isErr          func(error) bool
	}
	type args struct {
		ctx        context.Context
		aggCreator *es_models.AggregateCreator
		existing   *Org
		updated    *Org
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no existing org error",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				existing:   nil,
				updated:    &Org{},
			},
			res: res{
				aggregateCount: 0,
				isErr:          errors.IsPreconditionFailed,
			},
		},
		{
			name: "no updated org error",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				existing:   &Org{},
				updated:    nil,
			},
			res: res{
				aggregateCount: 0,
				isErr:          errors.IsPreconditionFailed,
			},
		},
		{
			name: "no changes",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				existing:   &Org{},
				updated:    &Org{},
			},
			res: res{
				aggregateCount: 0,
				isErr:          errors.IsPreconditionFailed,
			},
		},
		{
			name: "name changed",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				existing: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.ch",
					Name:   "coas",
				},
				updated: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.ch",
					Name:   "caos",
				},
			},
			res: res{
				aggregateCount: 2,
				isErr:          nil,
			},
		},
		{
			name: "domain changed",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				existing: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.swiss",
					Name:   "caos",
				},
				updated: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.ch",
					Name:   "caos",
				},
			},
			res: res{
				aggregateCount: 2,
				isErr:          nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := OrgUpdateAggregates(tt.args.ctx, tt.args.aggCreator, tt.args.existing, tt.args.updated)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got: %v", err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && len(got) != tt.res.aggregateCount {
				t.Errorf("OrgUpdateAggregates() aggregate count = %d, wanted count %d", len(got), tt.res.aggregateCount)
			}
		})
	}
}

func TestOrgCreatedAggregates(t *testing.T) {
	type res struct {
		aggregateCount int
		isErr          func(error) bool
	}
	type args struct {
		ctx        context.Context
		aggCreator *es_models.AggregateCreator
		org        *Org
	}
	tests := []struct {
		name string
		args args
		res  res
	}{
		{
			name: "no org error",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				org:        nil,
			},
			res: res{
				aggregateCount: 0,
				isErr:          errors.IsPreconditionFailed,
			},
		},
		{
			name: "org successful",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.ch",
					Name:   "caos",
				},
			},
			res: res{
				aggregateCount: 3,
				isErr:          nil,
			},
		},
		{
			name: "no domain error",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Name: "caos",
				},
			},
			res: res{
				aggregateCount: 2,
				isErr:          errors.IsPreconditionFailed,
			},
		},
		{
			name: "no name error",
			args: args{
				ctx:        auth.NewMockContext("org", "user"),
				aggCreator: es_models.NewAggregateCreator("test"),
				org: &Org{
					ObjectRoot: es_models.ObjectRoot{
						AggregateID: "sdaf",
						Sequence:    5,
					},
					Domain: "caos.ch",
				},
			},
			res: res{
				aggregateCount: 2,
				isErr:          errors.IsPreconditionFailed,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := orgCreatedAggregates(tt.args.ctx, tt.args.aggCreator, tt.args.org)
			if tt.res.isErr == nil && err != nil {
				t.Errorf("no error expected got %T: %v", err, err)
			}
			if tt.res.isErr != nil && !tt.res.isErr(err) {
				t.Errorf("wrong error got %T: %v", err, err)
			}
			if tt.res.isErr == nil && len(got) != tt.res.aggregateCount {
				t.Errorf("OrgUpdateAggregates() aggregate count = %d, wanted count %d", len(got), tt.res.aggregateCount)
			}
		})
	}
}