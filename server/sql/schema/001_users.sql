-- +goose Up
create table users (
    id UUID primary key default gen_random_uuid(),
    email varchar(255),
    username varchar(255),
    password_hash varchar(255)
);

-- +goose Down
drop table users;
