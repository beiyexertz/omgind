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
	"github.com/wanhello/omgind/internal/gen/ent/sysuserrole"
)

// SysUserRoleQuery is the builder for querying SysUserRole entities.
type SysUserRoleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysUserRole
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysUserRoleQuery builder.
func (surq *SysUserRoleQuery) Where(ps ...predicate.SysUserRole) *SysUserRoleQuery {
	surq.predicates = append(surq.predicates, ps...)
	return surq
}

// Limit adds a limit step to the query.
func (surq *SysUserRoleQuery) Limit(limit int) *SysUserRoleQuery {
	surq.limit = &limit
	return surq
}

// Offset adds an offset step to the query.
func (surq *SysUserRoleQuery) Offset(offset int) *SysUserRoleQuery {
	surq.offset = &offset
	return surq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (surq *SysUserRoleQuery) Unique(unique bool) *SysUserRoleQuery {
	surq.unique = &unique
	return surq
}

// Order adds an order step to the query.
func (surq *SysUserRoleQuery) Order(o ...OrderFunc) *SysUserRoleQuery {
	surq.order = append(surq.order, o...)
	return surq
}

// First returns the first SysUserRole entity from the query.
// Returns a *NotFoundError when no SysUserRole was found.
func (surq *SysUserRoleQuery) First(ctx context.Context) (*SysUserRole, error) {
	nodes, err := surq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysuserrole.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (surq *SysUserRoleQuery) FirstX(ctx context.Context) *SysUserRole {
	node, err := surq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysUserRole ID from the query.
// Returns a *NotFoundError when no SysUserRole ID was found.
func (surq *SysUserRoleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = surq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysuserrole.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (surq *SysUserRoleQuery) FirstIDX(ctx context.Context) string {
	id, err := surq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysUserRole entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysUserRole entity is not found.
// Returns a *NotFoundError when no SysUserRole entities are found.
func (surq *SysUserRoleQuery) Only(ctx context.Context) (*SysUserRole, error) {
	nodes, err := surq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysuserrole.Label}
	default:
		return nil, &NotSingularError{sysuserrole.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (surq *SysUserRoleQuery) OnlyX(ctx context.Context) *SysUserRole {
	node, err := surq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysUserRole ID in the query.
// Returns a *NotSingularError when exactly one SysUserRole ID is not found.
// Returns a *NotFoundError when no entities are found.
func (surq *SysUserRoleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = surq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = &NotSingularError{sysuserrole.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (surq *SysUserRoleQuery) OnlyIDX(ctx context.Context) string {
	id, err := surq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysUserRoles.
func (surq *SysUserRoleQuery) All(ctx context.Context) ([]*SysUserRole, error) {
	if err := surq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return surq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (surq *SysUserRoleQuery) AllX(ctx context.Context) []*SysUserRole {
	nodes, err := surq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysUserRole IDs.
func (surq *SysUserRoleQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := surq.Select(sysuserrole.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (surq *SysUserRoleQuery) IDsX(ctx context.Context) []string {
	ids, err := surq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (surq *SysUserRoleQuery) Count(ctx context.Context) (int, error) {
	if err := surq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return surq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (surq *SysUserRoleQuery) CountX(ctx context.Context) int {
	count, err := surq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (surq *SysUserRoleQuery) Exist(ctx context.Context) (bool, error) {
	if err := surq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return surq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (surq *SysUserRoleQuery) ExistX(ctx context.Context) bool {
	exist, err := surq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysUserRoleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (surq *SysUserRoleQuery) Clone() *SysUserRoleQuery {
	if surq == nil {
		return nil
	}
	return &SysUserRoleQuery{
		config:     surq.config,
		limit:      surq.limit,
		offset:     surq.offset,
		order:      append([]OrderFunc{}, surq.order...),
		predicates: append([]predicate.SysUserRole{}, surq.predicates...),
		// clone intermediate query.
		sql:  surq.sql.Clone(),
		path: surq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		IsDel bool `json:"is_del,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SysUserRole.Query().
//		GroupBy(sysuserrole.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (surq *SysUserRoleQuery) GroupBy(field string, fields ...string) *SysUserRoleGroupBy {
	group := &SysUserRoleGroupBy{config: surq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := surq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return surq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		IsDel bool `json:"is_del,omitempty"`
//	}
//
//	client.SysUserRole.Query().
//		Select(sysuserrole.FieldIsDel).
//		Scan(ctx, &v)
//
func (surq *SysUserRoleQuery) Select(field string, fields ...string) *SysUserRoleSelect {
	surq.fields = append([]string{field}, fields...)
	return &SysUserRoleSelect{SysUserRoleQuery: surq}
}

func (surq *SysUserRoleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range surq.fields {
		if !sysuserrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if surq.path != nil {
		prev, err := surq.path(ctx)
		if err != nil {
			return err
		}
		surq.sql = prev
	}
	return nil
}

func (surq *SysUserRoleQuery) sqlAll(ctx context.Context) ([]*SysUserRole, error) {
	var (
		nodes = []*SysUserRole{}
		_spec = surq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysUserRole{config: surq.config}
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
	if err := sqlgraph.QueryNodes(ctx, surq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (surq *SysUserRoleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := surq.querySpec()
	return sqlgraph.CountNodes(ctx, surq.driver, _spec)
}

func (surq *SysUserRoleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := surq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (surq *SysUserRoleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysuserrole.Table,
			Columns: sysuserrole.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysuserrole.FieldID,
			},
		},
		From:   surq.sql,
		Unique: true,
	}
	if unique := surq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := surq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysuserrole.FieldID)
		for i := range fields {
			if fields[i] != sysuserrole.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := surq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := surq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := surq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := surq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (surq *SysUserRoleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(surq.driver.Dialect())
	t1 := builder.Table(sysuserrole.Table)
	selector := builder.Select(t1.Columns(sysuserrole.Columns...)...).From(t1)
	if surq.sql != nil {
		selector = surq.sql
		selector.Select(selector.Columns(sysuserrole.Columns...)...)
	}
	for _, p := range surq.predicates {
		p(selector)
	}
	for _, p := range surq.order {
		p(selector)
	}
	if offset := surq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := surq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysUserRoleGroupBy is the group-by builder for SysUserRole entities.
type SysUserRoleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (surgb *SysUserRoleGroupBy) Aggregate(fns ...AggregateFunc) *SysUserRoleGroupBy {
	surgb.fns = append(surgb.fns, fns...)
	return surgb
}

// Scan applies the group-by query and scans the result into the given value.
func (surgb *SysUserRoleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := surgb.path(ctx)
	if err != nil {
		return err
	}
	surgb.sql = query
	return surgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := surgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(surgb.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := surgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) StringsX(ctx context.Context) []string {
	v, err := surgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = surgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) StringX(ctx context.Context) string {
	v, err := surgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(surgb.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := surgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) IntsX(ctx context.Context) []int {
	v, err := surgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = surgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) IntX(ctx context.Context) int {
	v, err := surgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(surgb.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := surgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := surgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = surgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := surgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(surgb.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := surgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := surgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (surgb *SysUserRoleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = surgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (surgb *SysUserRoleGroupBy) BoolX(ctx context.Context) bool {
	v, err := surgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (surgb *SysUserRoleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range surgb.fields {
		if !sysuserrole.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := surgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := surgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (surgb *SysUserRoleGroupBy) sqlQuery() *sql.Selector {
	selector := surgb.sql
	columns := make([]string, 0, len(surgb.fields)+len(surgb.fns))
	columns = append(columns, surgb.fields...)
	for _, fn := range surgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(surgb.fields...)
}

// SysUserRoleSelect is the builder for selecting fields of SysUserRole entities.
type SysUserRoleSelect struct {
	*SysUserRoleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (surs *SysUserRoleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := surs.prepareQuery(ctx); err != nil {
		return err
	}
	surs.sql = surs.SysUserRoleQuery.sqlQuery(ctx)
	return surs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (surs *SysUserRoleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := surs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(surs.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := surs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (surs *SysUserRoleSelect) StringsX(ctx context.Context) []string {
	v, err := surs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = surs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (surs *SysUserRoleSelect) StringX(ctx context.Context) string {
	v, err := surs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(surs.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := surs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (surs *SysUserRoleSelect) IntsX(ctx context.Context) []int {
	v, err := surs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = surs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (surs *SysUserRoleSelect) IntX(ctx context.Context) int {
	v, err := surs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(surs.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := surs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (surs *SysUserRoleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := surs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = surs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (surs *SysUserRoleSelect) Float64X(ctx context.Context) float64 {
	v, err := surs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(surs.fields) > 1 {
		return nil, errors.New("ent: SysUserRoleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := surs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (surs *SysUserRoleSelect) BoolsX(ctx context.Context) []bool {
	v, err := surs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (surs *SysUserRoleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = surs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysuserrole.Label}
	default:
		err = fmt.Errorf("ent: SysUserRoleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (surs *SysUserRoleSelect) BoolX(ctx context.Context) bool {
	v, err := surs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (surs *SysUserRoleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := surs.sqlQuery().Query()
	if err := surs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (surs *SysUserRoleSelect) sqlQuery() sql.Querier {
	selector := surs.sql
	selector.Select(selector.Columns(surs.fields...)...)
	return selector
}
