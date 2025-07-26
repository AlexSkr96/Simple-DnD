-- +goose Up
create table character_abilities (
    id UUID primary key default gen_random_uuid(),
    ability_id uuid references abilities(id),
    character_id uuid references characters(id),
    value int,
    is_proficient_saving_throw bool,
    unique(character_id, ability_id)
);

-- +goose Down
drop table character_abilities;
