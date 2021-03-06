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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// LeaderSupporterCount is an object representing the database table.
type LeaderSupporterCount struct {
	CandidateID    null.String `boil:"candidate_id" json:"candidate_id,omitempty" toml:"candidate_id" yaml:"candidate_id,omitempty"`
	SupporterCount null.Int64  `boil:"supporter_count" json:"supporter_count,omitempty" toml:"supporter_count" yaml:"supporter_count,omitempty"`
}

var LeaderSupporterCountColumns = struct {
	CandidateID    string
	SupporterCount string
}{
	CandidateID:    "candidate_id",
	SupporterCount: "supporter_count",
}

var LeaderSupporterCountTableColumns = struct {
	CandidateID    string
	SupporterCount string
}{
	CandidateID:    "leader_supporter_count.candidate_id",
	SupporterCount: "leader_supporter_count.supporter_count",
}

// Generated where

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var LeaderSupporterCountWhere = struct {
	CandidateID    whereHelpernull_String
	SupporterCount whereHelpernull_Int64
}{
	CandidateID:    whereHelpernull_String{field: "\"leaderseek\".\"leader_supporter_count\".\"candidate_id\""},
	SupporterCount: whereHelpernull_Int64{field: "\"leaderseek\".\"leader_supporter_count\".\"supporter_count\""},
}

var (
	leaderSupporterCountAllColumns            = []string{"candidate_id", "supporter_count"}
	leaderSupporterCountColumnsWithoutDefault = []string{}
	leaderSupporterCountColumnsWithDefault    = []string{"candidate_id", "supporter_count"}
	leaderSupporterCountPrimaryKeyColumns     = []string{}
	leaderSupporterCountGeneratedColumns      = []string{}
)

type (
	// LeaderSupporterCountSlice is an alias for a slice of pointers to LeaderSupporterCount.
	// This should almost always be used instead of []LeaderSupporterCount.
	LeaderSupporterCountSlice []*LeaderSupporterCount
	// LeaderSupporterCountHook is the signature for custom LeaderSupporterCount hook methods
	LeaderSupporterCountHook func(context.Context, boil.ContextExecutor, *LeaderSupporterCount) error

	leaderSupporterCountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	leaderSupporterCountType           = reflect.TypeOf(&LeaderSupporterCount{})
	leaderSupporterCountMapping        = queries.MakeStructMapping(leaderSupporterCountType)
	leaderSupporterCountInsertCacheMut sync.RWMutex
	leaderSupporterCountInsertCache    = make(map[string]insertCache)
	leaderSupporterCountUpdateCacheMut sync.RWMutex
	leaderSupporterCountUpdateCache    = make(map[string]updateCache)
	leaderSupporterCountUpsertCacheMut sync.RWMutex
	leaderSupporterCountUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
	// These are used in some views
	_ = fmt.Sprintln("")
	_ = reflect.Int
	_ = strings.Builder{}
	_ = sync.Mutex{}
	_ = strmangle.Plural("")
	_ = strconv.IntSize
)

var leaderSupporterCountAfterSelectHooks []LeaderSupporterCountHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *LeaderSupporterCount) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range leaderSupporterCountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddLeaderSupporterCountHook registers your hook function for all future operations.
func AddLeaderSupporterCountHook(hookPoint boil.HookPoint, leaderSupporterCountHook LeaderSupporterCountHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		leaderSupporterCountAfterSelectHooks = append(leaderSupporterCountAfterSelectHooks, leaderSupporterCountHook)
	}
}

// One returns a single leaderSupporterCount record from the query.
func (q leaderSupporterCountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*LeaderSupporterCount, error) {
	o := &LeaderSupporterCount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "repository: failed to execute a one query for leader_supporter_count")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all LeaderSupporterCount records from the query.
func (q leaderSupporterCountQuery) All(ctx context.Context, exec boil.ContextExecutor) (LeaderSupporterCountSlice, error) {
	var o []*LeaderSupporterCount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "repository: failed to assign all query results to LeaderSupporterCount slice")
	}

	if len(leaderSupporterCountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all LeaderSupporterCount records in the query.
func (q leaderSupporterCountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "repository: failed to count leader_supporter_count rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q leaderSupporterCountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "repository: failed to check if leader_supporter_count exists")
	}

	return count > 0, nil
}

// LeaderSupporterCounts retrieves all the records using an executor.
func LeaderSupporterCounts(mods ...qm.QueryMod) leaderSupporterCountQuery {
	mods = append(mods, qm.From("\"leaderseek\".\"leader_supporter_count\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"leaderseek\".\"leader_supporter_count\".*"})
	}

	return leaderSupporterCountQuery{q}
}
