-- +goose Up
create table damage_types (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null
);

insert into damage_types (name) values
('Acid'),
('Bludgeoning'),
('Cold'),
('Fire'),
('Force'),
('Lightning'),
('Necrotic'),
('Piercing'),
('Poison'),
('Psychic'),
('Radiant'),
('Slashing'),
('Thunder');

-- +goose Down
drop table damage_types;
