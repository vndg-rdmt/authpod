-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create schema auth;

create table if not exists auth.users
(
    id              bigserial                   primary key,
    created_at      timestamp with time zone    not null default current_timestamp,
    login           text                        not null unique,
    password_hash   text                        not null
);

create table if not exists auth.tokens
(
    token           uuid                        not null default gen_random_uuid(),
    user_id         bigint                      references auth.users(id),
    created_at      timestamp with time zone    not null default current_timestamp,
    expires_at      timestamp with time zone    not null
);

create table if not exists auth.websessions
(
    session_id      uuid                        not null default gen_random_uuid(),
    user_id         bigint                      references auth.users(id),
    created_at      timestamp with time zone    not null default current_timestamp,
    expires_at      timestamp with time zone    not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists auth.websessions;
drop table if exists auth.tokens;
drop table if exists auth.users;
drop schema auth;
-- +goose StatementEnd
