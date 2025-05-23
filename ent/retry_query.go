// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"
	"notification/ent/message"
	"notification/ent/predicate"
	"notification/ent/retry"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// RetryQuery is the builder for querying Retry entities.
type RetryQuery struct {
	config
	ctx         *QueryContext
	order       []retry.OrderOption
	inters      []Interceptor
	predicates  []predicate.Retry
	withMessage *MessageQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RetryQuery builder.
func (rq *RetryQuery) Where(ps ...predicate.Retry) *RetryQuery {
	rq.predicates = append(rq.predicates, ps...)
	return rq
}

// Limit the number of records to be returned by this query.
func (rq *RetryQuery) Limit(limit int) *RetryQuery {
	rq.ctx.Limit = &limit
	return rq
}

// Offset to start from.
func (rq *RetryQuery) Offset(offset int) *RetryQuery {
	rq.ctx.Offset = &offset
	return rq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rq *RetryQuery) Unique(unique bool) *RetryQuery {
	rq.ctx.Unique = &unique
	return rq
}

// Order specifies how the records should be ordered.
func (rq *RetryQuery) Order(o ...retry.OrderOption) *RetryQuery {
	rq.order = append(rq.order, o...)
	return rq
}

// QueryMessage chains the current query on the "message" edge.
func (rq *RetryQuery) QueryMessage() *MessageQuery {
	query := (&MessageClient{config: rq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(retry.Table, retry.FieldID, selector),
			sqlgraph.To(message.Table, message.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, retry.MessageTable, retry.MessageColumn),
		)
		fromU = sqlgraph.SetNeighbors(rq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Retry entity from the query.
// Returns a *NotFoundError when no Retry was found.
func (rq *RetryQuery) First(ctx context.Context) (*Retry, error) {
	nodes, err := rq.Limit(1).All(setContextOp(ctx, rq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{retry.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rq *RetryQuery) FirstX(ctx context.Context) *Retry {
	node, err := rq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Retry ID from the query.
// Returns a *NotFoundError when no Retry ID was found.
func (rq *RetryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(1).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{retry.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rq *RetryQuery) FirstIDX(ctx context.Context) int {
	id, err := rq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Retry entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Retry entity is found.
// Returns a *NotFoundError when no Retry entities are found.
func (rq *RetryQuery) Only(ctx context.Context) (*Retry, error) {
	nodes, err := rq.Limit(2).All(setContextOp(ctx, rq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{retry.Label}
	default:
		return nil, &NotSingularError{retry.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rq *RetryQuery) OnlyX(ctx context.Context) *Retry {
	node, err := rq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Retry ID in the query.
// Returns a *NotSingularError when more than one Retry ID is found.
// Returns a *NotFoundError when no entities are found.
func (rq *RetryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rq.Limit(2).IDs(setContextOp(ctx, rq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{retry.Label}
	default:
		err = &NotSingularError{retry.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rq *RetryQuery) OnlyIDX(ctx context.Context) int {
	id, err := rq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Retries.
func (rq *RetryQuery) All(ctx context.Context) ([]*Retry, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryAll)
	if err := rq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Retry, *RetryQuery]()
	return withInterceptors[[]*Retry](ctx, rq, qr, rq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rq *RetryQuery) AllX(ctx context.Context) []*Retry {
	nodes, err := rq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Retry IDs.
func (rq *RetryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rq.ctx.Unique == nil && rq.path != nil {
		rq.Unique(true)
	}
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryIDs)
	if err = rq.Select(retry.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rq *RetryQuery) IDsX(ctx context.Context) []int {
	ids, err := rq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rq *RetryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryCount)
	if err := rq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rq, querierCount[*RetryQuery](), rq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rq *RetryQuery) CountX(ctx context.Context) int {
	count, err := rq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rq *RetryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rq.ctx, ent.OpQueryExist)
	switch _, err := rq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rq *RetryQuery) ExistX(ctx context.Context) bool {
	exist, err := rq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RetryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rq *RetryQuery) Clone() *RetryQuery {
	if rq == nil {
		return nil
	}
	return &RetryQuery{
		config:      rq.config,
		ctx:         rq.ctx.Clone(),
		order:       append([]retry.OrderOption{}, rq.order...),
		inters:      append([]Interceptor{}, rq.inters...),
		predicates:  append([]predicate.Retry{}, rq.predicates...),
		withMessage: rq.withMessage.Clone(),
		// clone intermediate query.
		sql:  rq.sql.Clone(),
		path: rq.path,
	}
}

// WithMessage tells the query-builder to eager-load the nodes that are connected to
// the "message" edge. The optional arguments are used to configure the query builder of the edge.
func (rq *RetryQuery) WithMessage(opts ...func(*MessageQuery)) *RetryQuery {
	query := (&MessageClient{config: rq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rq.withMessage = query
	return rq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		MessageUUID uuid.UUID `json:"message_uuid,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Retry.Query().
//		GroupBy(retry.FieldMessageUUID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rq *RetryQuery) GroupBy(field string, fields ...string) *RetryGroupBy {
	rq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RetryGroupBy{build: rq}
	grbuild.flds = &rq.ctx.Fields
	grbuild.label = retry.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		MessageUUID uuid.UUID `json:"message_uuid,omitempty"`
//	}
//
//	client.Retry.Query().
//		Select(retry.FieldMessageUUID).
//		Scan(ctx, &v)
func (rq *RetryQuery) Select(fields ...string) *RetrySelect {
	rq.ctx.Fields = append(rq.ctx.Fields, fields...)
	sbuild := &RetrySelect{RetryQuery: rq}
	sbuild.label = retry.Label
	sbuild.flds, sbuild.scan = &rq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RetrySelect configured with the given aggregations.
func (rq *RetryQuery) Aggregate(fns ...AggregateFunc) *RetrySelect {
	return rq.Select().Aggregate(fns...)
}

func (rq *RetryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rq); err != nil {
				return err
			}
		}
	}
	for _, f := range rq.ctx.Fields {
		if !retry.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rq.path != nil {
		prev, err := rq.path(ctx)
		if err != nil {
			return err
		}
		rq.sql = prev
	}
	return nil
}

func (rq *RetryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Retry, error) {
	var (
		nodes       = []*Retry{}
		_spec       = rq.querySpec()
		loadedTypes = [1]bool{
			rq.withMessage != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Retry).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Retry{config: rq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rq.withMessage; query != nil {
		if err := rq.loadMessage(ctx, query, nodes, nil,
			func(n *Retry, e *Message) { n.Edges.Message = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rq *RetryQuery) loadMessage(ctx context.Context, query *MessageQuery, nodes []*Retry, init func(*Retry), assign func(*Retry, *Message)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*Retry)
	for i := range nodes {
		fk := nodes[i].MessageUUID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(message.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "message_uuid" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rq *RetryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rq.querySpec()
	_spec.Node.Columns = rq.ctx.Fields
	if len(rq.ctx.Fields) > 0 {
		_spec.Unique = rq.ctx.Unique != nil && *rq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rq.driver, _spec)
}

func (rq *RetryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(retry.Table, retry.Columns, sqlgraph.NewFieldSpec(retry.FieldID, field.TypeInt))
	_spec.From = rq.sql
	if unique := rq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rq.path != nil {
		_spec.Unique = true
	}
	if fields := rq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, retry.FieldID)
		for i := range fields {
			if fields[i] != retry.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if rq.withMessage != nil {
			_spec.Node.AddColumnOnce(retry.FieldMessageUUID)
		}
	}
	if ps := rq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rq *RetryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rq.driver.Dialect())
	t1 := builder.Table(retry.Table)
	columns := rq.ctx.Fields
	if len(columns) == 0 {
		columns = retry.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rq.sql != nil {
		selector = rq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rq.ctx.Unique != nil && *rq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rq.predicates {
		p(selector)
	}
	for _, p := range rq.order {
		p(selector)
	}
	if offset := rq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RetryGroupBy is the group-by builder for Retry entities.
type RetryGroupBy struct {
	selector
	build *RetryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rgb *RetryGroupBy) Aggregate(fns ...AggregateFunc) *RetryGroupBy {
	rgb.fns = append(rgb.fns, fns...)
	return rgb
}

// Scan applies the selector query and scans the result into the given value.
func (rgb *RetryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgb.build.ctx, ent.OpQueryGroupBy)
	if err := rgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RetryQuery, *RetryGroupBy](ctx, rgb.build, rgb, rgb.build.inters, v)
}

func (rgb *RetryGroupBy) sqlScan(ctx context.Context, root *RetryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rgb.fns))
	for _, fn := range rgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rgb.flds)+len(rgb.fns))
		for _, f := range *rgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RetrySelect is the builder for selecting fields of Retry entities.
type RetrySelect struct {
	*RetryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rs *RetrySelect) Aggregate(fns ...AggregateFunc) *RetrySelect {
	rs.fns = append(rs.fns, fns...)
	return rs
}

// Scan applies the selector query and scans the result into the given value.
func (rs *RetrySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rs.ctx, ent.OpQuerySelect)
	if err := rs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RetryQuery, *RetrySelect](ctx, rs.RetryQuery, rs, rs.inters, v)
}

func (rs *RetrySelect) sqlScan(ctx context.Context, root *RetryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rs.fns))
	for _, fn := range rs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
