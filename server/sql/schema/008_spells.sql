-- +goose Up
create table spells (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    description text,
    damage varchar(255),
    damage_type uuid references damage_types(id),
    saving_throw_ability_id uuid references abilities(id),
    range integer,
    spell_slot_level integer,
    school varchar(255)
);

-- +goose Down
drop table spells;
