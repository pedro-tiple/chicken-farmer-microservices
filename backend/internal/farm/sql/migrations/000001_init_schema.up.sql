CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE farms (
    id         uuid        DEFAULT uuid_generate_v4(),

    owner_id   uuid NOT NULL,
    name       text NOT NULL,

    created_at timestamptz DEFAULT now(),
    updated_at timestamptz,

    PRIMARY KEY (id)
);

INSERT INTO farms (
    id, owner_id, name
)
values (
    '93020a42-c32a-4b2c-a4b9-779f82841b11',
    '65e4d8ff-8766-48a7-bfcd-7160d149a319',
    'Default Farm'
);

CREATE TABLE barns (
    id              uuid        DEFAULT uuid_generate_v4(),
    farm_id         uuid    NOT NULL,

    feed            bigint  NOT NULL,
    has_auto_feeder boolean NOT NULL,

    created_at      timestamptz DEFAULT now(),
    updated_at      timestamptz,

    PRIMARY KEY (id),
    FOREIGN KEY (farm_id) REFERENCES farms (id)

);

CREATE TABLE chickens (
    id               uuid        DEFAULT uuid_generate_v4(),
    barn_id          uuid   NOT NULL,

    date_of_birth    bigint NOT NULL,
    resting_until    bigint NOT NULL,
    normal_eggs_laid bigint NOT NULL,
    gold_eggs_laid   bigint NOT NULL,
    gold_egg_chance  bigint NOT NULL,

    created_at       timestamptz DEFAULT now(),
    updated_at       timestamptz,

    PRIMARY KEY (id),
    FOREIGN KEY (barn_id) REFERENCES barns (id)
);
