-- +goose Up
create table character_prepared_spells(
    id UUID primary key default gen_random_uuid(),
    spell_id UUID references spells(id),
    character_id UUID references characters(id)
);
