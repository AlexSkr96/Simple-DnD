-- +goose Up
create table weapons_to_properties (
    id UUID primary key default gen_random_uuid(),
    weapon_id uuid references weapons(id),
    property_id uuid references weapon_properties(id)
);
