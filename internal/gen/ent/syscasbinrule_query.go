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
	"github.com/wanhello/omgind/internal/gen/ent/syscasbinrule"
)

// SysCasbinRuleQuery is the builder for querying SysCasbinRule entities.
type SysCasbinRuleQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysCasbinRule
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysCasbinRuleQuery builder.
func (scrq *SysCasbinRuleQuery) Where(ps ...predicate.SysCasbinRule) *SysCasbinRuleQuery {
	scrq.predicates = append(scrq.predicates, ps...)
	return scrq
}

// Limit adds a limit step to the query.
func (scrq *SysCasbinRuleQuery) Limit(limit int) *SysCasbinRuleQuery {
	scrq.limit = &limit
	return scrq
}

// Offset adds an offset step to the query.
func (scrq *SysCasbinRuleQuery) Offset(offset int) *SysCasbinRuleQuery {
	scrq.offset = &offset
	return scrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scrq *SysCasbinRuleQuery) Unique(unique bool) *SysCasbinRuleQuery {
	scrq.unique = &unique
	return scrq
}

// Order adds an order step to the query.
func (scrq *SysCasbinRuleQuery) Order(o ...OrderFunc) *SysCasbinRuleQuery {
	scrq.order = append(scrq.order, o...)
	return scrq
}

// First returns the first SysCasbinRule entity from the query.
// Returns a *NotFoundError when no SysCasbinRule was found.
func (scrq *SysCasbinRuleQuery) First(ctx context.Context) (*SysCasbinRule, error) {
	nodes, err := scrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{syscasbinrule.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) FirstX(ctx context.Context) *SysCasbinRule {
	node, err := scrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysCasbinRule ID from the query.
// Returns a *NotFoundError when no SysCasbinRule ID was found.
func (scrq *SysCasbinRuleQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = scrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{syscasbinrule.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) FirstIDX(ctx context.Context) string {
	id, err := scrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysCasbinRule entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysCasbinRule entity is not found.
// Returns a *NotFoundError when no SysCasbinRule entities are found.
func (scrq *SysCasbinRuleQuery) Only(ctx context.Context) (*SysCasbinRule, error) {
	nodes, err := scrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{syscasbinrule.Label}
	default:
		return nil, &NotSingularError{syscasbinrule.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) OnlyX(ctx context.Context) *SysCasbinRule {
	node, err := scrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysCasbinRule ID in the query.
// Returns a *NotSingularError when exactly one SysCasbinRule ID is not found.
// Returns a *NotFoundError when no entities are found.
func (scrq *SysCasbinRuleQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = scrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = &NotSingularError{syscasbinrule.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) OnlyIDX(ctx context.Context) string {
	id, err := scrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysCasbinRules.
func (scrq *SysCasbinRuleQuery) All(ctx context.Context) ([]*SysCasbinRule, error) {
	if err := scrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return scrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) AllX(ctx context.Context) []*SysCasbinRule {
	nodes, err := scrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysCasbinRule IDs.
func (scrq *SysCasbinRuleQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := scrq.Select(syscasbinrule.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) IDsX(ctx context.Context) []string {
	ids, err := scrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scrq *SysCasbinRuleQuery) Count(ctx context.Context) (int, error) {
	if err := scrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return scrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) CountX(ctx context.Context) int {
	count, err := scrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scrq *SysCasbinRuleQuery) Exist(ctx context.Context) (bool, error) {
	if err := scrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return scrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (scrq *SysCasbinRuleQuery) ExistX(ctx context.Context) bool {
	exist, err := scrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysCasbinRuleQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scrq *SysCasbinRuleQuery) Clone() *SysCasbinRuleQuery {
	if scrq == nil {
		return nil
	}
	return &SysCasbinRuleQuery{
		config:     scrq.config,
		limit:      scrq.limit,
		offset:     scrq.offset,
		order:      append([]OrderFunc{}, scrq.order...),
		predicates: append([]predicate.SysCasbinRule{}, scrq.predicates...),
		// clone intermediate query.
		sql:  scrq.sql.Clone(),
		path: scrq.path,
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
//	client.SysCasbinRule.Query().
//		GroupBy(syscasbinrule.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (scrq *SysCasbinRuleQuery) GroupBy(field string, fields ...string) *SysCasbinRuleGroupBy {
	group := &SysCasbinRuleGroupBy{config: scrq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := scrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return scrq.sqlQuery(ctx), nil
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
//	client.SysCasbinRule.Query().
//		Select(syscasbinrule.FieldIsDel).
//		Scan(ctx, &v)
//
func (scrq *SysCasbinRuleQuery) Select(field string, fields ...string) *SysCasbinRuleSelect {
	scrq.fields = append([]string{field}, fields...)
	return &SysCasbinRuleSelect{SysCasbinRuleQuery: scrq}
}

func (scrq *SysCasbinRuleQuery) prepareQuery(ctx context.Context) error {
	for _, f := range scrq.fields {
		if !syscasbinrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scrq.path != nil {
		prev, err := scrq.path(ctx)
		if err != nil {
			return err
		}
		scrq.sql = prev
	}
	return nil
}

func (scrq *SysCasbinRuleQuery) sqlAll(ctx context.Context) ([]*SysCasbinRule, error) {
	var (
		nodes = []*SysCasbinRule{}
		_spec = scrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysCasbinRule{config: scrq.config}
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
	if err := sqlgraph.QueryNodes(ctx, scrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (scrq *SysCasbinRuleQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scrq.querySpec()
	return sqlgraph.CountNodes(ctx, scrq.driver, _spec)
}

func (scrq *SysCasbinRuleQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := scrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (scrq *SysCasbinRuleQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   syscasbinrule.Table,
			Columns: syscasbinrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: syscasbinrule.FieldID,
			},
		},
		From:   scrq.sql,
		Unique: true,
	}
	if unique := scrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := scrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, syscasbinrule.FieldID)
		for i := range fields {
			if fields[i] != syscasbinrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scrq *SysCasbinRuleQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scrq.driver.Dialect())
	t1 := builder.Table(syscasbinrule.Table)
	selector := builder.Select(t1.Columns(syscasbinrule.Columns...)...).From(t1)
	if scrq.sql != nil {
		selector = scrq.sql
		selector.Select(selector.Columns(syscasbinrule.Columns...)...)
	}
	for _, p := range scrq.predicates {
		p(selector)
	}
	for _, p := range scrq.order {
		p(selector)
	}
	if offset := scrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysCasbinRuleGroupBy is the group-by builder for SysCasbinRule entities.
type SysCasbinRuleGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scrgb *SysCasbinRuleGroupBy) Aggregate(fns ...AggregateFunc) *SysCasbinRuleGroupBy {
	scrgb.fns = append(scrgb.fns, fns...)
	return scrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (scrgb *SysCasbinRuleGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := scrgb.path(ctx)
	if err != nil {
		return err
	}
	scrgb.sql = query
	return scrgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := scrgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(scrgb.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := scrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) StringsX(ctx context.Context) []string {
	v, err := scrgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scrgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) StringX(ctx context.Context) string {
	v, err := scrgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(scrgb.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := scrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) IntsX(ctx context.Context) []int {
	v, err := scrgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scrgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) IntX(ctx context.Context) int {
	v, err := scrgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(scrgb.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := scrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := scrgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scrgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) Float64X(ctx context.Context) float64 {
	v, err := scrgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(scrgb.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := scrgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := scrgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (scrgb *SysCasbinRuleGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scrgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scrgb *SysCasbinRuleGroupBy) BoolX(ctx context.Context) bool {
	v, err := scrgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scrgb *SysCasbinRuleGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range scrgb.fields {
		if !syscasbinrule.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := scrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scrgb *SysCasbinRuleGroupBy) sqlQuery() *sql.Selector {
	selector := scrgb.sql
	columns := make([]string, 0, len(scrgb.fields)+len(scrgb.fns))
	columns = append(columns, scrgb.fields...)
	for _, fn := range scrgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(scrgb.fields...)
}

// SysCasbinRuleSelect is the builder for selecting fields of SysCasbinRule entities.
type SysCasbinRuleSelect struct {
	*SysCasbinRuleQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (scrs *SysCasbinRuleSelect) Scan(ctx context.Context, v interface{}) error {
	if err := scrs.prepareQuery(ctx); err != nil {
		return err
	}
	scrs.sql = scrs.SysCasbinRuleQuery.sqlQuery(ctx)
	return scrs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) ScanX(ctx context.Context, v interface{}) {
	if err := scrs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Strings(ctx context.Context) ([]string, error) {
	if len(scrs.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := scrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) StringsX(ctx context.Context) []string {
	v, err := scrs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = scrs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) StringX(ctx context.Context) string {
	v, err := scrs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Ints(ctx context.Context) ([]int, error) {
	if len(scrs.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := scrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) IntsX(ctx context.Context) []int {
	v, err := scrs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = scrs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) IntX(ctx context.Context) int {
	v, err := scrs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(scrs.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := scrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) Float64sX(ctx context.Context) []float64 {
	v, err := scrs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = scrs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) Float64X(ctx context.Context) float64 {
	v, err := scrs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(scrs.fields) > 1 {
		return nil, errors.New("ent: SysCasbinRuleSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := scrs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) BoolsX(ctx context.Context) []bool {
	v, err := scrs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (scrs *SysCasbinRuleSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = scrs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{syscasbinrule.Label}
	default:
		err = fmt.Errorf("ent: SysCasbinRuleSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (scrs *SysCasbinRuleSelect) BoolX(ctx context.Context) bool {
	v, err := scrs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (scrs *SysCasbinRuleSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := scrs.sqlQuery().Query()
	if err := scrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scrs *SysCasbinRuleSelect) sqlQuery() sql.Querier {
	selector := scrs.sql
	selector.Select(selector.Columns(scrs.fields...)...)
	return selector
}
