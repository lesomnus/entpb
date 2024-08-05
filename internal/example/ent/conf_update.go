// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lesomnus/entpb/internal/example/ent/conf"
	"github.com/lesomnus/entpb/internal/example/ent/predicate"
)

// ConfUpdate is the builder for updating Conf entities.
type ConfUpdate struct {
	config
	hooks    []Hook
	mutation *ConfMutation
}

// Where appends a list predicates to the ConfUpdate builder.
func (cu *ConfUpdate) Where(ps ...predicate.Conf) *ConfUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetValue sets the "value" field.
func (cu *ConfUpdate) SetValue(s string) *ConfUpdate {
	cu.mutation.SetValue(s)
	return cu
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cu *ConfUpdate) SetNillableValue(s *string) *ConfUpdate {
	if s != nil {
		cu.SetValue(*s)
	}
	return cu
}

// SetDateUpdated sets the "date_updated" field.
func (cu *ConfUpdate) SetDateUpdated(t time.Time) *ConfUpdate {
	cu.mutation.SetDateUpdated(t)
	return cu
}

// Mutation returns the ConfMutation object of the builder.
func (cu *ConfUpdate) Mutation() *ConfMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConfUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConfUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConfUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConfUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConfUpdate) defaults() {
	if _, ok := cu.mutation.DateUpdated(); !ok {
		v := conf.UpdateDefaultDateUpdated()
		cu.mutation.SetDateUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConfUpdate) check() error {
	if v, ok := cu.mutation.Value(); ok {
		if err := conf.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Conf.value": %w`, err)}
		}
	}
	return nil
}

func (cu *ConfUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(conf.Table, conf.Columns, sqlgraph.NewFieldSpec(conf.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Value(); ok {
		_spec.SetField(conf.FieldValue, field.TypeString, value)
	}
	if value, ok := cu.mutation.DateUpdated(); ok {
		_spec.SetField(conf.FieldDateUpdated, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConfUpdateOne is the builder for updating a single Conf entity.
type ConfUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConfMutation
}

// SetValue sets the "value" field.
func (cuo *ConfUpdateOne) SetValue(s string) *ConfUpdateOne {
	cuo.mutation.SetValue(s)
	return cuo
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (cuo *ConfUpdateOne) SetNillableValue(s *string) *ConfUpdateOne {
	if s != nil {
		cuo.SetValue(*s)
	}
	return cuo
}

// SetDateUpdated sets the "date_updated" field.
func (cuo *ConfUpdateOne) SetDateUpdated(t time.Time) *ConfUpdateOne {
	cuo.mutation.SetDateUpdated(t)
	return cuo
}

// Mutation returns the ConfMutation object of the builder.
func (cuo *ConfUpdateOne) Mutation() *ConfMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConfUpdate builder.
func (cuo *ConfUpdateOne) Where(ps ...predicate.Conf) *ConfUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConfUpdateOne) Select(field string, fields ...string) *ConfUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Conf entity.
func (cuo *ConfUpdateOne) Save(ctx context.Context) (*Conf, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConfUpdateOne) SaveX(ctx context.Context) *Conf {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConfUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConfUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConfUpdateOne) defaults() {
	if _, ok := cuo.mutation.DateUpdated(); !ok {
		v := conf.UpdateDefaultDateUpdated()
		cuo.mutation.SetDateUpdated(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConfUpdateOne) check() error {
	if v, ok := cuo.mutation.Value(); ok {
		if err := conf.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "Conf.value": %w`, err)}
		}
	}
	return nil
}

func (cuo *ConfUpdateOne) sqlSave(ctx context.Context) (_node *Conf, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(conf.Table, conf.Columns, sqlgraph.NewFieldSpec(conf.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Conf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conf.FieldID)
		for _, f := range fields {
			if !conf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != conf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Value(); ok {
		_spec.SetField(conf.FieldValue, field.TypeString, value)
	}
	if value, ok := cuo.mutation.DateUpdated(); ok {
		_spec.SetField(conf.FieldDateUpdated, field.TypeTime, value)
	}
	_node = &Conf{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
