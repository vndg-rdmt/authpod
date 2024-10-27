package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vndg-rdmt/authpod/internal/repository/users"
)

const migrateup = `
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
`

const migratedown = `
drop table if exists auth.websessions;
drop table if exists auth.tokens;
drop table if exists auth.users;
drop schema auth;
`

func initroot(repo users.Repository) {
	userlogin := os.Getenv("USE_ROOTUSER")
	if userlogin != "" {

		password := generatePassword(20, true, true, true)
		if _, err := repo.Create(context.Background(), userlogin, hashPassword(password)); err != nil {
			fmt.Printf("failed to init user: %v\n", err)
		} else {
			fmt.Println("new user created: write down the credetials:")
			fmt.Printf(" - login:    %s\n", userlogin)
			fmt.Printf(" - password: %s\n", password)
		}
	} else {
		fmt.Println("skipping root user creation")
	}
}

func migrate(postgresql *pgxpool.Pool) {
	flag := os.Getenv("USE_MIGRATION")

	switch flag {
	case "true":
		if _, err := postgresql.Exec(context.Background(), migrateup); err != nil {
			fmt.Printf("failed to migrate up: %v\n", err)
			if casted, ok := err.(*pgconn.PgError); ok {
				if casted.Code == "42P06" {
					fmt.Println("seems like just 'schema' already exists, no errors occuried")
					return
				}
				b, _ := json.Marshal(casted)
				fmt.Println(string(b))
			}
		} else {
			fmt.Println("migrated up")
		}
	case "false":
		if _, err := postgresql.Exec(context.Background(), migratedown); err != nil {

			fmt.Printf("failed to migrate down: %v\n", err)
			if casted, ok := err.(*pgconn.PgError); ok {
				b, _ := json.Marshal(casted)
				fmt.Println(string(b))
			}
		} else {
			fmt.Println("migrated down")
		}
	default:
		return
	}
}
