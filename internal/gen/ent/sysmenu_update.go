// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
)

// SysMenuUpdate is the builder for updating SysMenu entities.
type SysMenuUpdate struct {
	config
	hooks    []Hook
	mutation *SysMenuMutation
}

// Where adds a new predicate for the SysMenuUpdate builder.
func (smu *SysMenuUpdate) Where(ps ...predicate.SysMenu) *SysMenuUpdate {
	smu.mutation.predicates = append(smu.mutation.predicates, ps...)
	return smu
}

// SetIsDel sets the "is_del" field.
func (smu *SysMenuUpdate) SetIsDel(b bool) *SysMenuUpdate {
	smu.mutation.SetIsDel(b)
	return smu
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableIsDel(b *bool) *SysMenuUpdate {
	if b != nil {
		smu.SetIsDel(*b)
	}
	return smu
}

// SetMemo sets the "memo" field.
func (smu *SysMenuUpdate) SetMemo(s string) *SysMenuUpdate {
	smu.mutation.SetMemo(s)
	return smu
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableMemo(s *string) *SysMenuUpdate {
	if s != nil {
		smu.SetMemo(*s)
	}
	return smu
}

// SetSort sets the "sort" field.
func (smu *SysMenuUpdate) SetSort(i int32) *SysMenuUpdate {
	smu.mutation.ResetSort()
	smu.mutation.SetSort(i)
	return smu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableSort(i *int32) *SysMenuUpdate {
	if i != nil {
		smu.SetSort(*i)
	}
	return smu
}

// AddSort adds i to the "sort" field.
func (smu *SysMenuUpdate) AddSort(i int32) *SysMenuUpdate {
	smu.mutation.AddSort(i)
	return smu
}

// SetUpdatedAt sets the "updated_at" field.
func (smu *SysMenuUpdate) SetUpdatedAt(t time.Time) *SysMenuUpdate {
	smu.mutation.SetUpdatedAt(t)
	return smu
}

// SetDeletedAt sets the "deleted_at" field.
func (smu *SysMenuUpdate) SetDeletedAt(t time.Time) *SysMenuUpdate {
	smu.mutation.SetDeletedAt(t)
	return smu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableDeletedAt(t *time.Time) *SysMenuUpdate {
	if t != nil {
		smu.SetDeletedAt(*t)
	}
	return smu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smu *SysMenuUpdate) ClearDeletedAt() *SysMenuUpdate {
	smu.mutation.ClearDeletedAt()
	return smu
}

// SetStatus sets the "status" field.
func (smu *SysMenuUpdate) SetStatus(i int16) *SysMenuUpdate {
	smu.mutation.ResetStatus()
	smu.mutation.SetStatus(i)
	return smu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableStatus(i *int16) *SysMenuUpdate {
	if i != nil {
		smu.SetStatus(*i)
	}
	return smu
}

// AddStatus adds i to the "status" field.
func (smu *SysMenuUpdate) AddStatus(i int16) *SysMenuUpdate {
	smu.mutation.AddStatus(i)
	return smu
}

// SetName sets the "name" field.
func (smu *SysMenuUpdate) SetName(s string) *SysMenuUpdate {
	smu.mutation.SetName(s)
	return smu
}

// SetIcon sets the "icon" field.
func (smu *SysMenuUpdate) SetIcon(s string) *SysMenuUpdate {
	smu.mutation.SetIcon(s)
	return smu
}

// SetRouter sets the "router" field.
func (smu *SysMenuUpdate) SetRouter(s string) *SysMenuUpdate {
	smu.mutation.SetRouter(s)
	return smu
}

// SetIsShow sets the "is_show" field.
func (smu *SysMenuUpdate) SetIsShow(b bool) *SysMenuUpdate {
	smu.mutation.SetIsShow(b)
	return smu
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableIsShow(b *bool) *SysMenuUpdate {
	if b != nil {
		smu.SetIsShow(*b)
	}
	return smu
}

// SetParentID sets the "parent_id" field.
func (smu *SysMenuUpdate) SetParentID(s string) *SysMenuUpdate {
	smu.mutation.SetParentID(s)
	return smu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableParentID(s *string) *SysMenuUpdate {
	if s != nil {
		smu.SetParentID(*s)
	}
	return smu
}

// ClearParentID clears the value of the "parent_id" field.
func (smu *SysMenuUpdate) ClearParentID() *SysMenuUpdate {
	smu.mutation.ClearParentID()
	return smu
}

// SetParentPath sets the "parent_path" field.
func (smu *SysMenuUpdate) SetParentPath(s string) *SysMenuUpdate {
	smu.mutation.SetParentPath(s)
	return smu
}

// SetNillableParentPath sets the "parent_path" field if the given value is not nil.
func (smu *SysMenuUpdate) SetNillableParentPath(s *string) *SysMenuUpdate {
	if s != nil {
		smu.SetParentPath(*s)
	}
	return smu
}

// ClearParentPath clears the value of the "parent_path" field.
func (smu *SysMenuUpdate) ClearParentPath() *SysMenuUpdate {
	smu.mutation.ClearParentPath()
	return smu
}

// Mutation returns the SysMenuMutation object of the builder.
func (smu *SysMenuUpdate) Mutation() *SysMenuMutation {
	return smu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (smu *SysMenuUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	smu.defaults()
	if len(smu.hooks) == 0 {
		if err = smu.check(); err != nil {
			return 0, err
		}
		affected, err = smu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smu.check(); err != nil {
				return 0, err
			}
			smu.mutation = mutation
			affected, err = smu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(smu.hooks) - 1; i >= 0; i-- {
			mut = smu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (smu *SysMenuUpdate) SaveX(ctx context.Context) int {
	affected, err := smu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (smu *SysMenuUpdate) Exec(ctx context.Context) error {
	_, err := smu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smu *SysMenuUpdate) ExecX(ctx context.Context) {
	if err := smu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smu *SysMenuUpdate) defaults() {
	if _, ok := smu.mutation.UpdatedAt(); !ok {
		v := sysmenu.UpdateDefaultUpdatedAt()
		smu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smu *SysMenuUpdate) check() error {
	if v, ok := smu.mutation.Memo(); ok {
		if err := sysmenu.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := smu.mutation.Name(); ok {
		if err := sysmenu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := smu.mutation.Icon(); ok {
		if err := sysmenu.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	if v, ok := smu.mutation.Router(); ok {
		if err := sysmenu.RouterValidator(v); err != nil {
			return &ValidationError{Name: "router", err: fmt.Errorf("ent: validator failed for field \"router\": %w", err)}
		}
	}
	if v, ok := smu.mutation.ParentID(); ok {
		if err := sysmenu.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf("ent: validator failed for field \"parent_id\": %w", err)}
		}
	}
	if v, ok := smu.mutation.ParentPath(); ok {
		if err := sysmenu.ParentPathValidator(v); err != nil {
			return &ValidationError{Name: "parent_path", err: fmt.Errorf("ent: validator failed for field \"parent_path\": %w", err)}
		}
	}
	return nil
}

func (smu *SysMenuUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenu.FieldID,
			},
		},
	}
	if ps := smu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smu.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsDel,
		})
	}
	if value, ok := smu.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMemo,
		})
	}
	if value, ok := smu.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
	}
	if value, ok := smu.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
	}
	if value, ok := smu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
	}
	if value, ok := smu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if smu.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if value, ok := smu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysmenu.FieldStatus,
		})
	}
	if value, ok := smu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysmenu.FieldStatus,
		})
	}
	if value, ok := smu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
	}
	if value, ok := smu.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldIcon,
		})
	}
	if value, ok := smu.mutation.Router(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRouter,
		})
	}
	if value, ok := smu.mutation.IsShow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsShow,
		})
	}
	if value, ok := smu.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if smu.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smu.mutation.ParentPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentPath,
		})
	}
	if smu.mutation.ParentPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldParentPath,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, smu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SysMenuUpdateOne is the builder for updating a single SysMenu entity.
type SysMenuUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SysMenuMutation
}

// SetIsDel sets the "is_del" field.
func (smuo *SysMenuUpdateOne) SetIsDel(b bool) *SysMenuUpdateOne {
	smuo.mutation.SetIsDel(b)
	return smuo
}

// SetNillableIsDel sets the "is_del" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableIsDel(b *bool) *SysMenuUpdateOne {
	if b != nil {
		smuo.SetIsDel(*b)
	}
	return smuo
}

// SetMemo sets the "memo" field.
func (smuo *SysMenuUpdateOne) SetMemo(s string) *SysMenuUpdateOne {
	smuo.mutation.SetMemo(s)
	return smuo
}

// SetNillableMemo sets the "memo" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableMemo(s *string) *SysMenuUpdateOne {
	if s != nil {
		smuo.SetMemo(*s)
	}
	return smuo
}

// SetSort sets the "sort" field.
func (smuo *SysMenuUpdateOne) SetSort(i int32) *SysMenuUpdateOne {
	smuo.mutation.ResetSort()
	smuo.mutation.SetSort(i)
	return smuo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableSort(i *int32) *SysMenuUpdateOne {
	if i != nil {
		smuo.SetSort(*i)
	}
	return smuo
}

// AddSort adds i to the "sort" field.
func (smuo *SysMenuUpdateOne) AddSort(i int32) *SysMenuUpdateOne {
	smuo.mutation.AddSort(i)
	return smuo
}

// SetUpdatedAt sets the "updated_at" field.
func (smuo *SysMenuUpdateOne) SetUpdatedAt(t time.Time) *SysMenuUpdateOne {
	smuo.mutation.SetUpdatedAt(t)
	return smuo
}

// SetDeletedAt sets the "deleted_at" field.
func (smuo *SysMenuUpdateOne) SetDeletedAt(t time.Time) *SysMenuUpdateOne {
	smuo.mutation.SetDeletedAt(t)
	return smuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableDeletedAt(t *time.Time) *SysMenuUpdateOne {
	if t != nil {
		smuo.SetDeletedAt(*t)
	}
	return smuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (smuo *SysMenuUpdateOne) ClearDeletedAt() *SysMenuUpdateOne {
	smuo.mutation.ClearDeletedAt()
	return smuo
}

// SetStatus sets the "status" field.
func (smuo *SysMenuUpdateOne) SetStatus(i int16) *SysMenuUpdateOne {
	smuo.mutation.ResetStatus()
	smuo.mutation.SetStatus(i)
	return smuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableStatus(i *int16) *SysMenuUpdateOne {
	if i != nil {
		smuo.SetStatus(*i)
	}
	return smuo
}

// AddStatus adds i to the "status" field.
func (smuo *SysMenuUpdateOne) AddStatus(i int16) *SysMenuUpdateOne {
	smuo.mutation.AddStatus(i)
	return smuo
}

// SetName sets the "name" field.
func (smuo *SysMenuUpdateOne) SetName(s string) *SysMenuUpdateOne {
	smuo.mutation.SetName(s)
	return smuo
}

// SetIcon sets the "icon" field.
func (smuo *SysMenuUpdateOne) SetIcon(s string) *SysMenuUpdateOne {
	smuo.mutation.SetIcon(s)
	return smuo
}

// SetRouter sets the "router" field.
func (smuo *SysMenuUpdateOne) SetRouter(s string) *SysMenuUpdateOne {
	smuo.mutation.SetRouter(s)
	return smuo
}

// SetIsShow sets the "is_show" field.
func (smuo *SysMenuUpdateOne) SetIsShow(b bool) *SysMenuUpdateOne {
	smuo.mutation.SetIsShow(b)
	return smuo
}

// SetNillableIsShow sets the "is_show" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableIsShow(b *bool) *SysMenuUpdateOne {
	if b != nil {
		smuo.SetIsShow(*b)
	}
	return smuo
}

// SetParentID sets the "parent_id" field.
func (smuo *SysMenuUpdateOne) SetParentID(s string) *SysMenuUpdateOne {
	smuo.mutation.SetParentID(s)
	return smuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableParentID(s *string) *SysMenuUpdateOne {
	if s != nil {
		smuo.SetParentID(*s)
	}
	return smuo
}

// ClearParentID clears the value of the "parent_id" field.
func (smuo *SysMenuUpdateOne) ClearParentID() *SysMenuUpdateOne {
	smuo.mutation.ClearParentID()
	return smuo
}

// SetParentPath sets the "parent_path" field.
func (smuo *SysMenuUpdateOne) SetParentPath(s string) *SysMenuUpdateOne {
	smuo.mutation.SetParentPath(s)
	return smuo
}

// SetNillableParentPath sets the "parent_path" field if the given value is not nil.
func (smuo *SysMenuUpdateOne) SetNillableParentPath(s *string) *SysMenuUpdateOne {
	if s != nil {
		smuo.SetParentPath(*s)
	}
	return smuo
}

// ClearParentPath clears the value of the "parent_path" field.
func (smuo *SysMenuUpdateOne) ClearParentPath() *SysMenuUpdateOne {
	smuo.mutation.ClearParentPath()
	return smuo
}

// Mutation returns the SysMenuMutation object of the builder.
func (smuo *SysMenuUpdateOne) Mutation() *SysMenuMutation {
	return smuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (smuo *SysMenuUpdateOne) Select(field string, fields ...string) *SysMenuUpdateOne {
	smuo.fields = append([]string{field}, fields...)
	return smuo
}

// Save executes the query and returns the updated SysMenu entity.
func (smuo *SysMenuUpdateOne) Save(ctx context.Context) (*SysMenu, error) {
	var (
		err  error
		node *SysMenu
	)
	smuo.defaults()
	if len(smuo.hooks) == 0 {
		if err = smuo.check(); err != nil {
			return nil, err
		}
		node, err = smuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SysMenuMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = smuo.check(); err != nil {
				return nil, err
			}
			smuo.mutation = mutation
			node, err = smuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(smuo.hooks) - 1; i >= 0; i-- {
			mut = smuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, smuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (smuo *SysMenuUpdateOne) SaveX(ctx context.Context) *SysMenu {
	node, err := smuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (smuo *SysMenuUpdateOne) Exec(ctx context.Context) error {
	_, err := smuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smuo *SysMenuUpdateOne) ExecX(ctx context.Context) {
	if err := smuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smuo *SysMenuUpdateOne) defaults() {
	if _, ok := smuo.mutation.UpdatedAt(); !ok {
		v := sysmenu.UpdateDefaultUpdatedAt()
		smuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smuo *SysMenuUpdateOne) check() error {
	if v, ok := smuo.mutation.Memo(); ok {
		if err := sysmenu.MemoValidator(v); err != nil {
			return &ValidationError{Name: "memo", err: fmt.Errorf("ent: validator failed for field \"memo\": %w", err)}
		}
	}
	if v, ok := smuo.mutation.Name(); ok {
		if err := sysmenu.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := smuo.mutation.Icon(); ok {
		if err := sysmenu.IconValidator(v); err != nil {
			return &ValidationError{Name: "icon", err: fmt.Errorf("ent: validator failed for field \"icon\": %w", err)}
		}
	}
	if v, ok := smuo.mutation.Router(); ok {
		if err := sysmenu.RouterValidator(v); err != nil {
			return &ValidationError{Name: "router", err: fmt.Errorf("ent: validator failed for field \"router\": %w", err)}
		}
	}
	if v, ok := smuo.mutation.ParentID(); ok {
		if err := sysmenu.ParentIDValidator(v); err != nil {
			return &ValidationError{Name: "parent_id", err: fmt.Errorf("ent: validator failed for field \"parent_id\": %w", err)}
		}
	}
	if v, ok := smuo.mutation.ParentPath(); ok {
		if err := sysmenu.ParentPathValidator(v); err != nil {
			return &ValidationError{Name: "parent_path", err: fmt.Errorf("ent: validator failed for field \"parent_path\": %w", err)}
		}
	}
	return nil
}

func (smuo *SysMenuUpdateOne) sqlSave(ctx context.Context) (_node *SysMenu, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenu.Table,
			Columns: sysmenu.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenu.FieldID,
			},
		},
	}
	id, ok := smuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SysMenu.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := smuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenu.FieldID)
		for _, f := range fields {
			if !sysmenu.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != sysmenu.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := smuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := smuo.mutation.IsDel(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsDel,
		})
	}
	if value, ok := smuo.mutation.Memo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldMemo,
		})
	}
	if value, ok := smuo.mutation.Sort(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
	}
	if value, ok := smuo.mutation.AddedSort(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: sysmenu.FieldSort,
		})
	}
	if value, ok := smuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldUpdatedAt,
		})
	}
	if value, ok := smuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if smuo.mutation.DeletedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: sysmenu.FieldDeletedAt,
		})
	}
	if value, ok := smuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysmenu.FieldStatus,
		})
	}
	if value, ok := smuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt16,
			Value:  value,
			Column: sysmenu.FieldStatus,
		})
	}
	if value, ok := smuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldName,
		})
	}
	if value, ok := smuo.mutation.Icon(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldIcon,
		})
	}
	if value, ok := smuo.mutation.Router(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldRouter,
		})
	}
	if value, ok := smuo.mutation.IsShow(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: sysmenu.FieldIsShow,
		})
	}
	if value, ok := smuo.mutation.ParentID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentID,
		})
	}
	if smuo.mutation.ParentIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldParentID,
		})
	}
	if value, ok := smuo.mutation.ParentPath(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: sysmenu.FieldParentPath,
		})
	}
	if smuo.mutation.ParentPathCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: sysmenu.FieldParentPath,
		})
	}
	_node = &SysMenu{config: smuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, smuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{sysmenu.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
