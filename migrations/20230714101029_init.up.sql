CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;

CREATE TABLE LEAGUE
(
    id         uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name       text                                  NOT NULL,
    start_date date,
    end_date   date,
    website    text
);

CREATE TABLE PARTICIPANT
(
    id           uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name         text                                  NOT NULL,
    contact_info text,
    wins         int  DEFAULT 0,
    losses       int  DEFAULT 0,
    rank         int  DEFAULT 0
);

CREATE TABLE MATCH
(
    id                 uuid DEFAULT public.gen_random_uuid() NOT NULL,
    group_id           uuid,
    participant_one_id uuid,
    participant_two_id uuid,
    replay_url         text
);

CREATE TABLE MATCH_VETO
(
    id             uuid DEFAULT public.gen_random_uuid() NOT NULL,
    match_id       uuid,
    participant_id uuid,
    game_map_id    uuid
);

CREATE TABLE GAME
(
    id                uuid DEFAULT public.gen_random_uuid() NOT NULL,
    match_id          uuid,
    round_number      int,
    winning_player_id uuid,
    game_map_id       uuid,
    game_date         timestamp
);

CREATE TABLE "group"
(
    id             uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name           text,
    league_id      uuid,
    admin          text,
    vetos          int,
    maps_per_match int,
    status         text
);

CREATE TABLE GAME_MAP
(
    id   uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name text
);

CREATE TABLE PARTICIPANT_GROUP
(
    id             uuid DEFAULT public.gen_random_uuid() NOT NULL,
    participant_id uuid,
    group_id       uuid
)