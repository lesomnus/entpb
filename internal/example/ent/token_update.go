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
	"github.com/google/uuid"
	"github.com/lesomnus/entpb/internal/example/ent/predicate"
	"github.com/lesomnus/entpb/internal/example/ent/token"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks    []Hook
	mutation *TokenMutation
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetName sets the "name" field.
func (tu *TokenUpdate) SetName(s string) *TokenUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableName(s *string) *TokenUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetUseCountLimit sets the "use_count_limit" field.
func (tu *TokenUpdate) SetUseCountLimit(u uint64) *TokenUpdate {
	tu.mutation.ResetUseCountLimit()
	tu.mutation.SetUseCountLimit(u)
	return tu
}

// SetNillableUseCountLimit sets the "use_count_limit" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableUseCountLimit(u *uint64) *TokenUpdate {
	if u != nil {
		tu.SetUseCountLimit(*u)
	}
	return tu
}

// AddUseCountLimit adds u to the "use_count_limit" field.
func (tu *TokenUpdate) AddUseCountLimit(u int64) *TokenUpdate {
	tu.mutation.AddUseCountLimit(u)
	return tu
}

// SetDateExpired sets the "date_expired" field.
func (tu *TokenUpdate) SetDateExpired(t time.Time) *TokenUpdate {
	tu.mutation.SetDateExpired(t)
	return tu
}

// SetNillableDateExpired sets the "date_expired" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableDateExpired(t *time.Time) *TokenUpdate {
	if t != nil {
		tu.SetDateExpired(*t)
	}
	return tu
}

// AddChildIDs adds the "children" edge to the Token entity by IDs.
func (tu *TokenUpdate) AddChildIDs(ids ...uuid.UUID) *TokenUpdate {
	tu.mutation.AddChildIDs(ids...)
	return tu
}

// AddChildren adds the "children" edges to the Token entity.
func (tu *TokenUpdate) AddChildren(t ...*Token) *TokenUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddChildIDs(ids...)
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// ClearChildren clears all "children" edges to the Token entity.
func (tu *TokenUpdate) ClearChildren() *TokenUpdate {
	tu.mutation.ClearChildren()
	return tu
}

// RemoveChildIDs removes the "children" edge to Token entities by IDs.
func (tu *TokenUpdate) RemoveChildIDs(ids ...uuid.UUID) *TokenUpdate {
	tu.mutation.RemoveChildIDs(ids...)
	return tu
}

// RemoveChildren removes "children" edges to Token entities.
func (tu *TokenUpdate) RemoveChildren(t ...*Token) *TokenUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if _, ok := tu.mutation.OwnerID(); tu.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Token.owner"`)
	}
	return nil
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(token.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.UseCountLimit(); ok {
		_spec.SetField(token.FieldUseCountLimit, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.AddedUseCountLimit(); ok {
		_spec.AddField(token.FieldUseCountLimit, field.TypeUint64, value)
	}
	if value, ok := tu.mutation.DateExpired(); ok {
		_spec.SetField(token.FieldDateExpired, field.TypeTime, value)
	}
	if tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TokenMutation
}

// SetName sets the "name" field.
func (tuo *TokenUpdateOne) SetName(s string) *TokenUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableName(s *string) *TokenUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetUseCountLimit sets the "use_count_limit" field.
func (tuo *TokenUpdateOne) SetUseCountLimit(u uint64) *TokenUpdateOne {
	tuo.mutation.ResetUseCountLimit()
	tuo.mutation.SetUseCountLimit(u)
	return tuo
}

// SetNillableUseCountLimit sets the "use_count_limit" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableUseCountLimit(u *uint64) *TokenUpdateOne {
	if u != nil {
		tuo.SetUseCountLimit(*u)
	}
	return tuo
}

// AddUseCountLimit adds u to the "use_count_limit" field.
func (tuo *TokenUpdateOne) AddUseCountLimit(u int64) *TokenUpdateOne {
	tuo.mutation.AddUseCountLimit(u)
	return tuo
}

// SetDateExpired sets the "date_expired" field.
func (tuo *TokenUpdateOne) SetDateExpired(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetDateExpired(t)
	return tuo
}

// SetNillableDateExpired sets the "date_expired" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableDateExpired(t *time.Time) *TokenUpdateOne {
	if t != nil {
		tuo.SetDateExpired(*t)
	}
	return tuo
}

// AddChildIDs adds the "children" edge to the Token entity by IDs.
func (tuo *TokenUpdateOne) AddChildIDs(ids ...uuid.UUID) *TokenUpdateOne {
	tuo.mutation.AddChildIDs(ids...)
	return tuo
}

// AddChildren adds the "children" edges to the Token entity.
func (tuo *TokenUpdateOne) AddChildren(t ...*Token) *TokenUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddChildIDs(ids...)
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// ClearChildren clears all "children" edges to the Token entity.
func (tuo *TokenUpdateOne) ClearChildren() *TokenUpdateOne {
	tuo.mutation.ClearChildren()
	return tuo
}

// RemoveChildIDs removes the "children" edge to Token entities by IDs.
func (tuo *TokenUpdateOne) RemoveChildIDs(ids ...uuid.UUID) *TokenUpdateOne {
	tuo.mutation.RemoveChildIDs(ids...)
	return tuo
}

// RemoveChildren removes "children" edges to Token entities.
func (tuo *TokenUpdateOne) RemoveChildren(t ...*Token) *TokenUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the TokenUpdate builder.
func (tuo *TokenUpdateOne) Where(ps ...predicate.Token) *TokenUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if _, ok := tuo.mutation.OwnerID(); tuo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Token.owner"`)
	}
	return nil
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(token.Table, token.Columns, sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(token.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.UseCountLimit(); ok {
		_spec.SetField(token.FieldUseCountLimit, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.AddedUseCountLimit(); ok {
		_spec.AddField(token.FieldUseCountLimit, field.TypeUint64, value)
	}
	if value, ok := tuo.mutation.DateExpired(); ok {
		_spec.SetField(token.FieldDateExpired, field.TypeTime, value)
	}
	if tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   token.ChildrenTable,
			Columns: []string{token.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
