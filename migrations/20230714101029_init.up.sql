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

CREATE TABLE TEAM
(
    id          uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name        text                                  NOT NULL,
    members_ids uuid[]                                NOT NULL
);

CREATE TABLE MATCH
(
    id                 uuid DEFAULT public.gen_random_uuid() NOT NULL,
    participant_one_id uuid,
    participant_two_id uuid,
    match_date         timestamp,
    winner_id          uuid
);

CREATE TABLE ROUND
(
    round_id     uuid,
    match_ids    uuid[],
    round_number text
);

CREATE TABLE BRACKET
(
    id        uuid DEFAULT public.gen_random_uuid() NOT NULL,
    round_ids uuid[]
);

CREATE TABLE GAME_MAP
(
    id   uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name text
);

CREATE TABLE HERO_UNIT
(
    id   uuid DEFAULT public.gen_random_uuid() NOT NULL,
    name text
)