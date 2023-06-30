// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/concourse/dex/storage/ent/db/connector"
	"github.com/concourse/dex/storage/ent/db/predicate"
)

// ConnectorUpdate is the builder for updating Connector entities.
type ConnectorUpdate struct {
	config
	hooks    []Hook
	mutation *ConnectorMutation
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cu *ConnectorUpdate) Where(ps ...predicate.Connector) *ConnectorUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetType sets the "type" field.
func (cu *ConnectorUpdate) SetType(s string) *ConnectorUpdate {
	cu.mutation.SetType(s)
	return cu
}

// SetName sets the "name" field.
func (cu *ConnectorUpdate) SetName(s string) *ConnectorUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetResourceVersion sets the "resource_version" field.
func (cu *ConnectorUpdate) SetResourceVersion(s string) *ConnectorUpdate {
	cu.mutation.SetResourceVersion(s)
	return cu
}

// SetConfig sets the "config" field.
func (cu *ConnectorUpdate) SetConfig(b []byte) *ConnectorUpdate {
	cu.mutation.SetConfig(b)
	return cu
}

// Mutation returns the ConnectorMutation object of the builder.
func (cu *ConnectorUpdate) Mutation() *ConnectorMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectorUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectorUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectorUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectorUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConnectorUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := connector.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "Connector.type": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := connector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Connector.name": %w`, err)}
		}
	}
	return nil
}

func (cu *ConnectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(connector.Table, connector.Columns, sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.SetField(connector.FieldType, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.ResourceVersion(); ok {
		_spec.SetField(connector.FieldResourceVersion, field.TypeString, value)
	}
	if value, ok := cu.mutation.Config(); ok {
		_spec.SetField(connector.FieldConfig, field.TypeBytes, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ConnectorUpdateOne is the builder for updating a single Connector entity.
type ConnectorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConnectorMutation
}

// SetType sets the "type" field.
func (cuo *ConnectorUpdateOne) SetType(s string) *ConnectorUpdateOne {
	cuo.mutation.SetType(s)
	return cuo
}

// SetName sets the "name" field.
func (cuo *ConnectorUpdateOne) SetName(s string) *ConnectorUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetResourceVersion sets the "resource_version" field.
func (cuo *ConnectorUpdateOne) SetResourceVersion(s string) *ConnectorUpdateOne {
	cuo.mutation.SetResourceVersion(s)
	return cuo
}

// SetConfig sets the "config" field.
func (cuo *ConnectorUpdateOne) SetConfig(b []byte) *ConnectorUpdateOne {
	cuo.mutation.SetConfig(b)
	return cuo
}

// Mutation returns the ConnectorMutation object of the builder.
func (cuo *ConnectorUpdateOne) Mutation() *ConnectorMutation {
	return cuo.mutation
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cuo *ConnectorUpdateOne) Where(ps ...predicate.Connector) *ConnectorUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectorUpdateOne) Select(field string, fields ...string) *ConnectorUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connector entity.
func (cuo *ConnectorUpdateOne) Save(ctx context.Context) (*Connector, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) SaveX(ctx context.Context) *Connector {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectorUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConnectorUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := connector.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`db: validator failed for field "Connector.type": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := connector.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Connector.name": %w`, err)}
		}
	}
	return nil
}

func (cuo *ConnectorUpdateOne) sqlSave(ctx context.Context) (_node *Connector, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(connector.Table, connector.Columns, sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "Connector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for _, f := range fields {
			if !connector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != connector.FieldID {
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
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.SetField(connector.FieldType, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(connector.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.ResourceVersion(); ok {
		_spec.SetField(connector.FieldResourceVersion, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Config(); ok {
		_spec.SetField(connector.FieldConfig, field.TypeBytes, value)
	}
	_node = &Connector{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
