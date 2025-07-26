-- +goose Up
create table character_spell_slots
(
    id           UUID primary key default gen_random_uuid(),
    character_id UUID references characters (id),
    level        integer,
    max          integer,
    slots_left   integer
);

-- +goose Down
drop table character_spell_slots;
