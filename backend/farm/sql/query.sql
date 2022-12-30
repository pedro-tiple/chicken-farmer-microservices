-- name: GetFarm :one
SELECT id,
    owner_id,
    name
FROM farms
WHERE id = $1
LIMIT 1;

-- name: GetBarnsOfFarm :many
SELECT barns.id    as id,
    farm_id,
    feed,
    has_auto_feeder,
    farms.owner_id as owner_id
FROM barns
         INNER JOIN farms ON farm_id = farms.id
WHERE barns.id = $1
ORDER BY barns.created_at;

-- name: GetChickensOfBarn :many
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
ORDER by chickens.created_at;

-- name: GetChicken :one
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
WHERE chickens.id = $1;

-- name: InsertChicken :execlastid
INSERT INTO chickens (
    id, date_of_birth, resting_until, normal_eggs_laid, gold_eggs_laid,
    gold_egg_chance, barn_id
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
);

-- name: InsertBarn :execlastid
INSERT INTO barns (
    farm_id, feed, has_auto_feeder
)
VALUES (
    $1, $2, $3
);

-- name: IncrementBarnFeed :exec
UPDATE barns
SET feed = feed + $2
WHERE id = $1;

-- name: DecrementBarnFeed :exec
UPDATE barns
SET feed = feed - $2
WHERE id = $1;


-- name: IncrementChickenNormalEggLayCount :exec
UPDATE chickens
SET normal_eggs_laid = normal_eggs_laid + 1
WHERE id = $1;

-- name: IncrementChickenGoldEggLayCount :exec
UPDATE chickens
SET gold_eggs_laid = gold_eggs_laid + 1
WHERE id = $1;
