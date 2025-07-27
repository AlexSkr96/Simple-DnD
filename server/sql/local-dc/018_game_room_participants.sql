-- +goose Up
create table game_room_participants (
    id           UUID primary key default gen_random_uuid(),
    user_id      UUID references users (id),
    character_id UUID references characters (id),
    game_room_id UUID references game_rooms (id)
);
