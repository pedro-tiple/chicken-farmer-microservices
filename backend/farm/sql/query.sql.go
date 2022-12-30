// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: query.sql

package sql

import (
	"context"

	"github.com/google/uuid"
)

const decrementBarnFeed = `-- name: DecrementBarnFeed :exec
UPDATE barns
SET feed = feed - $2
WHERE id = $1
`

type DecrementBarnFeedParams struct {
	ID   uuid.UUID
	Feed int64
}

func (q *Queries) DecrementBarnFeed(ctx context.Context, arg DecrementBarnFeedParams) error {
	_, err := q.db.ExecContext(ctx, decrementBarnFeed, arg.ID, arg.Feed)
	return err
}

const getBarnsOfFarm = `-- name: GetBarnsOfFarm :many
SELECT barns.id    as id,
    farm_id,
    feed,
    has_auto_feeder,
    farms.owner_id as owner_id
FROM barns
         INNER JOIN farms ON farm_id = farms.id
WHERE barns.id = $1
ORDER BY barns.created_at
`

type GetBarnsOfFarmRow struct {
	ID            uuid.UUID
	FarmID        uuid.UUID
	Feed          int64
	HasAutoFeeder bool
	OwnerID       uuid.UUID
}

func (q *Queries) GetBarnsOfFarm(ctx context.Context, id uuid.UUID) ([]GetBarnsOfFarmRow, error) {
	rows, err := q.db.QueryContext(ctx, getBarnsOfFarm, id)
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

const incrementChickenGoldEggLayCount = `-- name: IncrementChickenGoldEggLayCount :exec
UPDATE chickens
SET gold_eggs_laid = gold_eggs_laid + 1
WHERE id = $1
`

func (q *Queries) IncrementChickenGoldEggLayCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, incrementChickenGoldEggLayCount, id)
	return err
}

const incrementChickenNormalEggLayCount = `-- name: IncrementChickenNormalEggLayCount :exec
UPDATE chickens
SET normal_eggs_laid = normal_eggs_laid + 1
WHERE id = $1
`

func (q *Queries) IncrementChickenNormalEggLayCount(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, incrementChickenNormalEggLayCount, id)
	return err
}

const insertBarn = `-- name: InsertBarn :execlastid
INSERT INTO barns (
    farm_id, feed, has_auto_feeder
)
VALUES (
    $1, $2, $3
)
`

type InsertBarnParams struct {
	FarmID        uuid.UUID
	Feed          int64
	HasAutoFeeder bool
}

func (q *Queries) InsertBarn(ctx context.Context, arg InsertBarnParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, insertBarn, arg.FarmID, arg.Feed, arg.HasAutoFeeder)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const insertChicken = `-- name: InsertChicken :execlastid
INSERT INTO chickens (
    id, date_of_birth, resting_until, normal_eggs_laid, gold_eggs_laid,
    gold_egg_chance, barn_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
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

func (q *Queries) InsertChicken(ctx context.Context, arg InsertChickenParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, insertChicken,
		arg.ID,
		arg.DateOfBirth,
		arg.RestingUntil,
		arg.NormalEggsLaid,
		arg.GoldEggsLaid,
		arg.GoldEggChance,
		arg.BarnID,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}
