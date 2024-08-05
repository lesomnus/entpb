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
	"github.com/lesomnus/entpb/internal/example/ent/invitation"
	"github.com/lesomnus/entpb/internal/example/ent/predicate"
)

// InvitationUpdate is the builder for updating Invitation entities.
type InvitationUpdate struct {
	config
	hooks    []Hook
	mutation *InvitationMutation
}

// Where appends a list predicates to the InvitationUpdate builder.
func (iu *InvitationUpdate) Where(ps ...predicate.Invitation) *InvitationUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetDateExpired sets the "date_expired" field.
func (iu *InvitationUpdate) SetDateExpired(t time.Time) *InvitationUpdate {
	iu.mutation.SetDateExpired(t)
	return iu
}

// SetNillableDateExpired sets the "date_expired" field if the given value is not nil.
func (iu *InvitationUpdate) SetNillableDateExpired(t *time.Time) *InvitationUpdate {
	if t != nil {
		iu.SetDateExpired(*t)
	}
	return iu
}

// SetDateAccepted sets the "date_accepted" field.
func (iu *InvitationUpdate) SetDateAccepted(t time.Time) *InvitationUpdate {
	iu.mutation.SetDateAccepted(t)
	return iu
}

// SetNillableDateAccepted sets the "date_accepted" field if the given value is not nil.
func (iu *InvitationUpdate) SetNillableDateAccepted(t *time.Time) *InvitationUpdate {
	if t != nil {
		iu.SetDateAccepted(*t)
	}
	return iu
}

// ClearDateAccepted clears the value of the "date_accepted" field.
func (iu *InvitationUpdate) ClearDateAccepted() *InvitationUpdate {
	iu.mutation.ClearDateAccepted()
	return iu
}

// SetDateDeclined sets the "date_declined" field.
func (iu *InvitationUpdate) SetDateDeclined(t time.Time) *InvitationUpdate {
	iu.mutation.SetDateDeclined(t)
	return iu
}

// SetNillableDateDeclined sets the "date_declined" field if the given value is not nil.
func (iu *InvitationUpdate) SetNillableDateDeclined(t *time.Time) *InvitationUpdate {
	if t != nil {
		iu.SetDateDeclined(*t)
	}
	return iu
}

// ClearDateDeclined clears the value of the "date_declined" field.
func (iu *InvitationUpdate) ClearDateDeclined() *InvitationUpdate {
	iu.mutation.ClearDateDeclined()
	return iu
}

// SetDateCanceled sets the "date_canceled" field.
func (iu *InvitationUpdate) SetDateCanceled(t time.Time) *InvitationUpdate {
	iu.mutation.SetDateCanceled(t)
	return iu
}

// SetNillableDateCanceled sets the "date_canceled" field if the given value is not nil.
func (iu *InvitationUpdate) SetNillableDateCanceled(t *time.Time) *InvitationUpdate {
	if t != nil {
		iu.SetDateCanceled(*t)
	}
	return iu
}

// ClearDateCanceled clears the value of the "date_canceled" field.
func (iu *InvitationUpdate) ClearDateCanceled() *InvitationUpdate {
	iu.mutation.ClearDateCanceled()
	return iu
}

// Mutation returns the InvitationMutation object of the builder.
func (iu *InvitationUpdate) Mutation() *InvitationMutation {
	return iu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *InvitationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *InvitationUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *InvitationUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *InvitationUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *InvitationUpdate) check() error {
	if _, ok := iu.mutation.InviterID(); iu.mutation.InviterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invitation.inviter"`)
	}
	if _, ok := iu.mutation.SiloID(); iu.mutation.SiloCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invitation.silo"`)
	}
	return nil
}

func (iu *InvitationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(invitation.Table, invitation.Columns, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeUUID))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.DateExpired(); ok {
		_spec.SetField(invitation.FieldDateExpired, field.TypeTime, value)
	}
	if value, ok := iu.mutation.DateAccepted(); ok {
		_spec.SetField(invitation.FieldDateAccepted, field.TypeTime, value)
	}
	if iu.mutation.DateAcceptedCleared() {
		_spec.ClearField(invitation.FieldDateAccepted, field.TypeTime)
	}
	if value, ok := iu.mutation.DateDeclined(); ok {
		_spec.SetField(invitation.FieldDateDeclined, field.TypeTime, value)
	}
	if iu.mutation.DateDeclinedCleared() {
		_spec.ClearField(invitation.FieldDateDeclined, field.TypeTime)
	}
	if value, ok := iu.mutation.DateCanceled(); ok {
		_spec.SetField(invitation.FieldDateCanceled, field.TypeTime, value)
	}
	if iu.mutation.DateCanceledCleared() {
		_spec.ClearField(invitation.FieldDateCanceled, field.TypeTime)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// InvitationUpdateOne is the builder for updating a single Invitation entity.
type InvitationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InvitationMutation
}

// SetDateExpired sets the "date_expired" field.
func (iuo *InvitationUpdateOne) SetDateExpired(t time.Time) *InvitationUpdateOne {
	iuo.mutation.SetDateExpired(t)
	return iuo
}

// SetNillableDateExpired sets the "date_expired" field if the given value is not nil.
func (iuo *InvitationUpdateOne) SetNillableDateExpired(t *time.Time) *InvitationUpdateOne {
	if t != nil {
		iuo.SetDateExpired(*t)
	}
	return iuo
}

// SetDateAccepted sets the "date_accepted" field.
func (iuo *InvitationUpdateOne) SetDateAccepted(t time.Time) *InvitationUpdateOne {
	iuo.mutation.SetDateAccepted(t)
	return iuo
}

// SetNillableDateAccepted sets the "date_accepted" field if the given value is not nil.
func (iuo *InvitationUpdateOne) SetNillableDateAccepted(t *time.Time) *InvitationUpdateOne {
	if t != nil {
		iuo.SetDateAccepted(*t)
	}
	return iuo
}

// ClearDateAccepted clears the value of the "date_accepted" field.
func (iuo *InvitationUpdateOne) ClearDateAccepted() *InvitationUpdateOne {
	iuo.mutation.ClearDateAccepted()
	return iuo
}

// SetDateDeclined sets the "date_declined" field.
func (iuo *InvitationUpdateOne) SetDateDeclined(t time.Time) *InvitationUpdateOne {
	iuo.mutation.SetDateDeclined(t)
	return iuo
}

// SetNillableDateDeclined sets the "date_declined" field if the given value is not nil.
func (iuo *InvitationUpdateOne) SetNillableDateDeclined(t *time.Time) *InvitationUpdateOne {
	if t != nil {
		iuo.SetDateDeclined(*t)
	}
	return iuo
}

// ClearDateDeclined clears the value of the "date_declined" field.
func (iuo *InvitationUpdateOne) ClearDateDeclined() *InvitationUpdateOne {
	iuo.mutation.ClearDateDeclined()
	return iuo
}

// SetDateCanceled sets the "date_canceled" field.
func (iuo *InvitationUpdateOne) SetDateCanceled(t time.Time) *InvitationUpdateOne {
	iuo.mutation.SetDateCanceled(t)
	return iuo
}

// SetNillableDateCanceled sets the "date_canceled" field if the given value is not nil.
func (iuo *InvitationUpdateOne) SetNillableDateCanceled(t *time.Time) *InvitationUpdateOne {
	if t != nil {
		iuo.SetDateCanceled(*t)
	}
	return iuo
}

// ClearDateCanceled clears the value of the "date_canceled" field.
func (iuo *InvitationUpdateOne) ClearDateCanceled() *InvitationUpdateOne {
	iuo.mutation.ClearDateCanceled()
	return iuo
}

// Mutation returns the InvitationMutation object of the builder.
func (iuo *InvitationUpdateOne) Mutation() *InvitationMutation {
	return iuo.mutation
}

// Where appends a list predicates to the InvitationUpdate builder.
func (iuo *InvitationUpdateOne) Where(ps ...predicate.Invitation) *InvitationUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *InvitationUpdateOne) Select(field string, fields ...string) *InvitationUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Invitation entity.
func (iuo *InvitationUpdateOne) Save(ctx context.Context) (*Invitation, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *InvitationUpdateOne) SaveX(ctx context.Context) *Invitation {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *InvitationUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *InvitationUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *InvitationUpdateOne) check() error {
	if _, ok := iuo.mutation.InviterID(); iuo.mutation.InviterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invitation.inviter"`)
	}
	if _, ok := iuo.mutation.SiloID(); iuo.mutation.SiloCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Invitation.silo"`)
	}
	return nil
}

func (iuo *InvitationUpdateOne) sqlSave(ctx context.Context) (_node *Invitation, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(invitation.Table, invitation.Columns, sqlgraph.NewFieldSpec(invitation.FieldID, field.TypeUUID))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Invitation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, invitation.FieldID)
		for _, f := range fields {
			if !invitation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != invitation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.DateExpired(); ok {
		_spec.SetField(invitation.FieldDateExpired, field.TypeTime, value)
	}
	if value, ok := iuo.mutation.DateAccepted(); ok {
		_spec.SetField(invitation.FieldDateAccepted, field.TypeTime, value)
	}
	if iuo.mutation.DateAcceptedCleared() {
		_spec.ClearField(invitation.FieldDateAccepted, field.TypeTime)
	}
	if value, ok := iuo.mutation.DateDeclined(); ok {
		_spec.SetField(invitation.FieldDateDeclined, field.TypeTime, value)
	}
	if iuo.mutation.DateDeclinedCleared() {
		_spec.ClearField(invitation.FieldDateDeclined, field.TypeTime)
	}
	if value, ok := iuo.mutation.DateCanceled(); ok {
		_spec.SetField(invitation.FieldDateCanceled, field.TypeTime, value)
	}
	if iuo.mutation.DateCanceledCleared() {
		_spec.ClearField(invitation.FieldDateCanceled, field.TypeTime)
	}
	_node = &Invitation{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{invitation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}
