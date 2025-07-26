-- +goose Up
create table weapons(
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    is_ranged bool,
    optimal_range integer,
    max_range integer,
    cost integer,
    damage varchar(255),
    damage_type_id uuid references damage_types(id)
);
