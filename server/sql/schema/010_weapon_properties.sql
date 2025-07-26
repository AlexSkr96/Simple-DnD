-- +goose Up
create table weapon_properties (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    description text
);

insert into weapon_properties (name) values
('Light'),
('Heavy'),
('Versatile'),
('Two-handed'),
('Reach'),
('Finesse');

-- +goose Down
drop table weapon_properties;
