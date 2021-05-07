// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/wanhello/omgind/internal/gen/ent/migrate"

	"github.com/wanhello/omgind/internal/gen/ent/sysmenu"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuaction"
	"github.com/wanhello/omgind/internal/gen/ent/sysmenuactionresource"
	"github.com/wanhello/omgind/internal/gen/ent/sysrole"
	"github.com/wanhello/omgind/internal/gen/ent/sysrolemenu"
	"github.com/wanhello/omgind/internal/gen/ent/sysuser"
	"github.com/wanhello/omgind/internal/gen/ent/sysuserrole"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// SysMenu is the client for interacting with the SysMenu builders.
	SysMenu *SysMenuClient
	// SysMenuAction is the client for interacting with the SysMenuAction builders.
	SysMenuAction *SysMenuActionClient
	// SysMenuActionResource is the client for interacting with the SysMenuActionResource builders.
	SysMenuActionResource *SysMenuActionResourceClient
	// SysRole is the client for interacting with the SysRole builders.
	SysRole *SysRoleClient
	// SysRoleMenu is the client for interacting with the SysRoleMenu builders.
	SysRoleMenu *SysRoleMenuClient
	// SysUser is the client for interacting with the SysUser builders.
	SysUser *SysUserClient
	// SysUserRole is the client for interacting with the SysUserRole builders.
	SysUserRole *SysUserRoleClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.SysMenu = NewSysMenuClient(c.config)
	c.SysMenuAction = NewSysMenuActionClient(c.config)
	c.SysMenuActionResource = NewSysMenuActionResourceClient(c.config)
	c.SysRole = NewSysRoleClient(c.config)
	c.SysRoleMenu = NewSysRoleMenuClient(c.config)
	c.SysUser = NewSysUserClient(c.config)
	c.SysUserRole = NewSysUserRoleClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:                   ctx,
		config:                cfg,
		SysMenu:               NewSysMenuClient(cfg),
		SysMenuAction:         NewSysMenuActionClient(cfg),
		SysMenuActionResource: NewSysMenuActionResourceClient(cfg),
		SysRole:               NewSysRoleClient(cfg),
		SysRoleMenu:           NewSysRoleMenuClient(cfg),
		SysUser:               NewSysUserClient(cfg),
		SysUserRole:           NewSysUserRoleClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:                cfg,
		SysMenu:               NewSysMenuClient(cfg),
		SysMenuAction:         NewSysMenuActionClient(cfg),
		SysMenuActionResource: NewSysMenuActionResourceClient(cfg),
		SysRole:               NewSysRoleClient(cfg),
		SysRoleMenu:           NewSysRoleMenuClient(cfg),
		SysUser:               NewSysUserClient(cfg),
		SysUserRole:           NewSysUserRoleClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		SysMenu.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.SysMenu.Use(hooks...)
	c.SysMenuAction.Use(hooks...)
	c.SysMenuActionResource.Use(hooks...)
	c.SysRole.Use(hooks...)
	c.SysRoleMenu.Use(hooks...)
	c.SysUser.Use(hooks...)
	c.SysUserRole.Use(hooks...)
}

// SysMenuClient is a client for the SysMenu schema.
type SysMenuClient struct {
	config
}

// NewSysMenuClient returns a client for the SysMenu from the given config.
func NewSysMenuClient(c config) *SysMenuClient {
	return &SysMenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysmenu.Hooks(f(g(h())))`.
func (c *SysMenuClient) Use(hooks ...Hook) {
	c.hooks.SysMenu = append(c.hooks.SysMenu, hooks...)
}

// Create returns a create builder for SysMenu.
func (c *SysMenuClient) Create() *SysMenuCreate {
	mutation := newSysMenuMutation(c.config, OpCreate)
	return &SysMenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysMenu entities.
func (c *SysMenuClient) CreateBulk(builders ...*SysMenuCreate) *SysMenuCreateBulk {
	return &SysMenuCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysMenu.
func (c *SysMenuClient) Update() *SysMenuUpdate {
	mutation := newSysMenuMutation(c.config, OpUpdate)
	return &SysMenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysMenuClient) UpdateOne(sm *SysMenu) *SysMenuUpdateOne {
	mutation := newSysMenuMutation(c.config, OpUpdateOne, withSysMenu(sm))
	return &SysMenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysMenuClient) UpdateOneID(id int) *SysMenuUpdateOne {
	mutation := newSysMenuMutation(c.config, OpUpdateOne, withSysMenuID(id))
	return &SysMenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysMenu.
func (c *SysMenuClient) Delete() *SysMenuDelete {
	mutation := newSysMenuMutation(c.config, OpDelete)
	return &SysMenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysMenuClient) DeleteOne(sm *SysMenu) *SysMenuDeleteOne {
	return c.DeleteOneID(sm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysMenuClient) DeleteOneID(id int) *SysMenuDeleteOne {
	builder := c.Delete().Where(sysmenu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysMenuDeleteOne{builder}
}

// Query returns a query builder for SysMenu.
func (c *SysMenuClient) Query() *SysMenuQuery {
	return &SysMenuQuery{
		config: c.config,
	}
}

// Get returns a SysMenu entity by its id.
func (c *SysMenuClient) Get(ctx context.Context, id int) (*SysMenu, error) {
	return c.Query().Where(sysmenu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysMenuClient) GetX(ctx context.Context, id int) *SysMenu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysMenuClient) Hooks() []Hook {
	return c.hooks.SysMenu
}

// SysMenuActionClient is a client for the SysMenuAction schema.
type SysMenuActionClient struct {
	config
}

// NewSysMenuActionClient returns a client for the SysMenuAction from the given config.
func NewSysMenuActionClient(c config) *SysMenuActionClient {
	return &SysMenuActionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysmenuaction.Hooks(f(g(h())))`.
func (c *SysMenuActionClient) Use(hooks ...Hook) {
	c.hooks.SysMenuAction = append(c.hooks.SysMenuAction, hooks...)
}

// Create returns a create builder for SysMenuAction.
func (c *SysMenuActionClient) Create() *SysMenuActionCreate {
	mutation := newSysMenuActionMutation(c.config, OpCreate)
	return &SysMenuActionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysMenuAction entities.
func (c *SysMenuActionClient) CreateBulk(builders ...*SysMenuActionCreate) *SysMenuActionCreateBulk {
	return &SysMenuActionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysMenuAction.
func (c *SysMenuActionClient) Update() *SysMenuActionUpdate {
	mutation := newSysMenuActionMutation(c.config, OpUpdate)
	return &SysMenuActionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysMenuActionClient) UpdateOne(sma *SysMenuAction) *SysMenuActionUpdateOne {
	mutation := newSysMenuActionMutation(c.config, OpUpdateOne, withSysMenuAction(sma))
	return &SysMenuActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysMenuActionClient) UpdateOneID(id int) *SysMenuActionUpdateOne {
	mutation := newSysMenuActionMutation(c.config, OpUpdateOne, withSysMenuActionID(id))
	return &SysMenuActionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysMenuAction.
func (c *SysMenuActionClient) Delete() *SysMenuActionDelete {
	mutation := newSysMenuActionMutation(c.config, OpDelete)
	return &SysMenuActionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysMenuActionClient) DeleteOne(sma *SysMenuAction) *SysMenuActionDeleteOne {
	return c.DeleteOneID(sma.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysMenuActionClient) DeleteOneID(id int) *SysMenuActionDeleteOne {
	builder := c.Delete().Where(sysmenuaction.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysMenuActionDeleteOne{builder}
}

// Query returns a query builder for SysMenuAction.
func (c *SysMenuActionClient) Query() *SysMenuActionQuery {
	return &SysMenuActionQuery{
		config: c.config,
	}
}

// Get returns a SysMenuAction entity by its id.
func (c *SysMenuActionClient) Get(ctx context.Context, id int) (*SysMenuAction, error) {
	return c.Query().Where(sysmenuaction.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysMenuActionClient) GetX(ctx context.Context, id int) *SysMenuAction {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysMenuActionClient) Hooks() []Hook {
	return c.hooks.SysMenuAction
}

// SysMenuActionResourceClient is a client for the SysMenuActionResource schema.
type SysMenuActionResourceClient struct {
	config
}

// NewSysMenuActionResourceClient returns a client for the SysMenuActionResource from the given config.
func NewSysMenuActionResourceClient(c config) *SysMenuActionResourceClient {
	return &SysMenuActionResourceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysmenuactionresource.Hooks(f(g(h())))`.
func (c *SysMenuActionResourceClient) Use(hooks ...Hook) {
	c.hooks.SysMenuActionResource = append(c.hooks.SysMenuActionResource, hooks...)
}

// Create returns a create builder for SysMenuActionResource.
func (c *SysMenuActionResourceClient) Create() *SysMenuActionResourceCreate {
	mutation := newSysMenuActionResourceMutation(c.config, OpCreate)
	return &SysMenuActionResourceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysMenuActionResource entities.
func (c *SysMenuActionResourceClient) CreateBulk(builders ...*SysMenuActionResourceCreate) *SysMenuActionResourceCreateBulk {
	return &SysMenuActionResourceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysMenuActionResource.
func (c *SysMenuActionResourceClient) Update() *SysMenuActionResourceUpdate {
	mutation := newSysMenuActionResourceMutation(c.config, OpUpdate)
	return &SysMenuActionResourceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysMenuActionResourceClient) UpdateOne(smar *SysMenuActionResource) *SysMenuActionResourceUpdateOne {
	mutation := newSysMenuActionResourceMutation(c.config, OpUpdateOne, withSysMenuActionResource(smar))
	return &SysMenuActionResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysMenuActionResourceClient) UpdateOneID(id int) *SysMenuActionResourceUpdateOne {
	mutation := newSysMenuActionResourceMutation(c.config, OpUpdateOne, withSysMenuActionResourceID(id))
	return &SysMenuActionResourceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysMenuActionResource.
func (c *SysMenuActionResourceClient) Delete() *SysMenuActionResourceDelete {
	mutation := newSysMenuActionResourceMutation(c.config, OpDelete)
	return &SysMenuActionResourceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysMenuActionResourceClient) DeleteOne(smar *SysMenuActionResource) *SysMenuActionResourceDeleteOne {
	return c.DeleteOneID(smar.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysMenuActionResourceClient) DeleteOneID(id int) *SysMenuActionResourceDeleteOne {
	builder := c.Delete().Where(sysmenuactionresource.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysMenuActionResourceDeleteOne{builder}
}

// Query returns a query builder for SysMenuActionResource.
func (c *SysMenuActionResourceClient) Query() *SysMenuActionResourceQuery {
	return &SysMenuActionResourceQuery{
		config: c.config,
	}
}

// Get returns a SysMenuActionResource entity by its id.
func (c *SysMenuActionResourceClient) Get(ctx context.Context, id int) (*SysMenuActionResource, error) {
	return c.Query().Where(sysmenuactionresource.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysMenuActionResourceClient) GetX(ctx context.Context, id int) *SysMenuActionResource {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysMenuActionResourceClient) Hooks() []Hook {
	return c.hooks.SysMenuActionResource
}

// SysRoleClient is a client for the SysRole schema.
type SysRoleClient struct {
	config
}

// NewSysRoleClient returns a client for the SysRole from the given config.
func NewSysRoleClient(c config) *SysRoleClient {
	return &SysRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysrole.Hooks(f(g(h())))`.
func (c *SysRoleClient) Use(hooks ...Hook) {
	c.hooks.SysRole = append(c.hooks.SysRole, hooks...)
}

// Create returns a create builder for SysRole.
func (c *SysRoleClient) Create() *SysRoleCreate {
	mutation := newSysRoleMutation(c.config, OpCreate)
	return &SysRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysRole entities.
func (c *SysRoleClient) CreateBulk(builders ...*SysRoleCreate) *SysRoleCreateBulk {
	return &SysRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysRole.
func (c *SysRoleClient) Update() *SysRoleUpdate {
	mutation := newSysRoleMutation(c.config, OpUpdate)
	return &SysRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysRoleClient) UpdateOne(sr *SysRole) *SysRoleUpdateOne {
	mutation := newSysRoleMutation(c.config, OpUpdateOne, withSysRole(sr))
	return &SysRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysRoleClient) UpdateOneID(id int) *SysRoleUpdateOne {
	mutation := newSysRoleMutation(c.config, OpUpdateOne, withSysRoleID(id))
	return &SysRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysRole.
func (c *SysRoleClient) Delete() *SysRoleDelete {
	mutation := newSysRoleMutation(c.config, OpDelete)
	return &SysRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysRoleClient) DeleteOne(sr *SysRole) *SysRoleDeleteOne {
	return c.DeleteOneID(sr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysRoleClient) DeleteOneID(id int) *SysRoleDeleteOne {
	builder := c.Delete().Where(sysrole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysRoleDeleteOne{builder}
}

// Query returns a query builder for SysRole.
func (c *SysRoleClient) Query() *SysRoleQuery {
	return &SysRoleQuery{
		config: c.config,
	}
}

// Get returns a SysRole entity by its id.
func (c *SysRoleClient) Get(ctx context.Context, id int) (*SysRole, error) {
	return c.Query().Where(sysrole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysRoleClient) GetX(ctx context.Context, id int) *SysRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysRoleClient) Hooks() []Hook {
	return c.hooks.SysRole
}

// SysRoleMenuClient is a client for the SysRoleMenu schema.
type SysRoleMenuClient struct {
	config
}

// NewSysRoleMenuClient returns a client for the SysRoleMenu from the given config.
func NewSysRoleMenuClient(c config) *SysRoleMenuClient {
	return &SysRoleMenuClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysrolemenu.Hooks(f(g(h())))`.
func (c *SysRoleMenuClient) Use(hooks ...Hook) {
	c.hooks.SysRoleMenu = append(c.hooks.SysRoleMenu, hooks...)
}

// Create returns a create builder for SysRoleMenu.
func (c *SysRoleMenuClient) Create() *SysRoleMenuCreate {
	mutation := newSysRoleMenuMutation(c.config, OpCreate)
	return &SysRoleMenuCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysRoleMenu entities.
func (c *SysRoleMenuClient) CreateBulk(builders ...*SysRoleMenuCreate) *SysRoleMenuCreateBulk {
	return &SysRoleMenuCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysRoleMenu.
func (c *SysRoleMenuClient) Update() *SysRoleMenuUpdate {
	mutation := newSysRoleMenuMutation(c.config, OpUpdate)
	return &SysRoleMenuUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysRoleMenuClient) UpdateOne(srm *SysRoleMenu) *SysRoleMenuUpdateOne {
	mutation := newSysRoleMenuMutation(c.config, OpUpdateOne, withSysRoleMenu(srm))
	return &SysRoleMenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysRoleMenuClient) UpdateOneID(id int) *SysRoleMenuUpdateOne {
	mutation := newSysRoleMenuMutation(c.config, OpUpdateOne, withSysRoleMenuID(id))
	return &SysRoleMenuUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysRoleMenu.
func (c *SysRoleMenuClient) Delete() *SysRoleMenuDelete {
	mutation := newSysRoleMenuMutation(c.config, OpDelete)
	return &SysRoleMenuDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysRoleMenuClient) DeleteOne(srm *SysRoleMenu) *SysRoleMenuDeleteOne {
	return c.DeleteOneID(srm.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysRoleMenuClient) DeleteOneID(id int) *SysRoleMenuDeleteOne {
	builder := c.Delete().Where(sysrolemenu.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysRoleMenuDeleteOne{builder}
}

// Query returns a query builder for SysRoleMenu.
func (c *SysRoleMenuClient) Query() *SysRoleMenuQuery {
	return &SysRoleMenuQuery{
		config: c.config,
	}
}

// Get returns a SysRoleMenu entity by its id.
func (c *SysRoleMenuClient) Get(ctx context.Context, id int) (*SysRoleMenu, error) {
	return c.Query().Where(sysrolemenu.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysRoleMenuClient) GetX(ctx context.Context, id int) *SysRoleMenu {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysRoleMenuClient) Hooks() []Hook {
	return c.hooks.SysRoleMenu
}

// SysUserClient is a client for the SysUser schema.
type SysUserClient struct {
	config
}

// NewSysUserClient returns a client for the SysUser from the given config.
func NewSysUserClient(c config) *SysUserClient {
	return &SysUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysuser.Hooks(f(g(h())))`.
func (c *SysUserClient) Use(hooks ...Hook) {
	c.hooks.SysUser = append(c.hooks.SysUser, hooks...)
}

// Create returns a create builder for SysUser.
func (c *SysUserClient) Create() *SysUserCreate {
	mutation := newSysUserMutation(c.config, OpCreate)
	return &SysUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysUser entities.
func (c *SysUserClient) CreateBulk(builders ...*SysUserCreate) *SysUserCreateBulk {
	return &SysUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysUser.
func (c *SysUserClient) Update() *SysUserUpdate {
	mutation := newSysUserMutation(c.config, OpUpdate)
	return &SysUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysUserClient) UpdateOne(su *SysUser) *SysUserUpdateOne {
	mutation := newSysUserMutation(c.config, OpUpdateOne, withSysUser(su))
	return &SysUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysUserClient) UpdateOneID(id int) *SysUserUpdateOne {
	mutation := newSysUserMutation(c.config, OpUpdateOne, withSysUserID(id))
	return &SysUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysUser.
func (c *SysUserClient) Delete() *SysUserDelete {
	mutation := newSysUserMutation(c.config, OpDelete)
	return &SysUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysUserClient) DeleteOne(su *SysUser) *SysUserDeleteOne {
	return c.DeleteOneID(su.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysUserClient) DeleteOneID(id int) *SysUserDeleteOne {
	builder := c.Delete().Where(sysuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysUserDeleteOne{builder}
}

// Query returns a query builder for SysUser.
func (c *SysUserClient) Query() *SysUserQuery {
	return &SysUserQuery{
		config: c.config,
	}
}

// Get returns a SysUser entity by its id.
func (c *SysUserClient) Get(ctx context.Context, id int) (*SysUser, error) {
	return c.Query().Where(sysuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysUserClient) GetX(ctx context.Context, id int) *SysUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysUserClient) Hooks() []Hook {
	return c.hooks.SysUser
}

// SysUserRoleClient is a client for the SysUserRole schema.
type SysUserRoleClient struct {
	config
}

// NewSysUserRoleClient returns a client for the SysUserRole from the given config.
func NewSysUserRoleClient(c config) *SysUserRoleClient {
	return &SysUserRoleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `sysuserrole.Hooks(f(g(h())))`.
func (c *SysUserRoleClient) Use(hooks ...Hook) {
	c.hooks.SysUserRole = append(c.hooks.SysUserRole, hooks...)
}

// Create returns a create builder for SysUserRole.
func (c *SysUserRoleClient) Create() *SysUserRoleCreate {
	mutation := newSysUserRoleMutation(c.config, OpCreate)
	return &SysUserRoleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of SysUserRole entities.
func (c *SysUserRoleClient) CreateBulk(builders ...*SysUserRoleCreate) *SysUserRoleCreateBulk {
	return &SysUserRoleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for SysUserRole.
func (c *SysUserRoleClient) Update() *SysUserRoleUpdate {
	mutation := newSysUserRoleMutation(c.config, OpUpdate)
	return &SysUserRoleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SysUserRoleClient) UpdateOne(sur *SysUserRole) *SysUserRoleUpdateOne {
	mutation := newSysUserRoleMutation(c.config, OpUpdateOne, withSysUserRole(sur))
	return &SysUserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SysUserRoleClient) UpdateOneID(id int) *SysUserRoleUpdateOne {
	mutation := newSysUserRoleMutation(c.config, OpUpdateOne, withSysUserRoleID(id))
	return &SysUserRoleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for SysUserRole.
func (c *SysUserRoleClient) Delete() *SysUserRoleDelete {
	mutation := newSysUserRoleMutation(c.config, OpDelete)
	return &SysUserRoleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SysUserRoleClient) DeleteOne(sur *SysUserRole) *SysUserRoleDeleteOne {
	return c.DeleteOneID(sur.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SysUserRoleClient) DeleteOneID(id int) *SysUserRoleDeleteOne {
	builder := c.Delete().Where(sysuserrole.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SysUserRoleDeleteOne{builder}
}

// Query returns a query builder for SysUserRole.
func (c *SysUserRoleClient) Query() *SysUserRoleQuery {
	return &SysUserRoleQuery{
		config: c.config,
	}
}

// Get returns a SysUserRole entity by its id.
func (c *SysUserRoleClient) Get(ctx context.Context, id int) (*SysUserRole, error) {
	return c.Query().Where(sysuserrole.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SysUserRoleClient) GetX(ctx context.Context, id int) *SysUserRole {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SysUserRoleClient) Hooks() []Hook {
	return c.hooks.SysUserRole
}
