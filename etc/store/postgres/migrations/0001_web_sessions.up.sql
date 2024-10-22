create schema auth;

create table auth.web_sessions
(
    id              text                         primary key unique,
    user_id         text                         not null,
    created_at      timestamp with time zone     not null default CURRENT_TIMESTAMP,
    fingerprint     text                         not null
);
