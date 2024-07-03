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
	"github.com/lesomnus/entpb/internal/example/ent/identity"
	"github.com/lesomnus/entpb/internal/example/ent/user"
)

// IdentityCreate is the builder for creating a Identity entity.
type IdentityCreate struct {
	config
	mutation *IdentityMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (ic *IdentityCreate) SetDateCreated(t time.Time) *IdentityCreate {
	ic.mutation.SetDateCreated(t)
	return ic
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableDateCreated(t *time.Time) *IdentityCreate {
	if t != nil {
		ic.SetDateCreated(*t)
	}
	return ic
}

// SetName sets the "name" field.
func (ic *IdentityCreate) SetName(s string) *IdentityCreate {
	ic.mutation.SetName(s)
	return ic
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableName(s *string) *IdentityCreate {
	if s != nil {
		ic.SetName(*s)
	}
	return ic
}

// SetDescription sets the "description" field.
func (ic *IdentityCreate) SetDescription(s string) *IdentityCreate {
	ic.mutation.SetDescription(s)
	return ic
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableDescription(s *string) *IdentityCreate {
	if s != nil {
		ic.SetDescription(*s)
	}
	return ic
}

// SetKind sets the "kind" field.
func (ic *IdentityCreate) SetKind(s string) *IdentityCreate {
	ic.mutation.SetKind(s)
	return ic
}

// SetVerifier sets the "verifier" field.
func (ic *IdentityCreate) SetVerifier(s string) *IdentityCreate {
	ic.mutation.SetVerifier(s)
	return ic
}

// SetID sets the "id" field.
func (ic *IdentityCreate) SetID(u uuid.UUID) *IdentityCreate {
	ic.mutation.SetID(u)
	return ic
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ic *IdentityCreate) SetNillableID(u *uuid.UUID) *IdentityCreate {
	if u != nil {
		ic.SetID(*u)
	}
	return ic
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ic *IdentityCreate) SetOwnerID(id uuid.UUID) *IdentityCreate {
	ic.mutation.SetOwnerID(id)
	return ic
}

// SetOwner sets the "owner" edge to the User entity.
func (ic *IdentityCreate) SetOwner(u *User) *IdentityCreate {
	return ic.SetOwnerID(u.ID)
}

// Mutation returns the IdentityMutation object of the builder.
func (ic *IdentityCreate) Mutation() *IdentityMutation {
	return ic.mutation
}

// Save creates the Identity in the database.
func (ic *IdentityCreate) Save(ctx context.Context) (*Identity, error) {
	ic.defaults()
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IdentityCreate) SaveX(ctx context.Context) *Identity {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IdentityCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IdentityCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *IdentityCreate) defaults() {
	if _, ok := ic.mutation.DateCreated(); !ok {
		v := identity.DefaultDateCreated()
		ic.mutation.SetDateCreated(v)
	}
	if _, ok := ic.mutation.Name(); !ok {
		v := identity.DefaultName
		ic.mutation.SetName(v)
	}
	if _, ok := ic.mutation.Description(); !ok {
		v := identity.DefaultDescription
		ic.mutation.SetDescription(v)
	}
	if _, ok := ic.mutation.ID(); !ok {
		v := identity.DefaultID()
		ic.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IdentityCreate) check() error {
	if _, ok := ic.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Identity.date_created"`)}
	}
	if _, ok := ic.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Identity.name"`)}
	}
	if v, ok := ic.mutation.Name(); ok {
		if err := identity.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Identity.name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Identity.description"`)}
	}
	if v, ok := ic.mutation.Description(); ok {
		if err := identity.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Identity.description": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`ent: missing required field "Identity.kind"`)}
	}
	if v, ok := ic.mutation.Kind(); ok {
		if err := identity.KindValidator(v); err != nil {
			return &ValidationError{Name: "kind", err: fmt.Errorf(`ent: validator failed for field "Identity.kind": %w`, err)}
		}
	}
	if _, ok := ic.mutation.Verifier(); !ok {
		return &ValidationError{Name: "verifier", err: errors.New(`ent: missing required field "Identity.verifier"`)}
	}
	if v, ok := ic.mutation.Verifier(); ok {
		if err := identity.VerifierValidator(v); err != nil {
			return &ValidationError{Name: "verifier", err: fmt.Errorf(`ent: validator failed for field "Identity.verifier": %w`, err)}
		}
	}
	if _, ok := ic.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Identity.owner"`)}
	}
	return nil
}

func (ic *IdentityCreate) sqlSave(ctx context.Context) (*Identity, error) {
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

func (ic *IdentityCreate) createSpec() (*Identity, *sqlgraph.CreateSpec) {
	var (
		_node = &Identity{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(identity.Table, sqlgraph.NewFieldSpec(identity.FieldID, field.TypeUUID))
	)
	if id, ok := ic.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ic.mutation.DateCreated(); ok {
		_spec.SetField(identity.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := ic.mutation.Name(); ok {
		_spec.SetField(identity.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ic.mutation.Description(); ok {
		_spec.SetField(identity.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := ic.mutation.Kind(); ok {
		_spec.SetField(identity.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := ic.mutation.Verifier(); ok {
		_spec.SetField(identity.FieldVerifier, field.TypeString, value)
		_node.Verifier = value
	}
	if nodes := ic.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   identity.OwnerTable,
			Columns: []string{identity.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_identities = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IdentityCreateBulk is the builder for creating many Identity entities in bulk.
type IdentityCreateBulk struct {
	config
	err      error
	builders []*IdentityCreate
}

// Save creates the Identity entities in the database.
func (icb *IdentityCreateBulk) Save(ctx context.Context) ([]*Identity, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Identity, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IdentityMutation)
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
func (icb *IdentityCreateBulk) SaveX(ctx context.Context) []*Identity {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IdentityCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IdentityCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
