-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists users
(
    id              text default uuid_generate_v4() not null,
    username        text,
    hasedpassword   text,
    created_at      timestamp with time zone,
    updated_at      timestamp with time zone,
    constraint users_pkey
        primary key (id)
);

create table if not exists cards
(
    id          text default uuid_generate_v4() not null,
    name        text,
    user_id     text,
    number      text,
    holder      text,
    expire      text,
    cvc         text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    constraint cards_pkey
        primary key (id),
    constraint cards_users_id_fkey
        foreign key (user_id) references users
);

create table if not exists accounts
(
    id          text default uuid_generate_v4() not null,
    name        text,
    user_id     text,
    login       text,
    password    text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    constraint accounts_pkey
        primary key (id),
    constraint accounts_users_id_fkey
        foreign key (user_id) references users
);

create table if not exists notes
(
    id          text default uuid_generate_v4() not null,
    name        text,
    user_id     text,
    text        text,
    created_at  timestamp with time zone,
    updated_at  timestamp with time zone,
    constraint notes_pkey
        primary key (id),
    constraint notes_users_id_fkey
        foreign key (user_id) references users
);

-- +migrate Down
drop table if exists notes;
drop table if exists accounts;
drop table if exists cards;
drop table if exists users;