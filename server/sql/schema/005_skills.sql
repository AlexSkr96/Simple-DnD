-- +goose Up
create table skills (
    id UUID primary key default gen_random_uuid(),
    name varchar(255),
    base_ability_id uuid references abilities(id)
);

-- +goose Down
drop table skills;
