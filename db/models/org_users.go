// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package dbmodels

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// OrgUser is an object representing the database table.
type OrgUser struct {
	ID         string `boil:"id" json:"id" toml:"id" yaml:"id"`
	InternalID int    `boil:"internal_id" json:"internal_id" toml:"internal_id" yaml:"internal_id"`
	UserID     int    `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Admin      bool   `boil:"admin" json:"admin" toml:"admin" yaml:"admin"`
	OrgID      int    `boil:"org_id" json:"org_id" toml:"org_id" yaml:"org_id"`

	R *orgUserR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orgUserL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrgUserColumns = struct {
	ID         string
	InternalID string
	UserID     string
	Admin      string
	OrgID      string
}{
	ID:         "id",
	InternalID: "internal_id",
	UserID:     "user_id",
	Admin:      "admin",
	OrgID:      "org_id",
}

var OrgUserTableColumns = struct {
	ID         string
	InternalID string
	UserID     string
	Admin      string
	OrgID      string
}{
	ID:         "org_users.id",
	InternalID: "org_users.internal_id",
	UserID:     "org_users.user_id",
	Admin:      "org_users.admin",
	OrgID:      "org_users.org_id",
}

// Generated where

var OrgUserWhere = struct {
	ID         whereHelperstring
	InternalID whereHelperint
	UserID     whereHelperint
	Admin      whereHelperbool
	OrgID      whereHelperint
}{
	ID:         whereHelperstring{field: "\"org_users\".\"id\""},
	InternalID: whereHelperint{field: "\"org_users\".\"internal_id\""},
	UserID:     whereHelperint{field: "\"org_users\".\"user_id\""},
	Admin:      whereHelperbool{field: "\"org_users\".\"admin\""},
	OrgID:      whereHelperint{field: "\"org_users\".\"org_id\""},
}

// OrgUserRels is where relationship names are stored.
var OrgUserRels = struct {
	User string
	Org  string
}{
	User: "User",
	Org:  "Org",
}

// orgUserR is where relationships are stored.
type orgUserR struct {
	User *AbstractUser `boil:"User" json:"User" toml:"User" yaml:"User"`
	Org  *Org          `boil:"Org" json:"Org" toml:"Org" yaml:"Org"`
}

// NewStruct creates a new relationship struct
func (*orgUserR) NewStruct() *orgUserR {
	return &orgUserR{}
}

func (r *orgUserR) GetUser() *AbstractUser {
	if r == nil {
		return nil
	}
	return r.User
}

func (r *orgUserR) GetOrg() *Org {
	if r == nil {
		return nil
	}
	return r.Org
}

// orgUserL is where Load methods for each relationship are stored.
type orgUserL struct{}

var (
	orgUserAllColumns            = []string{"id", "internal_id", "user_id", "admin", "org_id"}
	orgUserColumnsWithoutDefault = []string{"id"}
	orgUserColumnsWithDefault    = []string{"internal_id", "user_id", "admin", "org_id"}
	orgUserPrimaryKeyColumns     = []string{"id"}
	orgUserGeneratedColumns      = []string{}
)

type (
	// OrgUserSlice is an alias for a slice of pointers to OrgUser.
	// This should almost always be used instead of []OrgUser.
	OrgUserSlice []*OrgUser
	// OrgUserHook is the signature for custom OrgUser hook methods
	OrgUserHook func(context.Context, boil.ContextExecutor, *OrgUser) error

	orgUserQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orgUserType                 = reflect.TypeOf(&OrgUser{})
	orgUserMapping              = queries.MakeStructMapping(orgUserType)
	orgUserPrimaryKeyMapping, _ = queries.BindMapping(orgUserType, orgUserMapping, orgUserPrimaryKeyColumns)
	orgUserInsertCacheMut       sync.RWMutex
	orgUserInsertCache          = make(map[string]insertCache)
	orgUserUpdateCacheMut       sync.RWMutex
	orgUserUpdateCache          = make(map[string]updateCache)
	orgUserUpsertCacheMut       sync.RWMutex
	orgUserUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var orgUserAfterSelectHooks []OrgUserHook

var orgUserBeforeInsertHooks []OrgUserHook
var orgUserAfterInsertHooks []OrgUserHook

var orgUserBeforeUpdateHooks []OrgUserHook
var orgUserAfterUpdateHooks []OrgUserHook

var orgUserBeforeDeleteHooks []OrgUserHook
var orgUserAfterDeleteHooks []OrgUserHook

var orgUserBeforeUpsertHooks []OrgUserHook
var orgUserAfterUpsertHooks []OrgUserHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OrgUser) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OrgUser) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OrgUser) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OrgUser) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OrgUser) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OrgUser) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OrgUser) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OrgUser) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OrgUser) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range orgUserAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOrgUserHook registers your hook function for all future operations.
func AddOrgUserHook(hookPoint boil.HookPoint, orgUserHook OrgUserHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		orgUserAfterSelectHooks = append(orgUserAfterSelectHooks, orgUserHook)
	case boil.BeforeInsertHook:
		orgUserBeforeInsertHooks = append(orgUserBeforeInsertHooks, orgUserHook)
	case boil.AfterInsertHook:
		orgUserAfterInsertHooks = append(orgUserAfterInsertHooks, orgUserHook)
	case boil.BeforeUpdateHook:
		orgUserBeforeUpdateHooks = append(orgUserBeforeUpdateHooks, orgUserHook)
	case boil.AfterUpdateHook:
		orgUserAfterUpdateHooks = append(orgUserAfterUpdateHooks, orgUserHook)
	case boil.BeforeDeleteHook:
		orgUserBeforeDeleteHooks = append(orgUserBeforeDeleteHooks, orgUserHook)
	case boil.AfterDeleteHook:
		orgUserAfterDeleteHooks = append(orgUserAfterDeleteHooks, orgUserHook)
	case boil.BeforeUpsertHook:
		orgUserBeforeUpsertHooks = append(orgUserBeforeUpsertHooks, orgUserHook)
	case boil.AfterUpsertHook:
		orgUserAfterUpsertHooks = append(orgUserAfterUpsertHooks, orgUserHook)
	}
}

// One returns a single orgUser record from the query.
func (q orgUserQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrgUser, error) {
	o := &OrgUser{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: failed to execute a one query for org_users")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OrgUser records from the query.
func (q orgUserQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrgUserSlice, error) {
	var o []*OrgUser

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "dbmodels: failed to assign all query results to OrgUser slice")
	}

	if len(orgUserAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OrgUser records in the query.
func (q orgUserQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to count org_users rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orgUserQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: failed to check if org_users exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *OrgUser) User(mods ...qm.QueryMod) abstractUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"internal_id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return AbstractUsers(queryMods...)
}

// Org pointed to by the foreign key.
func (o *OrgUser) Org(mods ...qm.QueryMod) orgQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"internal_id\" = ?", o.OrgID),
	}

	queryMods = append(queryMods, mods...)

	return Orgs(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orgUserL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrgUser interface{}, mods queries.Applicator) error {
	var slice []*OrgUser
	var object *OrgUser

	if singular {
		var ok bool
		object, ok = maybeOrgUser.(*OrgUser)
		if !ok {
			object = new(OrgUser)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrgUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrgUser))
			}
		}
	} else {
		s, ok := maybeOrgUser.(*[]*OrgUser)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrgUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrgUser))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orgUserR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orgUserR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`abstract_users`),
		qm.WhereIn(`abstract_users.internal_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AbstractUser")
	}

	var resultSlice []*AbstractUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AbstractUser")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for abstract_users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for abstract_users")
	}

	if len(abstractUserAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &abstractUserR{}
		}
		foreign.R.UserOrgUsers = append(foreign.R.UserOrgUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.InternalID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &abstractUserR{}
				}
				foreign.R.UserOrgUsers = append(foreign.R.UserOrgUsers, local)
				break
			}
		}
	}

	return nil
}

// LoadOrg allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (orgUserL) LoadOrg(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOrgUser interface{}, mods queries.Applicator) error {
	var slice []*OrgUser
	var object *OrgUser

	if singular {
		var ok bool
		object, ok = maybeOrgUser.(*OrgUser)
		if !ok {
			object = new(OrgUser)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeOrgUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeOrgUser))
			}
		}
	} else {
		s, ok := maybeOrgUser.(*[]*OrgUser)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeOrgUser)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeOrgUser))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &orgUserR{}
		}
		args = append(args, object.OrgID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &orgUserR{}
			}

			for _, a := range args {
				if a == obj.OrgID {
					continue Outer
				}
			}

			args = append(args, obj.OrgID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orgs`),
		qm.WhereIn(`orgs.internal_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Org")
	}

	var resultSlice []*Org
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Org")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orgs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orgs")
	}

	if len(orgAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Org = foreign
		if foreign.R == nil {
			foreign.R = &orgR{}
		}
		foreign.R.OrgUsers = append(foreign.R.OrgUsers, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.OrgID == foreign.InternalID {
				local.R.Org = foreign
				if foreign.R == nil {
					foreign.R = &orgR{}
				}
				foreign.R.OrgUsers = append(foreign.R.OrgUsers, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the orgUser to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserOrgUsers.
func (o *OrgUser) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AbstractUser) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"org_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, orgUserPrimaryKeyColumns),
	)
	values := []interface{}{related.InternalID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.InternalID
	if o.R == nil {
		o.R = &orgUserR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &abstractUserR{
			UserOrgUsers: OrgUserSlice{o},
		}
	} else {
		related.R.UserOrgUsers = append(related.R.UserOrgUsers, o)
	}

	return nil
}

// SetOrg of the orgUser to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.OrgUsers.
func (o *OrgUser) SetOrg(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Org) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"org_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
		strmangle.WhereClause("\"", "\"", 2, orgUserPrimaryKeyColumns),
	)
	values := []interface{}{related.InternalID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.OrgID = related.InternalID
	if o.R == nil {
		o.R = &orgUserR{
			Org: related,
		}
	} else {
		o.R.Org = related
	}

	if related.R == nil {
		related.R = &orgR{
			OrgUsers: OrgUserSlice{o},
		}
	} else {
		related.R.OrgUsers = append(related.R.OrgUsers, o)
	}

	return nil
}

// OrgUsers retrieves all the records using an executor.
func OrgUsers(mods ...qm.QueryMod) orgUserQuery {
	mods = append(mods, qm.From("\"org_users\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"org_users\".*"})
	}

	return orgUserQuery{q}
}

// FindOrgUser retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrgUser(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*OrgUser, error) {
	orgUserObj := &OrgUser{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"org_users\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orgUserObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "dbmodels: unable to select from org_users")
	}

	if err = orgUserObj.doAfterSelectHooks(ctx, exec); err != nil {
		return orgUserObj, err
	}

	return orgUserObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrgUser) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no org_users provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orgUserColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orgUserInsertCacheMut.RLock()
	cache, cached := orgUserInsertCache[key]
	orgUserInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orgUserAllColumns,
			orgUserColumnsWithDefault,
			orgUserColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orgUserType, orgUserMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orgUserType, orgUserMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"org_users\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"org_users\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to insert into org_users")
	}

	if !cached {
		orgUserInsertCacheMut.Lock()
		orgUserInsertCache[key] = cache
		orgUserInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OrgUser.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrgUser) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	orgUserUpdateCacheMut.RLock()
	cache, cached := orgUserUpdateCache[key]
	orgUserUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orgUserAllColumns,
			orgUserPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("dbmodels: unable to update org_users, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"org_users\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orgUserPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orgUserType, orgUserMapping, append(wl, orgUserPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update org_users row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by update for org_users")
	}

	if !cached {
		orgUserUpdateCacheMut.Lock()
		orgUserUpdateCache[key] = cache
		orgUserUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q orgUserQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all for org_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected for org_users")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrgUserSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("dbmodels: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orgUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"org_users\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orgUserPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to update all in orgUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to retrieve rows affected all in update all orgUser")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrgUser) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("dbmodels: no org_users provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(orgUserColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orgUserUpsertCacheMut.RLock()
	cache, cached := orgUserUpsertCache[key]
	orgUserUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orgUserAllColumns,
			orgUserColumnsWithDefault,
			orgUserColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			orgUserAllColumns,
			orgUserPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("dbmodels: unable to upsert org_users, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orgUserPrimaryKeyColumns))
			copy(conflict, orgUserPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"org_users\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orgUserType, orgUserMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orgUserType, orgUserMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to upsert org_users")
	}

	if !cached {
		orgUserUpsertCacheMut.Lock()
		orgUserUpsertCache[key] = cache
		orgUserUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OrgUser record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrgUser) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("dbmodels: no OrgUser provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orgUserPrimaryKeyMapping)
	sql := "DELETE FROM \"org_users\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete from org_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by delete for org_users")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orgUserQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("dbmodels: no orgUserQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from org_users")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for org_users")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrgUserSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(orgUserBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orgUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"org_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orgUserPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: unable to delete all from orgUser slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "dbmodels: failed to get rows affected by deleteall for org_users")
	}

	if len(orgUserAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrgUser) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrgUser(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrgUserSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrgUserSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orgUserPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"org_users\".* FROM \"org_users\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orgUserPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "dbmodels: unable to reload all in OrgUserSlice")
	}

	*o = slice

	return nil
}

// OrgUserExists checks if the OrgUser row exists.
func OrgUserExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"org_users\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "dbmodels: unable to check if org_users exists")
	}

	return exists, nil
}

// Exists checks if the OrgUser row exists.
func (o *OrgUser) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return OrgUserExists(ctx, exec, o.ID)
}