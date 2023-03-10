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
	"github.com/wanhello/omgind/internal/gen/ent/sysdictitem"
)

// SysDictItemQuery is the builder for querying SysDictItem entities.
type SysDictItemQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysDictItem
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysDictItemQuery builder.
func (sdiq *SysDictItemQuery) Where(ps ...predicate.SysDictItem) *SysDictItemQuery {
	sdiq.predicates = append(sdiq.predicates, ps...)
	return sdiq
}

// Limit adds a limit step to the query.
func (sdiq *SysDictItemQuery) Limit(limit int) *SysDictItemQuery {
	sdiq.limit = &limit
	return sdiq
}

// Offset adds an offset step to the query.
func (sdiq *SysDictItemQuery) Offset(offset int) *SysDictItemQuery {
	sdiq.offset = &offset
	return sdiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sdiq *SysDictItemQuery) Unique(unique bool) *SysDictItemQuery {
	sdiq.unique = &unique
	return sdiq
}

// Order adds an order step to the query.
func (sdiq *SysDictItemQuery) Order(o ...OrderFunc) *SysDictItemQuery {
	sdiq.order = append(sdiq.order, o...)
	return sdiq
}

// First returns the first SysDictItem entity from the query.
// Returns a *NotFoundError when no SysDictItem was found.
func (sdiq *SysDictItemQuery) First(ctx context.Context) (*SysDictItem, error) {
	nodes, err := sdiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysdictitem.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sdiq *SysDictItemQuery) FirstX(ctx context.Context) *SysDictItem {
	node, err := sdiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysDictItem ID from the query.
// Returns a *NotFoundError when no SysDictItem ID was found.
func (sdiq *SysDictItemQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysdictitem.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sdiq *SysDictItemQuery) FirstIDX(ctx context.Context) string {
	id, err := sdiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysDictItem entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysDictItem entity is not found.
// Returns a *NotFoundError when no SysDictItem entities are found.
func (sdiq *SysDictItemQuery) Only(ctx context.Context) (*SysDictItem, error) {
	nodes, err := sdiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysdictitem.Label}
	default:
		return nil, &NotSingularError{sysdictitem.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sdiq *SysDictItemQuery) OnlyX(ctx context.Context) *SysDictItem {
	node, err := sdiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysDictItem ID in the query.
// Returns a *NotSingularError when exactly one SysDictItem ID is not found.
// Returns a *NotFoundError when no entities are found.
func (sdiq *SysDictItemQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = sdiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = &NotSingularError{sysdictitem.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sdiq *SysDictItemQuery) OnlyIDX(ctx context.Context) string {
	id, err := sdiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysDictItems.
func (sdiq *SysDictItemQuery) All(ctx context.Context) ([]*SysDictItem, error) {
	if err := sdiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sdiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sdiq *SysDictItemQuery) AllX(ctx context.Context) []*SysDictItem {
	nodes, err := sdiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysDictItem IDs.
func (sdiq *SysDictItemQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := sdiq.Select(sysdictitem.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sdiq *SysDictItemQuery) IDsX(ctx context.Context) []string {
	ids, err := sdiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sdiq *SysDictItemQuery) Count(ctx context.Context) (int, error) {
	if err := sdiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sdiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sdiq *SysDictItemQuery) CountX(ctx context.Context) int {
	count, err := sdiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sdiq *SysDictItemQuery) Exist(ctx context.Context) (bool, error) {
	if err := sdiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sdiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sdiq *SysDictItemQuery) ExistX(ctx context.Context) bool {
	exist, err := sdiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysDictItemQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sdiq *SysDictItemQuery) Clone() *SysDictItemQuery {
	if sdiq == nil {
		return nil
	}
	return &SysDictItemQuery{
		config:     sdiq.config,
		limit:      sdiq.limit,
		offset:     sdiq.offset,
		order:      append([]OrderFunc{}, sdiq.order...),
		predicates: append([]predicate.SysDictItem{}, sdiq.predicates...),
		// clone intermediate query.
		sql:  sdiq.sql.Clone(),
		path: sdiq.path,
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
//	client.SysDictItem.Query().
//		GroupBy(sysdictitem.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sdiq *SysDictItemQuery) GroupBy(field string, fields ...string) *SysDictItemGroupBy {
	group := &SysDictItemGroupBy{config: sdiq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sdiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sdiq.sqlQuery(ctx), nil
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
//	client.SysDictItem.Query().
//		Select(sysdictitem.FieldIsDel).
//		Scan(ctx, &v)
//
func (sdiq *SysDictItemQuery) Select(field string, fields ...string) *SysDictItemSelect {
	sdiq.fields = append([]string{field}, fields...)
	return &SysDictItemSelect{SysDictItemQuery: sdiq}
}

func (sdiq *SysDictItemQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sdiq.fields {
		if !sysdictitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sdiq.path != nil {
		prev, err := sdiq.path(ctx)
		if err != nil {
			return err
		}
		sdiq.sql = prev
	}
	return nil
}

func (sdiq *SysDictItemQuery) sqlAll(ctx context.Context) ([]*SysDictItem, error) {
	var (
		nodes = []*SysDictItem{}
		_spec = sdiq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysDictItem{config: sdiq.config}
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
	if err := sqlgraph.QueryNodes(ctx, sdiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (sdiq *SysDictItemQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sdiq.querySpec()
	return sqlgraph.CountNodes(ctx, sdiq.driver, _spec)
}

func (sdiq *SysDictItemQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sdiq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (sdiq *SysDictItemQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysdictitem.Table,
			Columns: sysdictitem.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysdictitem.FieldID,
			},
		},
		From:   sdiq.sql,
		Unique: true,
	}
	if unique := sdiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sdiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysdictitem.FieldID)
		for i := range fields {
			if fields[i] != sysdictitem.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sdiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sdiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sdiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sdiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sdiq *SysDictItemQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sdiq.driver.Dialect())
	t1 := builder.Table(sysdictitem.Table)
	selector := builder.Select(t1.Columns(sysdictitem.Columns...)...).From(t1)
	if sdiq.sql != nil {
		selector = sdiq.sql
		selector.Select(selector.Columns(sysdictitem.Columns...)...)
	}
	for _, p := range sdiq.predicates {
		p(selector)
	}
	for _, p := range sdiq.order {
		p(selector)
	}
	if offset := sdiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sdiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysDictItemGroupBy is the group-by builder for SysDictItem entities.
type SysDictItemGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sdigb *SysDictItemGroupBy) Aggregate(fns ...AggregateFunc) *SysDictItemGroupBy {
	sdigb.fns = append(sdigb.fns, fns...)
	return sdigb
}

// Scan applies the group-by query and scans the result into the given value.
func (sdigb *SysDictItemGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := sdigb.path(ctx)
	if err != nil {
		return err
	}
	sdigb.sql = query
	return sdigb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sdigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sdigb.fields) > 1 {
		return nil, errors.New("ent: SysDictItemGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sdigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) StringsX(ctx context.Context) []string {
	v, err := sdigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sdigb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) StringX(ctx context.Context) string {
	v, err := sdigb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sdigb.fields) > 1 {
		return nil, errors.New("ent: SysDictItemGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sdigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) IntsX(ctx context.Context) []int {
	v, err := sdigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sdigb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) IntX(ctx context.Context) int {
	v, err := sdigb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sdigb.fields) > 1 {
		return nil, errors.New("ent: SysDictItemGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sdigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sdigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sdigb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) Float64X(ctx context.Context) float64 {
	v, err := sdigb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sdigb.fields) > 1 {
		return nil, errors.New("ent: SysDictItemGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sdigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sdigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (sdigb *SysDictItemGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sdigb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sdigb *SysDictItemGroupBy) BoolX(ctx context.Context) bool {
	v, err := sdigb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sdigb *SysDictItemGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range sdigb.fields {
		if !sysdictitem.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := sdigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sdigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sdigb *SysDictItemGroupBy) sqlQuery() *sql.Selector {
	selector := sdigb.sql
	columns := make([]string, 0, len(sdigb.fields)+len(sdigb.fns))
	columns = append(columns, sdigb.fields...)
	for _, fn := range sdigb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sdigb.fields...)
}

// SysDictItemSelect is the builder for selecting fields of SysDictItem entities.
type SysDictItemSelect struct {
	*SysDictItemQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sdis *SysDictItemSelect) Scan(ctx context.Context, v interface{}) error {
	if err := sdis.prepareQuery(ctx); err != nil {
		return err
	}
	sdis.sql = sdis.SysDictItemQuery.sqlQuery(ctx)
	return sdis.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sdis *SysDictItemSelect) ScanX(ctx context.Context, v interface{}) {
	if err := sdis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Strings(ctx context.Context) ([]string, error) {
	if len(sdis.fields) > 1 {
		return nil, errors.New("ent: SysDictItemSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := sdis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sdis *SysDictItemSelect) StringsX(ctx context.Context) []string {
	v, err := sdis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = sdis.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (sdis *SysDictItemSelect) StringX(ctx context.Context) string {
	v, err := sdis.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Ints(ctx context.Context) ([]int, error) {
	if len(sdis.fields) > 1 {
		return nil, errors.New("ent: SysDictItemSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := sdis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sdis *SysDictItemSelect) IntsX(ctx context.Context) []int {
	v, err := sdis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = sdis.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (sdis *SysDictItemSelect) IntX(ctx context.Context) int {
	v, err := sdis.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(sdis.fields) > 1 {
		return nil, errors.New("ent: SysDictItemSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := sdis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sdis *SysDictItemSelect) Float64sX(ctx context.Context) []float64 {
	v, err := sdis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = sdis.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (sdis *SysDictItemSelect) Float64X(ctx context.Context) float64 {
	v, err := sdis.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(sdis.fields) > 1 {
		return nil, errors.New("ent: SysDictItemSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := sdis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sdis *SysDictItemSelect) BoolsX(ctx context.Context) []bool {
	v, err := sdis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (sdis *SysDictItemSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = sdis.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysdictitem.Label}
	default:
		err = fmt.Errorf("ent: SysDictItemSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (sdis *SysDictItemSelect) BoolX(ctx context.Context) bool {
	v, err := sdis.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sdis *SysDictItemSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sdis.sqlQuery().Query()
	if err := sdis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sdis *SysDictItemSelect) sqlQuery() sql.Querier {
	selector := sdis.sql
	selector.Select(selector.Columns(sdis.fields...)...)
	return selector
}
