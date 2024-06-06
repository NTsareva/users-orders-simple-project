// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NTsareva/orders-users-simple-service/cmd/order-service/ent/order"
	"github.com/NTsareva/orders-users-simple-service/cmd/order-service/ent/predicate"
)

// OrderUpdate is the builder for updating Order entities.
type OrderUpdate struct {
	config
	hooks    []Hook
	mutation *OrderMutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ou *OrderUpdate) Where(ps ...predicate.Order) *OrderUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// SetTitle sets the "title" field.
func (ou *OrderUpdate) SetTitle(s string) *OrderUpdate {
	ou.mutation.SetTitle(s)
	return ou
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableTitle(s *string) *OrderUpdate {
	if s != nil {
		ou.SetTitle(*s)
	}
	return ou
}

// SetDescription sets the "description" field.
func (ou *OrderUpdate) SetDescription(s string) *OrderUpdate {
	ou.mutation.SetDescription(s)
	return ou
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableDescription(s *string) *OrderUpdate {
	if s != nil {
		ou.SetDescription(*s)
	}
	return ou
}

// SetUserID sets the "user_id" field.
func (ou *OrderUpdate) SetUserID(i int) *OrderUpdate {
	ou.mutation.ResetUserID()
	ou.mutation.SetUserID(i)
	return ou
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ou *OrderUpdate) SetNillableUserID(i *int) *OrderUpdate {
	if i != nil {
		ou.SetUserID(*i)
	}
	return ou
}

// AddUserID adds i to the "user_id" field.
func (ou *OrderUpdate) AddUserID(i int) *OrderUpdate {
	ou.mutation.AddUserID(i)
	return ou
}

// Mutation returns the OrderMutation object of the builder.
func (ou *OrderUpdate) Mutation() *OrderMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OrderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OrderUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OrderUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OrderUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ou *OrderUpdate) check() error {
	if v, ok := ou.mutation.Title(); ok {
		if err := order.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Order.title": %w`, err)}
		}
	}
	if v, ok := ou.mutation.Description(); ok {
		if err := order.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Order.description": %w`, err)}
		}
	}
	if v, ok := ou.mutation.UserID(); ok {
		if err := order.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Order.user_id": %w`, err)}
		}
	}
	return nil
}

func (ou *OrderUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ou.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ou.mutation.Title(); ok {
		_spec.SetField(order.FieldTitle, field.TypeString, value)
	}
	if value, ok := ou.mutation.Description(); ok {
		_spec.SetField(order.FieldDescription, field.TypeString, value)
	}
	if value, ok := ou.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ou.mutation.AddedUserID(); ok {
		_spec.AddField(order.FieldUserID, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OrderUpdateOne is the builder for updating a single Order entity.
type OrderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrderMutation
}

// SetTitle sets the "title" field.
func (ouo *OrderUpdateOne) SetTitle(s string) *OrderUpdateOne {
	ouo.mutation.SetTitle(s)
	return ouo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableTitle(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetTitle(*s)
	}
	return ouo
}

// SetDescription sets the "description" field.
func (ouo *OrderUpdateOne) SetDescription(s string) *OrderUpdateOne {
	ouo.mutation.SetDescription(s)
	return ouo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableDescription(s *string) *OrderUpdateOne {
	if s != nil {
		ouo.SetDescription(*s)
	}
	return ouo
}

// SetUserID sets the "user_id" field.
func (ouo *OrderUpdateOne) SetUserID(i int) *OrderUpdateOne {
	ouo.mutation.ResetUserID()
	ouo.mutation.SetUserID(i)
	return ouo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ouo *OrderUpdateOne) SetNillableUserID(i *int) *OrderUpdateOne {
	if i != nil {
		ouo.SetUserID(*i)
	}
	return ouo
}

// AddUserID adds i to the "user_id" field.
func (ouo *OrderUpdateOne) AddUserID(i int) *OrderUpdateOne {
	ouo.mutation.AddUserID(i)
	return ouo
}

// Mutation returns the OrderMutation object of the builder.
func (ouo *OrderUpdateOne) Mutation() *OrderMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OrderUpdate builder.
func (ouo *OrderUpdateOne) Where(ps ...predicate.Order) *OrderUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OrderUpdateOne) Select(field string, fields ...string) *OrderUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Order entity.
func (ouo *OrderUpdateOne) Save(ctx context.Context) (*Order, error) {
	return withHooks(ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OrderUpdateOne) SaveX(ctx context.Context) *Order {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OrderUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OrderUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ouo *OrderUpdateOne) check() error {
	if v, ok := ouo.mutation.Title(); ok {
		if err := order.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Order.title": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.Description(); ok {
		if err := order.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Order.description": %w`, err)}
		}
	}
	if v, ok := ouo.mutation.UserID(); ok {
		if err := order.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Order.user_id": %w`, err)}
		}
	}
	return nil
}

func (ouo *OrderUpdateOne) sqlSave(ctx context.Context) (_node *Order, err error) {
	if err := ouo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(order.Table, order.Columns, sqlgraph.NewFieldSpec(order.FieldID, field.TypeString))
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Order.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, order.FieldID)
		for _, f := range fields {
			if !order.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != order.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ouo.mutation.Title(); ok {
		_spec.SetField(order.FieldTitle, field.TypeString, value)
	}
	if value, ok := ouo.mutation.Description(); ok {
		_spec.SetField(order.FieldDescription, field.TypeString, value)
	}
	if value, ok := ouo.mutation.UserID(); ok {
		_spec.SetField(order.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ouo.mutation.AddedUserID(); ok {
		_spec.AddField(order.FieldUserID, field.TypeInt, value)
	}
	_node = &Order{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{order.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
