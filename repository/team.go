// Code generated by SQLBoiler 4.10.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package repository

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

// Team is an object representing the database table.
type Team struct {
	ID          string `boil:"id" json:"id" toml:"id" yaml:"id"`
	DisplayName string `boil:"display_name" json:"display_name" toml:"display_name" yaml:"display_name"`

	R *teamR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L teamL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TeamColumns = struct {
	ID          string
	DisplayName string
}{
	ID:          "id",
	DisplayName: "display_name",
}

var TeamTableColumns = struct {
	ID          string
	DisplayName string
}{
	ID:          "team.id",
	DisplayName: "team.display_name",
}

// Generated where

var TeamWhere = struct {
	ID          whereHelperstring
	DisplayName whereHelperstring
}{
	ID:          whereHelperstring{field: "\"leaderseek\".\"team\".\"id\""},
	DisplayName: whereHelperstring{field: "\"leaderseek\".\"team\".\"display_name\""},
}

// TeamRels is where relationship names are stored.
var TeamRels = struct {
	TeamMembers string
}{
	TeamMembers: "TeamMembers",
}

// teamR is where relationships are stored.
type teamR struct {
	TeamMembers TeamMemberSlice `boil:"TeamMembers" json:"TeamMembers" toml:"TeamMembers" yaml:"TeamMembers"`
}

// NewStruct creates a new relationship struct
func (*teamR) NewStruct() *teamR {
	return &teamR{}
}

// teamL is where Load methods for each relationship are stored.
type teamL struct{}

var (
	teamAllColumns            = []string{"id", "display_name"}
	teamColumnsWithoutDefault = []string{"id", "display_name"}
	teamColumnsWithDefault    = []string{}
	teamPrimaryKeyColumns     = []string{"id"}
	teamGeneratedColumns      = []string{}
)

type (
	// TeamSlice is an alias for a slice of pointers to Team.
	// This should almost always be used instead of []Team.
	TeamSlice []*Team
	// TeamHook is the signature for custom Team hook methods
	TeamHook func(context.Context, boil.ContextExecutor, *Team) error

	teamQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	teamType                 = reflect.TypeOf(&Team{})
	teamMapping              = queries.MakeStructMapping(teamType)
	teamPrimaryKeyMapping, _ = queries.BindMapping(teamType, teamMapping, teamPrimaryKeyColumns)
	teamInsertCacheMut       sync.RWMutex
	teamInsertCache          = make(map[string]insertCache)
	teamUpdateCacheMut       sync.RWMutex
	teamUpdateCache          = make(map[string]updateCache)
	teamUpsertCacheMut       sync.RWMutex
	teamUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var teamAfterSelectHooks []TeamHook

var teamBeforeInsertHooks []TeamHook
var teamAfterInsertHooks []TeamHook

var teamBeforeUpdateHooks []TeamHook
var teamAfterUpdateHooks []TeamHook

var teamBeforeDeleteHooks []TeamHook
var teamAfterDeleteHooks []TeamHook

var teamBeforeUpsertHooks []TeamHook
var teamAfterUpsertHooks []TeamHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Team) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Team) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Team) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Team) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Team) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Team) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Team) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Team) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Team) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range teamAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTeamHook registers your hook function for all future operations.
func AddTeamHook(hookPoint boil.HookPoint, teamHook TeamHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		teamAfterSelectHooks = append(teamAfterSelectHooks, teamHook)
	case boil.BeforeInsertHook:
		teamBeforeInsertHooks = append(teamBeforeInsertHooks, teamHook)
	case boil.AfterInsertHook:
		teamAfterInsertHooks = append(teamAfterInsertHooks, teamHook)
	case boil.BeforeUpdateHook:
		teamBeforeUpdateHooks = append(teamBeforeUpdateHooks, teamHook)
	case boil.AfterUpdateHook:
		teamAfterUpdateHooks = append(teamAfterUpdateHooks, teamHook)
	case boil.BeforeDeleteHook:
		teamBeforeDeleteHooks = append(teamBeforeDeleteHooks, teamHook)
	case boil.AfterDeleteHook:
		teamAfterDeleteHooks = append(teamAfterDeleteHooks, teamHook)
	case boil.BeforeUpsertHook:
		teamBeforeUpsertHooks = append(teamBeforeUpsertHooks, teamHook)
	case boil.AfterUpsertHook:
		teamAfterUpsertHooks = append(teamAfterUpsertHooks, teamHook)
	}
}

// One returns a single team record from the query.
func (q teamQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Team, error) {
	o := &Team{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for team")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Team records from the query.
func (q teamQuery) All(ctx context.Context, exec boil.ContextExecutor) (TeamSlice, error) {
	var o []*Team

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to Team slice")
	}

	if len(teamAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Team records in the query.
func (q teamQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count team rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q teamQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if team exists")
	}

	return count > 0, nil
}

// TeamMembers retrieves all the team_member's TeamMembers with an executor.
func (o *Team) TeamMembers(mods ...qm.QueryMod) teamMemberQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"leaderseek\".\"team_member\".\"team_id\"=?", o.ID),
	)

	return TeamMembers(queryMods...)
}

// LoadTeamMembers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (teamL) LoadTeamMembers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTeam interface{}, mods queries.Applicator) error {
	var slice []*Team
	var object *Team

	if singular {
		object = maybeTeam.(*Team)
	} else {
		slice = *maybeTeam.(*[]*Team)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &teamR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &teamR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`leaderseek.team_member`),
		qm.WhereIn(`leaderseek.team_member.team_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load team_member")
	}

	var resultSlice []*TeamMember
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice team_member")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on team_member")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for team_member")
	}

	if len(teamMemberAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TeamMembers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &teamMemberR{}
			}
			foreign.R.Team = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.TeamID) {
				local.R.TeamMembers = append(local.R.TeamMembers, foreign)
				if foreign.R == nil {
					foreign.R = &teamMemberR{}
				}
				foreign.R.Team = local
				break
			}
		}
	}

	return nil
}

// AddTeamMembers adds the given related objects to the existing relationships
// of the team, optionally inserting them as new records.
// Appends related to o.R.TeamMembers.
// Sets related.R.Team appropriately.
func (o *Team) AddTeamMembers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TeamMember) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.TeamID, o.ID)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"leaderseek\".\"team_member\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"team_id"}),
				strmangle.WhereClause("\"", "\"", 2, teamMemberPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.TeamID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &teamR{
			TeamMembers: related,
		}
	} else {
		o.R.TeamMembers = append(o.R.TeamMembers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &teamMemberR{
				Team: o,
			}
		} else {
			rel.R.Team = o
		}
	}
	return nil
}

// SetTeamMembers removes all previously related items of the
// team replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.Team's TeamMembers accordingly.
// Replaces o.R.TeamMembers with related.
// Sets related.R.Team's TeamMembers accordingly.
func (o *Team) SetTeamMembers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TeamMember) error {
	query := "update \"leaderseek\".\"team_member\" set \"team_id\" = null where \"team_id\" = $1"
	values := []interface{}{o.ID}
	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, query)
		fmt.Fprintln(writer, values)
	}
	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.TeamMembers {
			queries.SetScanner(&rel.TeamID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.Team = nil
		}
		o.R.TeamMembers = nil
	}

	return o.AddTeamMembers(ctx, exec, insert, related...)
}

// RemoveTeamMembers relationships from objects passed in.
// Removes related items from R.TeamMembers (uses pointer comparison, removal does not keep order)
// Sets related.R.Team.
func (o *Team) RemoveTeamMembers(ctx context.Context, exec boil.ContextExecutor, related ...*TeamMember) error {
	if len(related) == 0 {
		return nil
	}

	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.TeamID, nil)
		if rel.R != nil {
			rel.R.Team = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("team_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.TeamMembers {
			if rel != ri {
				continue
			}

			ln := len(o.R.TeamMembers)
			if ln > 1 && i < ln-1 {
				o.R.TeamMembers[i] = o.R.TeamMembers[ln-1]
			}
			o.R.TeamMembers = o.R.TeamMembers[:ln-1]
			break
		}
	}

	return nil
}

// Teams retrieves all the records using an executor.
func Teams(mods ...qm.QueryMod) teamQuery {
	mods = append(mods, qm.From("\"leaderseek\".\"team\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"leaderseek\".\"team\".*"})
	}

	return teamQuery{q}
}

// FindTeam retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTeam(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Team, error) {
	teamObj := &Team{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"leaderseek\".\"team\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, teamObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: unable to select from team")
	}

	if err = teamObj.doAfterSelectHooks(ctx, exec); err != nil {
		return teamObj, err
	}

	return teamObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Team) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no team provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(teamColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	teamInsertCacheMut.RLock()
	cache, cached := teamInsertCache[key]
	teamInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			teamAllColumns,
			teamColumnsWithDefault,
			teamColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(teamType, teamMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"leaderseek\".\"team\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"leaderseek\".\"team\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "repository: unable to insert into team")
	}

	if !cached {
		teamInsertCacheMut.Lock()
		teamInsertCache[key] = cache
		teamInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Team.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Team) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	teamUpdateCacheMut.RLock()
	cache, cached := teamUpdateCache[key]
	teamUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			teamAllColumns,
			teamPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("repository: unable to update team, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"leaderseek\".\"team\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, teamPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, append(wl, teamPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "repository: unable to update team row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by update for team")
	}

	if !cached {
		teamUpdateCacheMut.Lock()
		teamUpdateCache[key] = cache
		teamUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q teamQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all for team")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected for team")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TeamSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("repository: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"leaderseek\".\"team\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, teamPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to update all in team slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to retrieve rows affected all in update all team")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Team) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("repository: no team provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(teamColumnsWithDefault, o)

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

	teamUpsertCacheMut.RLock()
	cache, cached := teamUpsertCache[key]
	teamUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			teamAllColumns,
			teamColumnsWithDefault,
			teamColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			teamAllColumns,
			teamPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("repository: unable to upsert team, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(teamPrimaryKeyColumns))
			copy(conflict, teamPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"leaderseek\".\"team\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(teamType, teamMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(teamType, teamMapping, ret)
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
		return errors.Wrap(err, "repository: unable to upsert team")
	}

	if !cached {
		teamUpsertCacheMut.Lock()
		teamUpsertCache[key] = cache
		teamUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Team record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Team) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("repository: no Team provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), teamPrimaryKeyMapping)
	sql := "DELETE FROM \"leaderseek\".\"team\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete from team")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by delete for team")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q teamQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("repository: no teamQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from team")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for team")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TeamSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(teamBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"leaderseek\".\"team\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "repository: unable to delete all from team slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to get rows affected by deleteall for team")
	}

	if len(teamAfterDeleteHooks) != 0 {
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
func (o *Team) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTeam(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TeamSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TeamSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), teamPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"leaderseek\".\"team\".* FROM \"leaderseek\".\"team\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, teamPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "repository: unable to reload all in TeamSlice")
	}

	*o = slice

	return nil
}

// TeamExists checks if the Team row exists.
func TeamExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"leaderseek\".\"team\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "repository: unable to check if team exists")
	}

	return exists, nil
}
