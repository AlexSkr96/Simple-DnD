-- +goose Up
create table armor (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    ac integer not null,
    str_req integer not null default 0,
    stealth_disadvantage boolean not null default false,
    weight integer,
    cost integer
);
