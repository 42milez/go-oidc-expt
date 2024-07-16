// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/42milez/go-oidc-expt/pkg/ent/ent/migrate"
	"github.com/42milez/go-oidc-expt/pkg/typedef"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/authcode"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/consent"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/redirecturi"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/relyingparty"
	"github.com/42milez/go-oidc-expt/pkg/ent/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// AuthCode is the client for interacting with the AuthCode builders.
	AuthCode *AuthCodeClient
	// Consent is the client for interacting with the Consent builders.
	Consent *ConsentClient
	// RedirectURI is the client for interacting with the RedirectURI builders.
	RedirectURI *RedirectURIClient
	// RelyingParty is the client for interacting with the RelyingParty builders.
	RelyingParty *RelyingPartyClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.AuthCode = NewAuthCodeClient(c.config)
	c.Consent = NewConsentClient(c.config)
	c.RedirectURI = NewRedirectURIClient(c.config)
	c.RelyingParty = NewRelyingPartyClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:          ctx,
		config:       cfg,
		AuthCode:     NewAuthCodeClient(cfg),
		Consent:      NewConsentClient(cfg),
		RedirectURI:  NewRedirectURIClient(cfg),
		RelyingParty: NewRelyingPartyClient(cfg),
		User:         NewUserClient(cfg),
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
		ctx:          ctx,
		config:       cfg,
		AuthCode:     NewAuthCodeClient(cfg),
		Consent:      NewConsentClient(cfg),
		RedirectURI:  NewRedirectURIClient(cfg),
		RelyingParty: NewRelyingPartyClient(cfg),
		User:         NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		AuthCode.
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
	c.AuthCode.Use(hooks...)
	c.Consent.Use(hooks...)
	c.RedirectURI.Use(hooks...)
	c.RelyingParty.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.AuthCode.Intercept(interceptors...)
	c.Consent.Intercept(interceptors...)
	c.RedirectURI.Intercept(interceptors...)
	c.RelyingParty.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AuthCodeMutation:
		return c.AuthCode.mutate(ctx, m)
	case *ConsentMutation:
		return c.Consent.mutate(ctx, m)
	case *RedirectURIMutation:
		return c.RedirectURI.mutate(ctx, m)
	case *RelyingPartyMutation:
		return c.RelyingParty.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AuthCodeClient is a client for the AuthCode schema.
type AuthCodeClient struct {
	config
}

// NewAuthCodeClient returns a client for the AuthCode from the given config.
func NewAuthCodeClient(c config) *AuthCodeClient {
	return &AuthCodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `authcode.Hooks(f(g(h())))`.
func (c *AuthCodeClient) Use(hooks ...Hook) {
	c.hooks.AuthCode = append(c.hooks.AuthCode, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `authcode.Intercept(f(g(h())))`.
func (c *AuthCodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.AuthCode = append(c.inters.AuthCode, interceptors...)
}

// Create returns a builder for creating a AuthCode entity.
func (c *AuthCodeClient) Create() *AuthCodeCreate {
	mutation := newAuthCodeMutation(c.config, OpCreate)
	return &AuthCodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AuthCode entities.
func (c *AuthCodeClient) CreateBulk(builders ...*AuthCodeCreate) *AuthCodeCreateBulk {
	return &AuthCodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AuthCodeClient) MapCreateBulk(slice any, setFunc func(*AuthCodeCreate, int)) *AuthCodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AuthCodeCreateBulk{err: fmt.Errorf("calling to AuthCodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AuthCodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AuthCodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AuthCode.
func (c *AuthCodeClient) Update() *AuthCodeUpdate {
	mutation := newAuthCodeMutation(c.config, OpUpdate)
	return &AuthCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AuthCodeClient) UpdateOne(ac *AuthCode) *AuthCodeUpdateOne {
	mutation := newAuthCodeMutation(c.config, OpUpdateOne, withAuthCode(ac))
	return &AuthCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AuthCodeClient) UpdateOneID(id typedef.AuthCodeID) *AuthCodeUpdateOne {
	mutation := newAuthCodeMutation(c.config, OpUpdateOne, withAuthCodeID(id))
	return &AuthCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AuthCode.
func (c *AuthCodeClient) Delete() *AuthCodeDelete {
	mutation := newAuthCodeMutation(c.config, OpDelete)
	return &AuthCodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AuthCodeClient) DeleteOne(ac *AuthCode) *AuthCodeDeleteOne {
	return c.DeleteOneID(ac.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AuthCodeClient) DeleteOneID(id typedef.AuthCodeID) *AuthCodeDeleteOne {
	builder := c.Delete().Where(authcode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AuthCodeDeleteOne{builder}
}

// Query returns a query builder for AuthCode.
func (c *AuthCodeClient) Query() *AuthCodeQuery {
	return &AuthCodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAuthCode},
		inters: c.Interceptors(),
	}
}

// Get returns a AuthCode entity by its id.
func (c *AuthCodeClient) Get(ctx context.Context, id typedef.AuthCodeID) (*AuthCode, error) {
	return c.Query().Where(authcode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AuthCodeClient) GetX(ctx context.Context, id typedef.AuthCodeID) *AuthCode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRelyingParty queries the relying_party edge of a AuthCode.
func (c *AuthCodeClient) QueryRelyingParty(ac *AuthCode) *RelyingPartyQuery {
	query := (&RelyingPartyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ac.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(authcode.Table, authcode.FieldID, id),
			sqlgraph.To(relyingparty.Table, relyingparty.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, authcode.RelyingPartyTable, authcode.RelyingPartyColumn),
		)
		fromV = sqlgraph.Neighbors(ac.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AuthCodeClient) Hooks() []Hook {
	return c.hooks.AuthCode
}

// Interceptors returns the client interceptors.
func (c *AuthCodeClient) Interceptors() []Interceptor {
	return c.inters.AuthCode
}

func (c *AuthCodeClient) mutate(ctx context.Context, m *AuthCodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AuthCodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AuthCodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AuthCodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AuthCodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown AuthCode mutation op: %q", m.Op())
	}
}

// ConsentClient is a client for the Consent schema.
type ConsentClient struct {
	config
}

// NewConsentClient returns a client for the Consent from the given config.
func NewConsentClient(c config) *ConsentClient {
	return &ConsentClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `consent.Hooks(f(g(h())))`.
func (c *ConsentClient) Use(hooks ...Hook) {
	c.hooks.Consent = append(c.hooks.Consent, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `consent.Intercept(f(g(h())))`.
func (c *ConsentClient) Intercept(interceptors ...Interceptor) {
	c.inters.Consent = append(c.inters.Consent, interceptors...)
}

// Create returns a builder for creating a Consent entity.
func (c *ConsentClient) Create() *ConsentCreate {
	mutation := newConsentMutation(c.config, OpCreate)
	return &ConsentCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Consent entities.
func (c *ConsentClient) CreateBulk(builders ...*ConsentCreate) *ConsentCreateBulk {
	return &ConsentCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ConsentClient) MapCreateBulk(slice any, setFunc func(*ConsentCreate, int)) *ConsentCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ConsentCreateBulk{err: fmt.Errorf("calling to ConsentClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ConsentCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ConsentCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Consent.
func (c *ConsentClient) Update() *ConsentUpdate {
	mutation := newConsentMutation(c.config, OpUpdate)
	return &ConsentUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConsentClient) UpdateOne(co *Consent) *ConsentUpdateOne {
	mutation := newConsentMutation(c.config, OpUpdateOne, withConsent(co))
	return &ConsentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConsentClient) UpdateOneID(id typedef.ConsentID) *ConsentUpdateOne {
	mutation := newConsentMutation(c.config, OpUpdateOne, withConsentID(id))
	return &ConsentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Consent.
func (c *ConsentClient) Delete() *ConsentDelete {
	mutation := newConsentMutation(c.config, OpDelete)
	return &ConsentDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConsentClient) DeleteOne(co *Consent) *ConsentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ConsentClient) DeleteOneID(id typedef.ConsentID) *ConsentDeleteOne {
	builder := c.Delete().Where(consent.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConsentDeleteOne{builder}
}

// Query returns a query builder for Consent.
func (c *ConsentClient) Query() *ConsentQuery {
	return &ConsentQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeConsent},
		inters: c.Interceptors(),
	}
}

// Get returns a Consent entity by its id.
func (c *ConsentClient) Get(ctx context.Context, id typedef.ConsentID) (*Consent, error) {
	return c.Query().Where(consent.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConsentClient) GetX(ctx context.Context, id typedef.ConsentID) *Consent {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Consent.
func (c *ConsentClient) QueryUser(co *Consent) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(consent.Table, consent.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, consent.UserTable, consent.UserColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConsentClient) Hooks() []Hook {
	return c.hooks.Consent
}

// Interceptors returns the client interceptors.
func (c *ConsentClient) Interceptors() []Interceptor {
	return c.inters.Consent
}

func (c *ConsentClient) mutate(ctx context.Context, m *ConsentMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ConsentCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ConsentUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ConsentUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ConsentDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Consent mutation op: %q", m.Op())
	}
}

// RedirectURIClient is a client for the RedirectURI schema.
type RedirectURIClient struct {
	config
}

// NewRedirectURIClient returns a client for the RedirectURI from the given config.
func NewRedirectURIClient(c config) *RedirectURIClient {
	return &RedirectURIClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `redirecturi.Hooks(f(g(h())))`.
func (c *RedirectURIClient) Use(hooks ...Hook) {
	c.hooks.RedirectURI = append(c.hooks.RedirectURI, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `redirecturi.Intercept(f(g(h())))`.
func (c *RedirectURIClient) Intercept(interceptors ...Interceptor) {
	c.inters.RedirectURI = append(c.inters.RedirectURI, interceptors...)
}

// Create returns a builder for creating a RedirectURI entity.
func (c *RedirectURIClient) Create() *RedirectURICreate {
	mutation := newRedirectURIMutation(c.config, OpCreate)
	return &RedirectURICreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RedirectURI entities.
func (c *RedirectURIClient) CreateBulk(builders ...*RedirectURICreate) *RedirectURICreateBulk {
	return &RedirectURICreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RedirectURIClient) MapCreateBulk(slice any, setFunc func(*RedirectURICreate, int)) *RedirectURICreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RedirectURICreateBulk{err: fmt.Errorf("calling to RedirectURIClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RedirectURICreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RedirectURICreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RedirectURI.
func (c *RedirectURIClient) Update() *RedirectURIUpdate {
	mutation := newRedirectURIMutation(c.config, OpUpdate)
	return &RedirectURIUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RedirectURIClient) UpdateOne(ru *RedirectURI) *RedirectURIUpdateOne {
	mutation := newRedirectURIMutation(c.config, OpUpdateOne, withRedirectURI(ru))
	return &RedirectURIUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RedirectURIClient) UpdateOneID(id typedef.RedirectURIID) *RedirectURIUpdateOne {
	mutation := newRedirectURIMutation(c.config, OpUpdateOne, withRedirectURIID(id))
	return &RedirectURIUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RedirectURI.
func (c *RedirectURIClient) Delete() *RedirectURIDelete {
	mutation := newRedirectURIMutation(c.config, OpDelete)
	return &RedirectURIDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RedirectURIClient) DeleteOne(ru *RedirectURI) *RedirectURIDeleteOne {
	return c.DeleteOneID(ru.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RedirectURIClient) DeleteOneID(id typedef.RedirectURIID) *RedirectURIDeleteOne {
	builder := c.Delete().Where(redirecturi.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RedirectURIDeleteOne{builder}
}

// Query returns a query builder for RedirectURI.
func (c *RedirectURIClient) Query() *RedirectURIQuery {
	return &RedirectURIQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRedirectURI},
		inters: c.Interceptors(),
	}
}

// Get returns a RedirectURI entity by its id.
func (c *RedirectURIClient) Get(ctx context.Context, id typedef.RedirectURIID) (*RedirectURI, error) {
	return c.Query().Where(redirecturi.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RedirectURIClient) GetX(ctx context.Context, id typedef.RedirectURIID) *RedirectURI {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRelyingParty queries the relying_party edge of a RedirectURI.
func (c *RedirectURIClient) QueryRelyingParty(ru *RedirectURI) *RelyingPartyQuery {
	query := (&RelyingPartyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ru.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(redirecturi.Table, redirecturi.FieldID, id),
			sqlgraph.To(relyingparty.Table, relyingparty.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, redirecturi.RelyingPartyTable, redirecturi.RelyingPartyColumn),
		)
		fromV = sqlgraph.Neighbors(ru.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RedirectURIClient) Hooks() []Hook {
	return c.hooks.RedirectURI
}

// Interceptors returns the client interceptors.
func (c *RedirectURIClient) Interceptors() []Interceptor {
	return c.inters.RedirectURI
}

func (c *RedirectURIClient) mutate(ctx context.Context, m *RedirectURIMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RedirectURICreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RedirectURIUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RedirectURIUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RedirectURIDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RedirectURI mutation op: %q", m.Op())
	}
}

// RelyingPartyClient is a client for the RelyingParty schema.
type RelyingPartyClient struct {
	config
}

// NewRelyingPartyClient returns a client for the RelyingParty from the given config.
func NewRelyingPartyClient(c config) *RelyingPartyClient {
	return &RelyingPartyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `relyingparty.Hooks(f(g(h())))`.
func (c *RelyingPartyClient) Use(hooks ...Hook) {
	c.hooks.RelyingParty = append(c.hooks.RelyingParty, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `relyingparty.Intercept(f(g(h())))`.
func (c *RelyingPartyClient) Intercept(interceptors ...Interceptor) {
	c.inters.RelyingParty = append(c.inters.RelyingParty, interceptors...)
}

// Create returns a builder for creating a RelyingParty entity.
func (c *RelyingPartyClient) Create() *RelyingPartyCreate {
	mutation := newRelyingPartyMutation(c.config, OpCreate)
	return &RelyingPartyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RelyingParty entities.
func (c *RelyingPartyClient) CreateBulk(builders ...*RelyingPartyCreate) *RelyingPartyCreateBulk {
	return &RelyingPartyCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RelyingPartyClient) MapCreateBulk(slice any, setFunc func(*RelyingPartyCreate, int)) *RelyingPartyCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RelyingPartyCreateBulk{err: fmt.Errorf("calling to RelyingPartyClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RelyingPartyCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RelyingPartyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RelyingParty.
func (c *RelyingPartyClient) Update() *RelyingPartyUpdate {
	mutation := newRelyingPartyMutation(c.config, OpUpdate)
	return &RelyingPartyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RelyingPartyClient) UpdateOne(rp *RelyingParty) *RelyingPartyUpdateOne {
	mutation := newRelyingPartyMutation(c.config, OpUpdateOne, withRelyingParty(rp))
	return &RelyingPartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RelyingPartyClient) UpdateOneID(id typedef.RelyingPartyID) *RelyingPartyUpdateOne {
	mutation := newRelyingPartyMutation(c.config, OpUpdateOne, withRelyingPartyID(id))
	return &RelyingPartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RelyingParty.
func (c *RelyingPartyClient) Delete() *RelyingPartyDelete {
	mutation := newRelyingPartyMutation(c.config, OpDelete)
	return &RelyingPartyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RelyingPartyClient) DeleteOne(rp *RelyingParty) *RelyingPartyDeleteOne {
	return c.DeleteOneID(rp.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RelyingPartyClient) DeleteOneID(id typedef.RelyingPartyID) *RelyingPartyDeleteOne {
	builder := c.Delete().Where(relyingparty.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RelyingPartyDeleteOne{builder}
}

// Query returns a query builder for RelyingParty.
func (c *RelyingPartyClient) Query() *RelyingPartyQuery {
	return &RelyingPartyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRelyingParty},
		inters: c.Interceptors(),
	}
}

// Get returns a RelyingParty entity by its id.
func (c *RelyingPartyClient) Get(ctx context.Context, id typedef.RelyingPartyID) (*RelyingParty, error) {
	return c.Query().Where(relyingparty.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RelyingPartyClient) GetX(ctx context.Context, id typedef.RelyingPartyID) *RelyingParty {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAuthCodes queries the auth_codes edge of a RelyingParty.
func (c *RelyingPartyClient) QueryAuthCodes(rp *RelyingParty) *AuthCodeQuery {
	query := (&AuthCodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relyingparty.Table, relyingparty.FieldID, id),
			sqlgraph.To(authcode.Table, authcode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, relyingparty.AuthCodesTable, relyingparty.AuthCodesColumn),
		)
		fromV = sqlgraph.Neighbors(rp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryRedirectUris queries the redirect_uris edge of a RelyingParty.
func (c *RelyingPartyClient) QueryRedirectUris(rp *RelyingParty) *RedirectURIQuery {
	query := (&RedirectURIClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rp.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(relyingparty.Table, relyingparty.FieldID, id),
			sqlgraph.To(redirecturi.Table, redirecturi.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, relyingparty.RedirectUrisTable, relyingparty.RedirectUrisColumn),
		)
		fromV = sqlgraph.Neighbors(rp.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RelyingPartyClient) Hooks() []Hook {
	return c.hooks.RelyingParty
}

// Interceptors returns the client interceptors.
func (c *RelyingPartyClient) Interceptors() []Interceptor {
	return c.inters.RelyingParty
}

func (c *RelyingPartyClient) mutate(ctx context.Context, m *RelyingPartyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RelyingPartyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RelyingPartyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RelyingPartyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RelyingPartyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown RelyingParty mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id typedef.UserID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id typedef.UserID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id typedef.UserID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id typedef.UserID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryConsents queries the consents edge of a User.
func (c *UserClient) QueryConsents(u *User) *ConsentQuery {
	query := (&ConsentClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(consent.Table, consent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ConsentsTable, user.ConsentsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	hooks := c.hooks.User
	return append(hooks[:len(hooks):len(hooks)], user.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		AuthCode, Consent, RedirectURI, RelyingParty, User []ent.Hook
	}
	inters struct {
		AuthCode, Consent, RedirectURI, RelyingParty, User []ent.Interceptor
	}
)
