-- +goose Up
-- game room for online gaming
create table game_rooms (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    owner_id UUID not null references users(id) -- owner is always a game master
);

-- +goose Down
drop table game_rooms;
