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
	"github.com/lesomnus/entpb/internal/example/ent/membership"
	"github.com/lesomnus/entpb/internal/example/ent/team"
	"github.com/lesomnus/entpb/internal/example/role"
)

// MembershipCreate is the builder for creating a Membership entity.
type MembershipCreate struct {
	config
	mutation *MembershipMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (mc *MembershipCreate) SetDateCreated(t time.Time) *MembershipCreate {
	mc.mutation.SetDateCreated(t)
	return mc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableDateCreated(t *time.Time) *MembershipCreate {
	if t != nil {
		mc.SetDateCreated(*t)
	}
	return mc
}

// SetAccountID sets the "account_id" field.
func (mc *MembershipCreate) SetAccountID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetAccountID(u)
	return mc
}

// SetTeamID sets the "team_id" field.
func (mc *MembershipCreate) SetTeamID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetTeamID(u)
	return mc
}

// SetRole sets the "role" field.
func (mc *MembershipCreate) SetRole(r role.Role) *MembershipCreate {
	mc.mutation.SetRole(r)
	return mc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableRole(r *role.Role) *MembershipCreate {
	if r != nil {
		mc.SetRole(*r)
	}
	return mc
}

// SetID sets the "id" field.
func (mc *MembershipCreate) SetID(u uuid.UUID) *MembershipCreate {
	mc.mutation.SetID(u)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *MembershipCreate) SetNillableID(u *uuid.UUID) *MembershipCreate {
	if u != nil {
		mc.SetID(*u)
	}
	return mc
}

// SetAccount sets the "account" edge to the Account entity.
func (mc *MembershipCreate) SetAccount(a *Account) *MembershipCreate {
	return mc.SetAccountID(a.ID)
}

// SetTeam sets the "team" edge to the Team entity.
func (mc *MembershipCreate) SetTeam(t *Team) *MembershipCreate {
	return mc.SetTeamID(t.ID)
}

// Mutation returns the MembershipMutation object of the builder.
func (mc *MembershipCreate) Mutation() *MembershipMutation {
	return mc.mutation
}

// Save creates the Membership in the database.
func (mc *MembershipCreate) Save(ctx context.Context) (*Membership, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MembershipCreate) SaveX(ctx context.Context) *Membership {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MembershipCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MembershipCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MembershipCreate) defaults() {
	if _, ok := mc.mutation.DateCreated(); !ok {
		v := membership.DefaultDateCreated()
		mc.mutation.SetDateCreated(v)
	}
	if _, ok := mc.mutation.Role(); !ok {
		v := membership.DefaultRole
		mc.mutation.SetRole(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := membership.DefaultID()
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MembershipCreate) check() error {
	if _, ok := mc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Membership.date_created"`)}
	}
	if _, ok := mc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account_id", err: errors.New(`ent: missing required field "Membership.account_id"`)}
	}
	if _, ok := mc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team_id", err: errors.New(`ent: missing required field "Membership.team_id"`)}
	}
	if _, ok := mc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "Membership.role"`)}
	}
	if v, ok := mc.mutation.Role(); ok {
		if err := membership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "Membership.role": %w`, err)}
		}
	}
	if _, ok := mc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Membership.account"`)}
	}
	if _, ok := mc.mutation.TeamID(); !ok {
		return &ValidationError{Name: "team", err: errors.New(`ent: missing required edge "Membership.team"`)}
	}
	return nil
}

func (mc *MembershipCreate) sqlSave(ctx context.Context) (*Membership, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MembershipCreate) createSpec() (*Membership, *sqlgraph.CreateSpec) {
	var (
		_node = &Membership{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(membership.Table, sqlgraph.NewFieldSpec(membership.FieldID, field.TypeUUID))
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := mc.mutation.DateCreated(); ok {
		_spec.SetField(membership.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := mc.mutation.Role(); ok {
		_spec.SetField(membership.FieldRole, field.TypeEnum, value)
		_node.Role = value
	}
	if nodes := mc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.AccountTable,
			Columns: []string{membership.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.AccountID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.TeamIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   membership.TeamTable,
			Columns: []string{membership.TeamColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TeamID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MembershipCreateBulk is the builder for creating many Membership entities in bulk.
type MembershipCreateBulk struct {
	config
	err      error
	builders []*MembershipCreate
}

// Save creates the Membership entities in the database.
func (mcb *MembershipCreateBulk) Save(ctx context.Context) ([]*Membership, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Membership, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MembershipMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MembershipCreateBulk) SaveX(ctx context.Context) []*Membership {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MembershipCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MembershipCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
