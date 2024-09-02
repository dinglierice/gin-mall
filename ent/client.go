// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"mall/ent/migrate"

	"mall/ent/psconfig"
	"mall/ent/psstrategy"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// PsConfig is the client for interacting with the PsConfig builders.
	PsConfig *PsConfigClient
	// PsStrategy is the client for interacting with the PsStrategy builders.
	PsStrategy *PsStrategyClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.PsConfig = NewPsConfigClient(c.config)
	c.PsStrategy = NewPsStrategyClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		PsConfig:   NewPsConfigClient(cfg),
		PsStrategy: NewPsStrategyClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
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
		ctx:        ctx,
		config:     cfg,
		PsConfig:   NewPsConfigClient(cfg),
		PsStrategy: NewPsStrategyClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		PsConfig.
//		Query().
//		Count(ctx)
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
	c.PsConfig.Use(hooks...)
	c.PsStrategy.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.PsConfig.Intercept(interceptors...)
	c.PsStrategy.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PsConfigMutation:
		return c.PsConfig.mutate(ctx, m)
	case *PsStrategyMutation:
		return c.PsStrategy.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PsConfigClient is a client for the PsConfig schema.
type PsConfigClient struct {
	config
}

// NewPsConfigClient returns a client for the PsConfig from the given config.
func NewPsConfigClient(c config) *PsConfigClient {
	return &PsConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `psconfig.Hooks(f(g(h())))`.
func (c *PsConfigClient) Use(hooks ...Hook) {
	c.hooks.PsConfig = append(c.hooks.PsConfig, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `psconfig.Intercept(f(g(h())))`.
func (c *PsConfigClient) Intercept(interceptors ...Interceptor) {
	c.inters.PsConfig = append(c.inters.PsConfig, interceptors...)
}

// Create returns a builder for creating a PsConfig entity.
func (c *PsConfigClient) Create() *PsConfigCreate {
	mutation := newPsConfigMutation(c.config, OpCreate)
	return &PsConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PsConfig entities.
func (c *PsConfigClient) CreateBulk(builders ...*PsConfigCreate) *PsConfigCreateBulk {
	return &PsConfigCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PsConfigClient) MapCreateBulk(slice any, setFunc func(*PsConfigCreate, int)) *PsConfigCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PsConfigCreateBulk{err: fmt.Errorf("calling to PsConfigClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PsConfigCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PsConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PsConfig.
func (c *PsConfigClient) Update() *PsConfigUpdate {
	mutation := newPsConfigMutation(c.config, OpUpdate)
	return &PsConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PsConfigClient) UpdateOne(pc *PsConfig) *PsConfigUpdateOne {
	mutation := newPsConfigMutation(c.config, OpUpdateOne, withPsConfig(pc))
	return &PsConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PsConfigClient) UpdateOneID(id int) *PsConfigUpdateOne {
	mutation := newPsConfigMutation(c.config, OpUpdateOne, withPsConfigID(id))
	return &PsConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PsConfig.
func (c *PsConfigClient) Delete() *PsConfigDelete {
	mutation := newPsConfigMutation(c.config, OpDelete)
	return &PsConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PsConfigClient) DeleteOne(pc *PsConfig) *PsConfigDeleteOne {
	return c.DeleteOneID(pc.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PsConfigClient) DeleteOneID(id int) *PsConfigDeleteOne {
	builder := c.Delete().Where(psconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PsConfigDeleteOne{builder}
}

// Query returns a query builder for PsConfig.
func (c *PsConfigClient) Query() *PsConfigQuery {
	return &PsConfigQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePsConfig},
		inters: c.Interceptors(),
	}
}

// Get returns a PsConfig entity by its id.
func (c *PsConfigClient) Get(ctx context.Context, id int) (*PsConfig, error) {
	return c.Query().Where(psconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PsConfigClient) GetX(ctx context.Context, id int) *PsConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PsConfigClient) Hooks() []Hook {
	return c.hooks.PsConfig
}

// Interceptors returns the client interceptors.
func (c *PsConfigClient) Interceptors() []Interceptor {
	return c.inters.PsConfig
}

func (c *PsConfigClient) mutate(ctx context.Context, m *PsConfigMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PsConfigCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PsConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PsConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PsConfigDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PsConfig mutation op: %q", m.Op())
	}
}

// PsStrategyClient is a client for the PsStrategy schema.
type PsStrategyClient struct {
	config
}

// NewPsStrategyClient returns a client for the PsStrategy from the given config.
func NewPsStrategyClient(c config) *PsStrategyClient {
	return &PsStrategyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `psstrategy.Hooks(f(g(h())))`.
func (c *PsStrategyClient) Use(hooks ...Hook) {
	c.hooks.PsStrategy = append(c.hooks.PsStrategy, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `psstrategy.Intercept(f(g(h())))`.
func (c *PsStrategyClient) Intercept(interceptors ...Interceptor) {
	c.inters.PsStrategy = append(c.inters.PsStrategy, interceptors...)
}

// Create returns a builder for creating a PsStrategy entity.
func (c *PsStrategyClient) Create() *PsStrategyCreate {
	mutation := newPsStrategyMutation(c.config, OpCreate)
	return &PsStrategyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PsStrategy entities.
func (c *PsStrategyClient) CreateBulk(builders ...*PsStrategyCreate) *PsStrategyCreateBulk {
	return &PsStrategyCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PsStrategyClient) MapCreateBulk(slice any, setFunc func(*PsStrategyCreate, int)) *PsStrategyCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PsStrategyCreateBulk{err: fmt.Errorf("calling to PsStrategyClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PsStrategyCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PsStrategyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PsStrategy.
func (c *PsStrategyClient) Update() *PsStrategyUpdate {
	mutation := newPsStrategyMutation(c.config, OpUpdate)
	return &PsStrategyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PsStrategyClient) UpdateOne(ps *PsStrategy) *PsStrategyUpdateOne {
	mutation := newPsStrategyMutation(c.config, OpUpdateOne, withPsStrategy(ps))
	return &PsStrategyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PsStrategyClient) UpdateOneID(id int) *PsStrategyUpdateOne {
	mutation := newPsStrategyMutation(c.config, OpUpdateOne, withPsStrategyID(id))
	return &PsStrategyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PsStrategy.
func (c *PsStrategyClient) Delete() *PsStrategyDelete {
	mutation := newPsStrategyMutation(c.config, OpDelete)
	return &PsStrategyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PsStrategyClient) DeleteOne(ps *PsStrategy) *PsStrategyDeleteOne {
	return c.DeleteOneID(ps.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PsStrategyClient) DeleteOneID(id int) *PsStrategyDeleteOne {
	builder := c.Delete().Where(psstrategy.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PsStrategyDeleteOne{builder}
}

// Query returns a query builder for PsStrategy.
func (c *PsStrategyClient) Query() *PsStrategyQuery {
	return &PsStrategyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePsStrategy},
		inters: c.Interceptors(),
	}
}

// Get returns a PsStrategy entity by its id.
func (c *PsStrategyClient) Get(ctx context.Context, id int) (*PsStrategy, error) {
	return c.Query().Where(psstrategy.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PsStrategyClient) GetX(ctx context.Context, id int) *PsStrategy {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PsStrategyClient) Hooks() []Hook {
	return c.hooks.PsStrategy
}

// Interceptors returns the client interceptors.
func (c *PsStrategyClient) Interceptors() []Interceptor {
	return c.inters.PsStrategy
}

func (c *PsStrategyClient) mutate(ctx context.Context, m *PsStrategyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PsStrategyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PsStrategyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PsStrategyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PsStrategyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PsStrategy mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		PsConfig, PsStrategy []ent.Hook
	}
	inters struct {
		PsConfig, PsStrategy []ent.Interceptor
	}
)
