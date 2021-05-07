// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wanhello/omgind/internal/gen/ent/predicate"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuactionresource"
)

// SysMenuActionQuery is the builder for querying SysMenuAction entities.
type SysMenuActionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.SysMenuAction
	// eager-loading edges.
	withResources *SysMenuActionResourceQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SysMenuActionQuery builder.
func (smaq *SysMenuActionQuery) Where(ps ...predicate.SysMenuAction) *SysMenuActionQuery {
	smaq.predicates = append(smaq.predicates, ps...)
	return smaq
}

// Limit adds a limit step to the query.
func (smaq *SysMenuActionQuery) Limit(limit int) *SysMenuActionQuery {
	smaq.limit = &limit
	return smaq
}

// Offset adds an offset step to the query.
func (smaq *SysMenuActionQuery) Offset(offset int) *SysMenuActionQuery {
	smaq.offset = &offset
	return smaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (smaq *SysMenuActionQuery) Unique(unique bool) *SysMenuActionQuery {
	smaq.unique = &unique
	return smaq
}

// Order adds an order step to the query.
func (smaq *SysMenuActionQuery) Order(o ...OrderFunc) *SysMenuActionQuery {
	smaq.order = append(smaq.order, o...)
	return smaq
}

// QueryResources chains the current query on the "resources" edge.
func (smaq *SysMenuActionQuery) QueryResources() *SysMenuActionResourceQuery {
	query := &SysMenuActionResourceQuery{config: smaq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := smaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := smaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(sysmenuaction.Table, sysmenuaction.FieldID, selector),
			sqlgraph.To(sysmenuactionresource.Table, sysmenuactionresource.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, sysmenuaction.ResourcesTable, sysmenuaction.ResourcesColumn),
		)
		fromU = sqlgraph.SetNeighbors(smaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SysMenuAction entity from the query.
// Returns a *NotFoundError when no SysMenuAction was found.
func (smaq *SysMenuActionQuery) First(ctx context.Context) (*SysMenuAction, error) {
	nodes, err := smaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{sysmenuaction.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (smaq *SysMenuActionQuery) FirstX(ctx context.Context) *SysMenuAction {
	node, err := smaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SysMenuAction ID from the query.
// Returns a *NotFoundError when no SysMenuAction ID was found.
func (smaq *SysMenuActionQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smaq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{sysmenuaction.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (smaq *SysMenuActionQuery) FirstIDX(ctx context.Context) string {
	id, err := smaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SysMenuAction entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one SysMenuAction entity is not found.
// Returns a *NotFoundError when no SysMenuAction entities are found.
func (smaq *SysMenuActionQuery) Only(ctx context.Context) (*SysMenuAction, error) {
	nodes, err := smaq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{sysmenuaction.Label}
	default:
		return nil, &NotSingularError{sysmenuaction.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (smaq *SysMenuActionQuery) OnlyX(ctx context.Context) *SysMenuAction {
	node, err := smaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SysMenuAction ID in the query.
// Returns a *NotSingularError when exactly one SysMenuAction ID is not found.
// Returns a *NotFoundError when no entities are found.
func (smaq *SysMenuActionQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = smaq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = &NotSingularError{sysmenuaction.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (smaq *SysMenuActionQuery) OnlyIDX(ctx context.Context) string {
	id, err := smaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SysMenuActions.
func (smaq *SysMenuActionQuery) All(ctx context.Context) ([]*SysMenuAction, error) {
	if err := smaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return smaq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (smaq *SysMenuActionQuery) AllX(ctx context.Context) []*SysMenuAction {
	nodes, err := smaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SysMenuAction IDs.
func (smaq *SysMenuActionQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := smaq.Select(sysmenuaction.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (smaq *SysMenuActionQuery) IDsX(ctx context.Context) []string {
	ids, err := smaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (smaq *SysMenuActionQuery) Count(ctx context.Context) (int, error) {
	if err := smaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return smaq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (smaq *SysMenuActionQuery) CountX(ctx context.Context) int {
	count, err := smaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (smaq *SysMenuActionQuery) Exist(ctx context.Context) (bool, error) {
	if err := smaq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return smaq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (smaq *SysMenuActionQuery) ExistX(ctx context.Context) bool {
	exist, err := smaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SysMenuActionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (smaq *SysMenuActionQuery) Clone() *SysMenuActionQuery {
	if smaq == nil {
		return nil
	}
	return &SysMenuActionQuery{
		config:        smaq.config,
		limit:         smaq.limit,
		offset:        smaq.offset,
		order:         append([]OrderFunc{}, smaq.order...),
		predicates:    append([]predicate.SysMenuAction{}, smaq.predicates...),
		withResources: smaq.withResources.Clone(),
		// clone intermediate query.
		sql:  smaq.sql.Clone(),
		path: smaq.path,
	}
}

// WithResources tells the query-builder to eager-load the nodes that are connected to
// the "resources" edge. The optional arguments are used to configure the query builder of the edge.
func (smaq *SysMenuActionQuery) WithResources(opts ...func(*SysMenuActionResourceQuery)) *SysMenuActionQuery {
	query := &SysMenuActionResourceQuery{config: smaq.config}
	for _, opt := range opts {
		opt(query)
	}
	smaq.withResources = query
	return smaq
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
//	client.SysMenuAction.Query().
//		GroupBy(sysmenuaction.FieldIsDel).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (smaq *SysMenuActionQuery) GroupBy(field string, fields ...string) *SysMenuActionGroupBy {
	group := &SysMenuActionGroupBy{config: smaq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := smaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return smaq.sqlQuery(ctx), nil
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
//	client.SysMenuAction.Query().
//		Select(sysmenuaction.FieldIsDel).
//		Scan(ctx, &v)
//
func (smaq *SysMenuActionQuery) Select(field string, fields ...string) *SysMenuActionSelect {
	smaq.fields = append([]string{field}, fields...)
	return &SysMenuActionSelect{SysMenuActionQuery: smaq}
}

func (smaq *SysMenuActionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range smaq.fields {
		if !sysmenuaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if smaq.path != nil {
		prev, err := smaq.path(ctx)
		if err != nil {
			return err
		}
		smaq.sql = prev
	}
	return nil
}

func (smaq *SysMenuActionQuery) sqlAll(ctx context.Context) ([]*SysMenuAction, error) {
	var (
		nodes       = []*SysMenuAction{}
		_spec       = smaq.querySpec()
		loadedTypes = [1]bool{
			smaq.withResources != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &SysMenuAction{config: smaq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, smaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := smaq.withResources; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[string]*SysMenuAction)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Resources = []*SysMenuActionResource{}
		}
		query.Where(predicate.SysMenuActionResource(func(s *sql.Selector) {
			s.Where(sql.InValues(sysmenuaction.ResourcesColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ActionID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "action_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Resources = append(node.Edges.Resources, n)
		}
	}

	return nodes, nil
}

func (smaq *SysMenuActionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := smaq.querySpec()
	return sqlgraph.CountNodes(ctx, smaq.driver, _spec)
}

func (smaq *SysMenuActionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := smaq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (smaq *SysMenuActionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   sysmenuaction.Table,
			Columns: sysmenuaction.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: sysmenuaction.FieldID,
			},
		},
		From:   smaq.sql,
		Unique: true,
	}
	if unique := smaq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := smaq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, sysmenuaction.FieldID)
		for i := range fields {
			if fields[i] != sysmenuaction.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := smaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := smaq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := smaq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := smaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (smaq *SysMenuActionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(smaq.driver.Dialect())
	t1 := builder.Table(sysmenuaction.Table)
	selector := builder.Select(t1.Columns(sysmenuaction.Columns...)...).From(t1)
	if smaq.sql != nil {
		selector = smaq.sql
		selector.Select(selector.Columns(sysmenuaction.Columns...)...)
	}
	for _, p := range smaq.predicates {
		p(selector)
	}
	for _, p := range smaq.order {
		p(selector)
	}
	if offset := smaq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := smaq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SysMenuActionGroupBy is the group-by builder for SysMenuAction entities.
type SysMenuActionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (smagb *SysMenuActionGroupBy) Aggregate(fns ...AggregateFunc) *SysMenuActionGroupBy {
	smagb.fns = append(smagb.fns, fns...)
	return smagb
}

// Scan applies the group-by query and scans the result into the given value.
func (smagb *SysMenuActionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := smagb.path(ctx)
	if err != nil {
		return err
	}
	smagb.sql = query
	return smagb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := smagb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(smagb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := smagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) StringsX(ctx context.Context) []string {
	v, err := smagb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smagb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) StringX(ctx context.Context) string {
	v, err := smagb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(smagb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := smagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) IntsX(ctx context.Context) []int {
	v, err := smagb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smagb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) IntX(ctx context.Context) int {
	v, err := smagb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(smagb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := smagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := smagb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smagb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := smagb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(smagb.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := smagb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := smagb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (smagb *SysMenuActionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smagb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smagb *SysMenuActionGroupBy) BoolX(ctx context.Context) bool {
	v, err := smagb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smagb *SysMenuActionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range smagb.fields {
		if !sysmenuaction.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := smagb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := smagb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smagb *SysMenuActionGroupBy) sqlQuery() *sql.Selector {
	selector := smagb.sql
	columns := make([]string, 0, len(smagb.fields)+len(smagb.fns))
	columns = append(columns, smagb.fields...)
	for _, fn := range smagb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(smagb.fields...)
}

// SysMenuActionSelect is the builder for selecting fields of SysMenuAction entities.
type SysMenuActionSelect struct {
	*SysMenuActionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (smas *SysMenuActionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := smas.prepareQuery(ctx); err != nil {
		return err
	}
	smas.sql = smas.SysMenuActionQuery.sqlQuery(ctx)
	return smas.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (smas *SysMenuActionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := smas.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(smas.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := smas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (smas *SysMenuActionSelect) StringsX(ctx context.Context) []string {
	v, err := smas.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = smas.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (smas *SysMenuActionSelect) StringX(ctx context.Context) string {
	v, err := smas.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(smas.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := smas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (smas *SysMenuActionSelect) IntsX(ctx context.Context) []int {
	v, err := smas.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = smas.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (smas *SysMenuActionSelect) IntX(ctx context.Context) int {
	v, err := smas.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(smas.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := smas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (smas *SysMenuActionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := smas.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = smas.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (smas *SysMenuActionSelect) Float64X(ctx context.Context) float64 {
	v, err := smas.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(smas.fields) > 1 {
		return nil, errors.New("ent: SysMenuActionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := smas.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (smas *SysMenuActionSelect) BoolsX(ctx context.Context) []bool {
	v, err := smas.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (smas *SysMenuActionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = smas.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{sysmenuaction.Label}
	default:
		err = fmt.Errorf("ent: SysMenuActionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (smas *SysMenuActionSelect) BoolX(ctx context.Context) bool {
	v, err := smas.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (smas *SysMenuActionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := smas.sqlQuery().Query()
	if err := smas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (smas *SysMenuActionSelect) sqlQuery() sql.Querier {
	selector := smas.sql
	selector.Select(selector.Columns(smas.fields...)...)
	return selector
}
