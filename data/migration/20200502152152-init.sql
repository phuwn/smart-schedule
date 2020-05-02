
-- +migrate Up
create function gen_random_uuid() returns uuid
    language c
    as '$libdir/pgcrypto', 'pg_random_uuid';

create table lists (
	id uuid not null unique primary key default gen_random_uuid(),
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200)
);

create table statuses (
	id uuid not null unique primary key default gen_random_uuid(),
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200),
	color varchar(200)
);

create table tags (
	id uuid not null unique primary key default gen_random_uuid(),
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	name varchar(200),
	color varchar(200)
);

create table tickets (
	id uuid not null unique primary key default gen_random_uuid(),
	created_at timestamp(6) with time zone default now(),
	updated_at timestamp(6) with time zone,
	deleted_at timestamp(6) with time zone,
	list_id uuid references lists (id) on delete cascade not null,
	name varchar(200),
	notes text,
	status int,
	priority int,
	tags json,
	remind boolean,
	remind_date timestamp(6) with time zone,
	repeat int 
);

-- +migrate Down
drop table tickets;
drop table tags;
drop table statuses;
drop table lists;
drop function gen_random_uuid;