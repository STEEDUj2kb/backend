// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/STEEDUj2kb/platform/ent/migrate"

	"github.com/STEEDUj2kb/platform/ent/dailygaol"
	"github.com/STEEDUj2kb/platform/ent/study"
	"github.com/STEEDUj2kb/platform/ent/user"
	"github.com/STEEDUj2kb/platform/ent/weeklygaol"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// DailyGaol is the client for interacting with the DailyGaol builders.
	DailyGaol *DailyGaolClient
	// Study is the client for interacting with the Study builders.
	Study *StudyClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// WeeklyGaol is the client for interacting with the WeeklyGaol builders.
	WeeklyGaol *WeeklyGaolClient
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
	c.DailyGaol = NewDailyGaolClient(c.config)
	c.Study = NewStudyClient(c.config)
	c.User = NewUserClient(c.config)
	c.WeeklyGaol = NewWeeklyGaolClient(c.config)
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
		ctx:        ctx,
		config:     cfg,
		DailyGaol:  NewDailyGaolClient(cfg),
		Study:      NewStudyClient(cfg),
		User:       NewUserClient(cfg),
		WeeklyGaol: NewWeeklyGaolClient(cfg),
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
		config:     cfg,
		DailyGaol:  NewDailyGaolClient(cfg),
		Study:      NewStudyClient(cfg),
		User:       NewUserClient(cfg),
		WeeklyGaol: NewWeeklyGaolClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		DailyGaol.
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
	c.DailyGaol.Use(hooks...)
	c.Study.Use(hooks...)
	c.User.Use(hooks...)
	c.WeeklyGaol.Use(hooks...)
}

// DailyGaolClient is a client for the DailyGaol schema.
type DailyGaolClient struct {
	config
}

// NewDailyGaolClient returns a client for the DailyGaol from the given config.
func NewDailyGaolClient(c config) *DailyGaolClient {
	return &DailyGaolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `dailygaol.Hooks(f(g(h())))`.
func (c *DailyGaolClient) Use(hooks ...Hook) {
	c.hooks.DailyGaol = append(c.hooks.DailyGaol, hooks...)
}

// Create returns a create builder for DailyGaol.
func (c *DailyGaolClient) Create() *DailyGaolCreate {
	mutation := newDailyGaolMutation(c.config, OpCreate)
	return &DailyGaolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DailyGaol entities.
func (c *DailyGaolClient) CreateBulk(builders ...*DailyGaolCreate) *DailyGaolCreateBulk {
	return &DailyGaolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DailyGaol.
func (c *DailyGaolClient) Update() *DailyGaolUpdate {
	mutation := newDailyGaolMutation(c.config, OpUpdate)
	return &DailyGaolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DailyGaolClient) UpdateOne(dg *DailyGaol) *DailyGaolUpdateOne {
	mutation := newDailyGaolMutation(c.config, OpUpdateOne, withDailyGaol(dg))
	return &DailyGaolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DailyGaolClient) UpdateOneID(id int) *DailyGaolUpdateOne {
	mutation := newDailyGaolMutation(c.config, OpUpdateOne, withDailyGaolID(id))
	return &DailyGaolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DailyGaol.
func (c *DailyGaolClient) Delete() *DailyGaolDelete {
	mutation := newDailyGaolMutation(c.config, OpDelete)
	return &DailyGaolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DailyGaolClient) DeleteOne(dg *DailyGaol) *DailyGaolDeleteOne {
	return c.DeleteOneID(dg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DailyGaolClient) DeleteOneID(id int) *DailyGaolDeleteOne {
	builder := c.Delete().Where(dailygaol.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DailyGaolDeleteOne{builder}
}

// Query returns a query builder for DailyGaol.
func (c *DailyGaolClient) Query() *DailyGaolQuery {
	return &DailyGaolQuery{
		config: c.config,
	}
}

// Get returns a DailyGaol entity by its id.
func (c *DailyGaolClient) Get(ctx context.Context, id int) (*DailyGaol, error) {
	return c.Query().Where(dailygaol.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DailyGaolClient) GetX(ctx context.Context, id int) *DailyGaol {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStudy queries the study edge of a DailyGaol.
func (c *DailyGaolClient) QueryStudy(dg *DailyGaol) *StudyQuery {
	query := &StudyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dailygaol.Table, dailygaol.FieldID, id),
			sqlgraph.To(study.Table, study.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailygaol.StudyTable, dailygaol.StudyColumn),
		)
		fromV = sqlgraph.Neighbors(dg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWgoal queries the wgoal edge of a DailyGaol.
func (c *DailyGaolClient) QueryWgoal(dg *DailyGaol) *WeeklyGaolQuery {
	query := &WeeklyGaolQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := dg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(dailygaol.Table, dailygaol.FieldID, id),
			sqlgraph.To(weeklygaol.Table, weeklygaol.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, dailygaol.WgoalTable, dailygaol.WgoalColumn),
		)
		fromV = sqlgraph.Neighbors(dg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *DailyGaolClient) Hooks() []Hook {
	return c.hooks.DailyGaol
}

// StudyClient is a client for the Study schema.
type StudyClient struct {
	config
}

// NewStudyClient returns a client for the Study from the given config.
func NewStudyClient(c config) *StudyClient {
	return &StudyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `study.Hooks(f(g(h())))`.
func (c *StudyClient) Use(hooks ...Hook) {
	c.hooks.Study = append(c.hooks.Study, hooks...)
}

// Create returns a create builder for Study.
func (c *StudyClient) Create() *StudyCreate {
	mutation := newStudyMutation(c.config, OpCreate)
	return &StudyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Study entities.
func (c *StudyClient) CreateBulk(builders ...*StudyCreate) *StudyCreateBulk {
	return &StudyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Study.
func (c *StudyClient) Update() *StudyUpdate {
	mutation := newStudyMutation(c.config, OpUpdate)
	return &StudyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StudyClient) UpdateOne(s *Study) *StudyUpdateOne {
	mutation := newStudyMutation(c.config, OpUpdateOne, withStudy(s))
	return &StudyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StudyClient) UpdateOneID(id int) *StudyUpdateOne {
	mutation := newStudyMutation(c.config, OpUpdateOne, withStudyID(id))
	return &StudyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Study.
func (c *StudyClient) Delete() *StudyDelete {
	mutation := newStudyMutation(c.config, OpDelete)
	return &StudyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StudyClient) DeleteOne(s *Study) *StudyDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StudyClient) DeleteOneID(id int) *StudyDeleteOne {
	builder := c.Delete().Where(study.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StudyDeleteOne{builder}
}

// Query returns a query builder for Study.
func (c *StudyClient) Query() *StudyQuery {
	return &StudyQuery{
		config: c.config,
	}
}

// Get returns a Study entity by its id.
func (c *StudyClient) Get(ctx context.Context, id int) (*Study, error) {
	return c.Query().Where(study.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StudyClient) GetX(ctx context.Context, id int) *Study {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPlanner queries the planner edge of a Study.
func (c *StudyClient) QueryPlanner(s *Study) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(study.Table, study.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, study.PlannerTable, study.PlannerColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryDgoals queries the dgoals edge of a Study.
func (c *StudyClient) QueryDgoals(s *Study) *DailyGaolQuery {
	query := &DailyGaolQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(study.Table, study.FieldID, id),
			sqlgraph.To(dailygaol.Table, dailygaol.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, study.DgoalsTable, study.DgoalsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWgoals queries the wgoals edge of a Study.
func (c *StudyClient) QueryWgoals(s *Study) *WeeklyGaolQuery {
	query := &WeeklyGaolQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(study.Table, study.FieldID, id),
			sqlgraph.To(weeklygaol.Table, weeklygaol.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, study.WgoalsTable, study.WgoalsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StudyClient) Hooks() []Hook {
	return c.hooks.Study
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
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
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStudies queries the studies edge of a User.
func (c *UserClient) QueryStudies(u *User) *StudyQuery {
	query := &StudyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(study.Table, study.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.StudiesTable, user.StudiesColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// WeeklyGaolClient is a client for the WeeklyGaol schema.
type WeeklyGaolClient struct {
	config
}

// NewWeeklyGaolClient returns a client for the WeeklyGaol from the given config.
func NewWeeklyGaolClient(c config) *WeeklyGaolClient {
	return &WeeklyGaolClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `weeklygaol.Hooks(f(g(h())))`.
func (c *WeeklyGaolClient) Use(hooks ...Hook) {
	c.hooks.WeeklyGaol = append(c.hooks.WeeklyGaol, hooks...)
}

// Create returns a create builder for WeeklyGaol.
func (c *WeeklyGaolClient) Create() *WeeklyGaolCreate {
	mutation := newWeeklyGaolMutation(c.config, OpCreate)
	return &WeeklyGaolCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of WeeklyGaol entities.
func (c *WeeklyGaolClient) CreateBulk(builders ...*WeeklyGaolCreate) *WeeklyGaolCreateBulk {
	return &WeeklyGaolCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for WeeklyGaol.
func (c *WeeklyGaolClient) Update() *WeeklyGaolUpdate {
	mutation := newWeeklyGaolMutation(c.config, OpUpdate)
	return &WeeklyGaolUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WeeklyGaolClient) UpdateOne(wg *WeeklyGaol) *WeeklyGaolUpdateOne {
	mutation := newWeeklyGaolMutation(c.config, OpUpdateOne, withWeeklyGaol(wg))
	return &WeeklyGaolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WeeklyGaolClient) UpdateOneID(id int) *WeeklyGaolUpdateOne {
	mutation := newWeeklyGaolMutation(c.config, OpUpdateOne, withWeeklyGaolID(id))
	return &WeeklyGaolUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for WeeklyGaol.
func (c *WeeklyGaolClient) Delete() *WeeklyGaolDelete {
	mutation := newWeeklyGaolMutation(c.config, OpDelete)
	return &WeeklyGaolDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *WeeklyGaolClient) DeleteOne(wg *WeeklyGaol) *WeeklyGaolDeleteOne {
	return c.DeleteOneID(wg.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *WeeklyGaolClient) DeleteOneID(id int) *WeeklyGaolDeleteOne {
	builder := c.Delete().Where(weeklygaol.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WeeklyGaolDeleteOne{builder}
}

// Query returns a query builder for WeeklyGaol.
func (c *WeeklyGaolClient) Query() *WeeklyGaolQuery {
	return &WeeklyGaolQuery{
		config: c.config,
	}
}

// Get returns a WeeklyGaol entity by its id.
func (c *WeeklyGaolClient) Get(ctx context.Context, id int) (*WeeklyGaol, error) {
	return c.Query().Where(weeklygaol.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WeeklyGaolClient) GetX(ctx context.Context, id int) *WeeklyGaol {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryDgaols queries the dgaols edge of a WeeklyGaol.
func (c *WeeklyGaolClient) QueryDgaols(wg *WeeklyGaol) *DailyGaolQuery {
	query := &DailyGaolQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(weeklygaol.Table, weeklygaol.FieldID, id),
			sqlgraph.To(dailygaol.Table, dailygaol.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, weeklygaol.DgaolsTable, weeklygaol.DgaolsColumn),
		)
		fromV = sqlgraph.Neighbors(wg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryStudy queries the study edge of a WeeklyGaol.
func (c *WeeklyGaolClient) QueryStudy(wg *WeeklyGaol) *StudyQuery {
	query := &StudyQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := wg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(weeklygaol.Table, weeklygaol.FieldID, id),
			sqlgraph.To(study.Table, study.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, weeklygaol.StudyTable, weeklygaol.StudyColumn),
		)
		fromV = sqlgraph.Neighbors(wg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WeeklyGaolClient) Hooks() []Hook {
	return c.hooks.WeeklyGaol
}
