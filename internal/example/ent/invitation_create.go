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
)

// InvitationCreate is the builder for creating a Invitation entity.
type InvitationCreate struct {
	config
	mutation *InvitationMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (ic *InvitationCreate) SetDateCreated(t time.Time) *InvitationCreate {
	ic.mutation.SetDateCreated(t)
	return ic
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableDateCreated(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetDateCreated(*t)
	}
	return ic
}

// SetInvitee sets the "invitee" field.
func (ic *InvitationCreate) SetInvitee(s string) *InvitationCreate {
	ic.mutation.SetInvitee(s)
	return ic
}

// SetType sets the "type" field.
func (ic *InvitationCreate) SetType(s string) *InvitationCreate {
	ic.mutation.SetType(s)
	return ic
}

// SetDateExpired sets the "date_expired" field.
func (ic *InvitationCreate) SetDateExpired(t time.Time) *InvitationCreate {
	ic.mutation.SetDateExpired(t)
	return ic
}

// SetDateAccepted sets the "date_accepted" field.
func (ic *InvitationCreate) SetDateAccepted(t time.Time) *InvitationCreate {
	ic.mutation.SetDateAccepted(t)
	return ic
}

// SetNillableDateAccepted sets the "date_accepted" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableDateAccepted(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetDateAccepted(*t)
	}
	return ic
}

// SetDateDeclined sets the "date_declined" field.
func (ic *InvitationCreate) SetDateDeclined(t time.Time) *InvitationCreate {
	ic.mutation.SetDateDeclined(t)
	return ic
}

// SetNillableDateDeclined sets the "date_declined" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableDateDeclined(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetDateDeclined(*t)
	}
	return ic
}

// SetDateCanceled sets the "date_canceled" field.
func (ic *InvitationCreate) SetDateCanceled(t time.Time) *InvitationCreate {
	ic.mutation.SetDateCanceled(t)
	return ic
}

// SetNillableDateCanceled sets the "date_canceled" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableDateCanceled(t *time.Time) *InvitationCreate {
	if t != nil {
		ic.SetDateCanceled(*t)
	}
	return ic
}

// SetID sets the "id" field.
func (ic *InvitationCreate) SetID(u uuid.UUID) *InvitationCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *InvitationCreate) SetNillableID(u *uuid.UUID) *InvitationCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetSiloID sets the "silo" edge to the Silo entity by ID.
func (ic *InvitationCreate) SetSiloID(id uuid.UUID) *InvitationCreate {
	ic.mutation.SetSiloID(id)
	return ic
}

// SetSilo sets the "silo" edge to the Silo entity.
func (ic *InvitationCreate) SetSilo(s *Silo) *InvitationCreate {
	return ic.SetSiloID(s.ID)
}

// SetInviterID sets the "inviter" edge to the Account entity by ID.
func (ic *InvitationCreate) SetInviterID(id uuid.UUID) *InvitationCreate {
	ic.mutation.SetInviterID(id)
	return ic
}

// SetInviter sets the "inviter" edge to the Account entity.
func (ic *InvitationCreate) SetInviter(a *Account) *InvitationCreate {
	return ic.SetInviterID(a.ID)
}

// Mutation returns the InvitationMutation object of the builder.
func (ic *InvitationCreate) Mutation() *InvitationMutation {
	return ic.mutation
}

// Save creates the Invitation in the database.
func (ic *InvitationCreate) Save(ctx context.Context) (*Invitation, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *InvitationCreate) SaveX(ctx context.Context) *Invitation {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *InvitationCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *InvitationCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *InvitationCreate) defaults() {
	if _, ok := ic.mutation.DateCreated(); !ok {
		v := invitation.DefaultDateCreated()
		ic.mutation.SetDateCreated(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := invitation.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *InvitationCreate) check() error {
	if _, ok := ic.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Invitation.date_created"`)}
	}
	if _, ok := ic.mutation.Invitee(); !ok {
		return &ValidationError{Name: "invitee", err: errors.New(`ent: missing required field "Invitation.invitee"`)}
	}
	if v, ok := ic.mutation.Invitee(); ok {
		if err := invitation.InviteeValidator(v); err != nil {
			return &ValidationError{Name: "invitee", err: fmt.Errorf(`ent: validator failed for field "Invitation.invitee": %w`, err)}
		}
	}
	if _, ok := ic.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Invitation.type"`)}
	}
	if v, ok := ic.mutation.GetType(); ok {
		if err := invitation.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Invitation.type": %w`, err)}
		}
	}
	if _, ok := ic.mutation.DateExpired(); !ok {
		return &ValidationError{Name: "date_expired", err: errors.New(`ent: missing required field "Invitation.date_expired"`)}
	}
	if _, ok := ic.mutation.SiloID(); !ok {
		return &ValidationError{Name: "silo", err: errors.New(`ent: missing required edge "Invitation.silo"`)}
	}
	if _, ok := ic.mutation.InviterID(); !ok {
		return &ValidationError{Name: "inviter", err: errors.New(`ent: missing required edge "Invitation.inviter"`)}
	}
	return nil
}

func (ic *InvitationCreate) sqlSave(ctx context.Context) (*Invitation, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
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
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *InvitationCreate) createSpec() (*Invitation, *sqlgraph.CreateSpec) {
	var (
		_node = &Invitation{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(invitation.Table, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.DateCreated(); ok {
		_spec.SetField(invitation.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := ic.mutation.Invitee(); ok {
		_spec.SetField(invitation.FieldInvitee, field.TypeString, value)
		_node.Invitee = value
	}
	if value, ok := ic.mutation.GetType(); ok {
		_spec.SetField(invitation.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := ic.mutation.DateExpired(); ok {
		_spec.SetField(invitation.FieldDateExpired, field.TypeTime, value)
		_node.DateExpired = value
	}
	if value, ok := ic.mutation.DateAccepted(); ok {
		_spec.SetField(invitation.FieldDateAccepted, field.TypeTime, value)
		_node.DateAccepted = &value
	}
	if value, ok := ic.mutation.DateDeclined(); ok {
		_spec.SetField(invitation.FieldDateDeclined, field.TypeTime, value)
		_node.DateDeclined = &value
	}
	if value, ok := ic.mutation.DateCanceled(); ok {
		_spec.SetField(invitation.FieldDateCanceled, field.TypeTime, value)
		_node.DateCanceled = &value
	}
	if nodes := ic.mutation.SiloIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.SiloTable,
			Columns: []string{invitation.SiloColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(silo.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.silo_invitations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ic.mutation.InviterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   invitation.InviterTable,
			Columns: []string{invitation.InviterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(account.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_invitations = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// InvitationCreateBulk is the builder for creating many Invitation entities in bulk.
type InvitationCreateBulk struct {
	config
	err      error
	builders []*InvitationCreate
}

// Save creates the Invitation entities in the database.
func (icb *InvitationCreateBulk) Save(ctx context.Context) ([]*Invitation, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Invitation, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvitationMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *InvitationCreateBulk) SaveX(ctx context.Context) []*Invitation {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *InvitationCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *InvitationCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
