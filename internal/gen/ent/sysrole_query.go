// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
)

// SysRoleQuery is the builder for querying SysRole entities.
type SysRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysRole
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysRoleQuery builder.
func (srq *SysRoleQuery) Where(ps ...predicate.SysRole) *SysRoleQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit adds a limit step to the query.
func (srq *SysRoleQuery) Limit(limit int) *SysRoleQuery {
	srq.limit = &limit
	return srq
}

// Offset adds an offset step to the query.
func (srq *SysRoleQuery) Offset(offset int) *SysRoleQuery {
	srq.offset = &offset
	return srq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (srq *SysRoleQuery) Unique(unique bool) *SysRoleQuery {
	srq.unique = &unique
	return srq
}

// Order adds an order step to the query.
func (srq *SysRoleQuery) Order(o ...OrderFunc) *SysRoleQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// First returns the first SysRole entity from the query.
// Returns a *NotFoundError when no SysRole was found.
func (srq *SysRoleQuery) First(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *SysRoleQuery) FirstX(ctx context.Context) *SysRole {
	node, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysRole ID from the query.
// Returns a *NotFoundError when no SysRole ID was found.
func (srq *SysRoleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = srq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (srq *SysRoleQuery) FirstIDX(ctx context.Context) string {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysRole entity is not found.
// Returns a *NotFoundError when no SysRole entities are found.
func (srq *SysRoleQuery) Only(ctx context.Context) (*SysRole, error) {
	nodes, err := srq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysrole.Label}
	default:
		return nil, &NotSingularError{sysrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyX(ctx context.Context) *SysRole {
	node, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysRole ID in the query.
// Returns a *NotSingularError when exactly one SysRole ID is not found.
// Returns a *NotFoundError when no entities are found.
func (srq *SysRoleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = srq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = &NotSingularError{sysrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *SysRoleQuery) OnlyIDX(ctx context.Context) string {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysRoles.
func (srq *SysRoleQuery) All(ctx context.Context) ([]*SysRole, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return srq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (srq *SysRoleQuery) AllX(ctx context.Context) []*SysRole {
	nodes, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysRole IDs.
func (srq *SysRoleQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := srq.Select(sysrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *SysRoleQuery) IDsX(ctx context.Context) []string {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *SysRoleQuery) Count(ctx context.Context) (int, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return srq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (srq *SysRoleQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *SysRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return srq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *SysRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *SysRoleQuery) Clone() *SysRoleQuery {
	if srq == nil {
		return nil
	}
	return &SysRoleQuery{
		config:     srq.config,
		limit:      srq.limit,
		offset:     srq.offset,
		order:      append([]OrderFunc{}, srq.order...),
		predicates: append([]predicate.SysRole{}, srq.predicates...),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status int32 `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysRole.Query().
//		GroupBy(sysrole.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (srq *SysRoleQuery) GroupBy(field string, fields ...string) *SysRoleGroupBy {
	group := &SysRoleGroupBy{config: srq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status int32 `json:"status,omitempty"`
//	}
//
//	client.SysRole.Query().
//		Select(sysrole.FieldStatus).
//		Scan(ctx, &v)
//
func (srq *SysRoleQuery) Select(field string, fields ...string) *SysRoleSelect {
	srq.fields = append([]string{field}, fields...)
	return &SysRoleSelect{SysRoleQuery: srq}
}

func (srq *SysRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range srq.fields {
		if !sysrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *SysRoleQuery) sqlAll(ctx context.Context) ([]*SysRole, error) {
	var (
		nodes = []*SysRole{}
		_spec = srq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysRole{config: srq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (srq *SysRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *SysRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := srq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (srq *SysRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysrole.Table,
			Columns: sysrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysrole.FieldID,
			},
		},
		From:   srq.sql,
		Unique: true,
	}
	if unique := srq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := srq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysrole.FieldID)
		for i := range fields {
			if fields[i] != sysrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *SysRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(sysrole.Table)
	selector := builder.Select(t1.Columns(sysrole.Columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(sysrole.Columns...)...)
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysRoleGroupBy is the group-by builder for SysRole entities.
type SysRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *SysRoleGroupBy) Aggregate(fns ...AggregateFunc) *SysRoleGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the group-by query and scans the result into the given value.
func (srgb *SysRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := srgb.path(ctx)
	if err != nil {
		return err
	}
	srgb.sql = query
	return srgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srgb *SysRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := srgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: SysRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srgb *SysRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := srgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srgb *SysRoleGroupBy) StringX(ctx context.Context) string {
	v, err := srgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: SysRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srgb *SysRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := srgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srgb *SysRoleGroupBy) IntX(ctx context.Context) int {
	v, err := srgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: SysRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srgb *SysRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := srgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srgb *SysRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := srgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: SysRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srgb *SysRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := srgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (srgb *SysRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srgb *SysRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := srgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srgb *SysRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range srgb.fields {
		if !sysrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := srgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := srgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srgb *SysRoleGroupBy) sqlQuery() *sql.Selector {
	selector := srgb.sql
	columns := make([]string, 0, len(srgb.fields)+len(srgb.fns))
	columns = append(columns, srgb.fields...)
	for _, fn := range srgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(srgb.fields...)
}

// SysRoleSelect is the builder for selecting fields of SysRole entities.
type SysRoleSelect struct {
	*SysRoleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (srs *SysRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := srs.prepareQuery(ctx); err != nil {
		return err
	}
	srs.sql = srs.SysRoleQuery.sqlQuery(ctx)
	return srs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srs *SysRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := srs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: SysRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srs *SysRoleSelect) StringsX(ctx context.Context) []string {
	v, err := srs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srs *SysRoleSelect) StringX(ctx context.Context) string {
	v, err := srs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: SysRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srs *SysRoleSelect) IntsX(ctx context.Context) []int {
	v, err := srs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srs *SysRoleSelect) IntX(ctx context.Context) int {
	v, err := srs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: SysRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srs *SysRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := srs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srs *SysRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := srs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: SysRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srs *SysRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := srs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (srs *SysRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysrole.Label}
	default:
		err = fmt.Errorf("ent: SysRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srs *SysRoleSelect) BoolX(ctx context.Context) bool {
	v, err := srs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srs *SysRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srs.sqlQuery().Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srs *SysRoleSelect) sqlQuery() sql.Querier {
	selector := srs.sql
	selector.Select(selector.Columns(srs.fields...)...)
	return selector
}
