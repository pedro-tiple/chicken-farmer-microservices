// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sql

import (
	"context"

	"github.com/google/uuid"
)

const decrementBarnFeedGreaterEqualThan = `-- name: DecrementBarnFeedGreaterEqualThan :execrows
UPDATE barns
SET feed = feed - $2
WHERE id = $1
  AND feed >= $2
`

type DecrementBarnFeedGreaterEqualThanParams struct {
	ID   uuid.UUID
	Feed int64
}

func (q *Queries) DecrementBarnFeedGreaterEqualThan(ctx context.Context, arg DecrementBarnFeedGreaterEqualThanParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, decrementBarnFeedGreaterEqualThan, arg.ID, arg.Feed)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const deleteChicken = `-- name: DeleteChicken :exec
DELETE FROM chickens
WHERE ID = $1
`

func (q *Queries) DeleteChicken(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteChicken, id)
	return err
}

const getBarn = `-- name: GetBarn :one
SELECT barns.id    as id,
    farm_id,
    feed,
    has_auto_feeder,
    farms.owner_id as owner_id
FROM barns
         INNER JOIN farms ON farm_id = farms.id
WHERE barns.id = $1
`

type GetBarnRow struct {
	ID            uuid.UUID
	FarmID        uuid.UUID
	Feed          int64
	HasAutoFeeder bool
	OwnerID       uuid.UUID
}

func (q *Queries) GetBarn(ctx context.Context, id uuid.UUID) (GetBarnRow, error) {
	row := q.db.QueryRowContext(ctx, getBarn, id)
	var i GetBarnRow
	err := row.Scan(
		&i.ID,
		&i.FarmID,
		&i.Feed,
		&i.HasAutoFeeder,
		&i.OwnerID,
	)
	return i, err
}

const getBarnsOfFarm = `-- name: GetBarnsOfFarm :many
SELECT barns.id    as id,
    farm_id,
    feed,
    has_auto_feeder,
    farms.owner_id as owner_id
FROM barns
         INNER JOIN farms ON farm_id = farms.id
WHERE barns.farm_id = $1
ORDER BY barns.created_at
`

type GetBarnsOfFarmRow struct {
	ID            uuid.UUID
	FarmID        uuid.UUID
	Feed          int64
	HasAutoFeeder bool
	OwnerID       uuid.UUID
}

func (q *Queries) GetBarnsOfFarm(ctx context.Context, farmID uuid.UUID) ([]GetBarnsOfFarmRow, error) {
	rows, err := q.db.QueryContext(ctx, getBarnsOfFarm, farmID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetBarnsOfFarmRow
	for rows.Next() {
		var i GetBarnsOfFarmRow
		if err := rows.Scan(
			&i.ID,
			&i.FarmID,
			&i.Feed,
			&i.HasAutoFeeder,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getChicken = `-- name: GetChicken :one
SELECT chickens.id as id,
    date_of_birth,
    resting_until,
    normal_eggs_laid,
    gold_eggs_laid,
    gold_egg_chance,
    barns.id       as barn_id,
    farms.owner_id as owner_id
FROM chickens
         INNER JOIN barns ON barn_id = barns.id
         INNER JOIN farms ON farm_id = farms.id
WHERE chickens.id = $1
`

type GetChickenRow struct {
	ID             uuid.UUID
	DateOfBirth    int64
	RestingUntil   int64
	NormalEggsLaid int64
	GoldEggsLaid   int64
	GoldEggChance  int64
	BarnID         uuid.UUID
	OwnerID        uuid.UUID
}

func (q *Queries) GetChicken(ctx context.Context, id uuid.UUID) (GetChickenRow, error) {
	row := q.db.QueryRowContext(ctx, getChicken, id)
	var i GetChickenRow
	err := row.Scan(
		&i.ID,
		&i.DateOfBirth,
		&i.RestingUntil,
		&i.NormalEggsLaid,
		&i.GoldEggsLaid,
		&i.GoldEggChance,
		&i.BarnID,
		&i.OwnerID,
	)
	return i, err
}

const getChickensOfBarn = `-- name: GetChickensOfBarn :many
SELECT chickens.id as id,
    chickens.date_of_birth,
    chickens.resting_until,
    chickens.normal_eggs_laid,
    chickens.gold_eggs_laid,
    chickens.gold_egg_chance,
    barns.id       as barn_id,
    farms.owner_id as owner_id
FROM chickens
         INNER JOIN barns ON barn_id = barns.id
         INNER JOIN farms ON farm_id = farms.id
WHERE barns.id = $1
ORDER by chickens.created_at
`

type GetChickensOfBarnRow struct {
	ID             uuid.UUID
	DateOfBirth    int64
	RestingUntil   int64
	NormalEggsLaid int64
	GoldEggsLaid   int64
	GoldEggChance  int64
	BarnID         uuid.UUID
	OwnerID        uuid.UUID
}

func (q *Queries) GetChickensOfBarn(ctx context.Context, id uuid.UUID) ([]GetChickensOfBarnRow, error) {
	rows, err := q.db.QueryContext(ctx, getChickensOfBarn, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetChickensOfBarnRow
	for rows.Next() {
		var i GetChickensOfBarnRow
		if err := rows.Scan(
			&i.ID,
			&i.DateOfBirth,
			&i.RestingUntil,
			&i.NormalEggsLaid,
			&i.GoldEggsLaid,
			&i.GoldEggChance,
			&i.BarnID,
			&i.OwnerID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getFarm = `-- name: GetFarm :one
SELECT id,
    owner_id,
    name
FROM farms
WHERE id = $1
LIMIT 1
`

type GetFarmRow struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
}

func (q *Queries) GetFarm(ctx context.Context, id uuid.UUID) (GetFarmRow, error) {
	row := q.db.QueryRowContext(ctx, getFarm, id)
	var i GetFarmRow
	err := row.Scan(&i.ID, &i.OwnerID, &i.Name)
	return i, err
}

const incrementBarnFeed = `-- name: IncrementBarnFeed :exec
UPDATE barns
SET feed = feed + $2
WHERE id = $1
`

type IncrementBarnFeedParams struct {
	ID   uuid.UUID
	Feed int64
}

func (q *Queries) IncrementBarnFeed(ctx context.Context, arg IncrementBarnFeedParams) error {
	_, err := q.db.ExecContext(ctx, incrementBarnFeed, arg.ID, arg.Feed)
	return err
}

const incrementChickenEggLayCount = `-- name: IncrementChickenEggLayCount :exec
UPDATE chickens
SET normal_eggs_laid = normal_eggs_laid + $2, gold_eggs_laid = gold_eggs_laid + $3
WHERE id = $1
`

type IncrementChickenEggLayCountParams struct {
	ID             uuid.UUID
	NormalEggsLaid int64
	GoldEggsLaid   int64
}

func (q *Queries) IncrementChickenEggLayCount(ctx context.Context, arg IncrementChickenEggLayCountParams) error {
	_, err := q.db.ExecContext(ctx, incrementChickenEggLayCount, arg.ID, arg.NormalEggsLaid, arg.GoldEggsLaid)
	return err
}

const insertBarn = `-- name: InsertBarn :one
INSERT INTO barns (
    farm_id, feed, has_auto_feeder
)
VALUES (
    $1, $2, $3
)
RETURNING id, farm_id, feed, has_auto_feeder, created_at, updated_at
`

type InsertBarnParams struct {
	FarmID        uuid.UUID
	Feed          int64
	HasAutoFeeder bool
}

func (q *Queries) InsertBarn(ctx context.Context, arg InsertBarnParams) (Barn, error) {
	row := q.db.QueryRowContext(ctx, insertBarn, arg.FarmID, arg.Feed, arg.HasAutoFeeder)
	var i Barn
	err := row.Scan(
		&i.ID,
		&i.FarmID,
		&i.Feed,
		&i.HasAutoFeeder,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertChicken = `-- name: InsertChicken :one
INSERT INTO chickens (
    id, date_of_birth, resting_until, normal_eggs_laid, gold_eggs_laid,
    gold_egg_chance, barn_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING id, barn_id, date_of_birth, resting_until, normal_eggs_laid, gold_eggs_laid, gold_egg_chance, created_at, updated_at
`

type InsertChickenParams struct {
	ID             uuid.UUID
	DateOfBirth    int64
	RestingUntil   int64
	NormalEggsLaid int64
	GoldEggsLaid   int64
	GoldEggChance  int64
	BarnID         uuid.UUID
}

func (q *Queries) InsertChicken(ctx context.Context, arg InsertChickenParams) (Chicken, error) {
	row := q.db.QueryRowContext(ctx, insertChicken,
		arg.ID,
		arg.DateOfBirth,
		arg.RestingUntil,
		arg.NormalEggsLaid,
		arg.GoldEggsLaid,
		arg.GoldEggChance,
		arg.BarnID,
	)
	var i Chicken
	err := row.Scan(
		&i.ID,
		&i.BarnID,
		&i.DateOfBirth,
		&i.RestingUntil,
		&i.NormalEggsLaid,
		&i.GoldEggsLaid,
		&i.GoldEggChance,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const insertFarm = `-- name: InsertFarm :one
INSERT INTO farms (
    id, owner_id, name
)
VALUES (
    $1, $2, $3
)
RETURNING id, owner_id, name, created_at, updated_at
`

type InsertFarmParams struct {
	ID      uuid.UUID
	OwnerID uuid.UUID
	Name    string
}

func (q *Queries) InsertFarm(ctx context.Context, arg InsertFarmParams) (Farm, error) {
	row := q.db.QueryRowContext(ctx, insertFarm, arg.ID, arg.OwnerID, arg.Name)
	var i Farm
	err := row.Scan(
		&i.ID,
		&i.OwnerID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateChickenRestingUntil = `-- name: UpdateChickenRestingUntil :exec
UPDATE chickens
SET resting_until = $2
WHERE id = $1
`

type UpdateChickenRestingUntilParams struct {
	ID           uuid.UUID
	RestingUntil int64
}

func (q *Queries) UpdateChickenRestingUntil(ctx context.Context, arg UpdateChickenRestingUntilParams) error {
	_, err := q.db.ExecContext(ctx, updateChickenRestingUntil, arg.ID, arg.RestingUntil)
	return err
}
