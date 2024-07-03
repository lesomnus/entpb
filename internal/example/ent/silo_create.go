// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/lesomnus/entpb/internal/example/ent/account"
	"github.com/lesomnus/entpb/internal/example/ent/invitation"
	"github.com/lesomnus/entpb/internal/example/ent/silo"
	"github.com/lesomnus/entpb/internal/example/ent/team"
)

// SiloCreate is the builder for creating a Silo entity.
type SiloCreate struct {
	config
	mutation *SiloMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (sc *SiloCreate) SetDateCreated(t time.Time) *SiloCreate {
	sc.mutation.SetDateCreated(t)
	return sc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (sc *SiloCreate) SetNillableDateCreated(t *time.Time) *SiloCreate {
	if t != nil {
		sc.SetDateCreated(*t)
	}
	return sc
}

// SetAlias sets the "alias" field.
func (sc *SiloCreate) SetAlias(s string) *SiloCreate {
	sc.mutation.SetAlias(s)
	return sc
}

// SetNillableAlias sets the "alias" field if the given value is not nil.
func (sc *SiloCreate) SetNillableAlias(s *string) *SiloCreate {
	if s != nil {
		sc.SetAlias(*s)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SiloCreate) SetName(s string) *SiloCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *SiloCreate) SetNillableName(s *string) *SiloCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// SetDescription sets the "description" field.
func (sc *SiloCreate) SetDescription(s string) *SiloCreate {
	sc.mutation.SetDescription(s)
	return sc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (sc *SiloCreate) SetNillableDescription(s *string) *SiloCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SiloCreate) SetID(u uuid.UUID) *SiloCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SiloCreate) SetNillableID(u *uuid.UUID) *SiloCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (sc *SiloCreate) AddAccountIDs(ids ...uuid.UUID) *SiloCreate {
	sc.mutation.AddAccountIDs(ids...)
	return sc
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (sc *SiloCreate) AddAccounts(a ...*Account) *SiloCreate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return sc.AddAccountIDs(ids...)
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (sc *SiloCreate) AddTeamIDs(ids ...uuid.UUID) *SiloCreate {
	sc.mutation.AddTeamIDs(ids...)
	return sc
}

// AddTeams adds the "teams" edges to the Team entity.
func (sc *SiloCreate) AddTeams(t ...*Team) *SiloCreate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return sc.AddTeamIDs(ids...)
}

// AddInvitationIDs adds the "invitations" edge to the Invitation entity by IDs.
func (sc *SiloCreate) AddInvitationIDs(ids ...uuid.UUID) *SiloCreate {
	sc.mutation.AddInvitationIDs(ids...)
	return sc
}

// AddInvitations adds the "invitations" edges to the Invitation entity.
func (sc *SiloCreate) AddInvitations(i ...*Invitation) *SiloCreate {
	ids := make([]uuid.UUID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddInvitationIDs(ids...)
}

// Mutation returns the SiloMutation object of the builder.
func (sc *SiloCreate) Mutation() *SiloMutation {
	return sc.mutation
}

// Save creates the Silo in the database.
func (sc *SiloCreate) Save(ctx context.Context) (*Silo, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SiloCreate) SaveX(ctx context.Context) *Silo {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SiloCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SiloCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SiloCreate) defaults() {
	if _, ok := sc.mutation.DateCreated(); !ok {
		v := silo.DefaultDateCreated()
		sc.mutation.SetDateCreated(v)
	}
	if _, ok := sc.mutation.Alias(); !ok {
		v := silo.DefaultAlias()
		sc.mutation.SetAlias(v)
	}
	if _, ok := sc.mutation.Name(); !ok {
		v := silo.DefaultName
		sc.mutation.SetName(v)
	}
	if _, ok := sc.mutation.Description(); !ok {
		v := silo.DefaultDescription
		sc.mutation.SetDescription(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := silo.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SiloCreate) check() error {
	if _, ok := sc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Silo.date_created"`)}
	}
	if _, ok := sc.mutation.Alias(); !ok {
		return &ValidationError{Name: "alias", err: errors.New(`ent: missing required field "Silo.alias"`)}
	}
	if v, ok := sc.mutation.Alias(); ok {
		if err := silo.AliasValidator(v); err != nil {
			return &ValidationError{Name: "alias", err: fmt.Errorf(`ent: validator failed for field "Silo.alias": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Silo.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := silo.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Silo.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Silo.description"`)}
	}
	if v, ok := sc.mutation.Description(); ok {
		if err := silo.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Silo.description": %w`, err)}
		}
	}
	return nil
}

func (sc *SiloCreate) sqlSave(ctx context.Context) (*Silo, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SiloCreate) createSpec() (*Silo, *sqlgraph.CreateSpec) {
	var (
		_node = &Silo{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(silo.Table, sqlgraph.NewFieldSpec(silo.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.DateCreated(); ok {
		_spec.SetField(silo.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := sc.mutation.Alias(); ok {
		_spec.SetField(silo.FieldAlias, field.TypeString, value)
		_node.Alias = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(silo.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Description(); ok {
		_spec.SetField(silo.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := sc.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   silo.AccountsTable,
			Columns: []string{silo.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   silo.TeamsTable,
			Columns: []string{silo.TeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.InvitationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   silo.InvitationsTable,
			Columns: []string{silo.InvitationsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SiloCreateBulk is the builder for creating many Silo entities in bulk.
type SiloCreateBulk struct {
	config
	err      error
	builders []*SiloCreate
}

// Save creates the Silo entities in the database.
func (scb *SiloCreateBulk) Save(ctx context.Context) ([]*Silo, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Silo, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SiloMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SiloCreateBulk) SaveX(ctx context.Context) []*Silo {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SiloCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SiloCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}