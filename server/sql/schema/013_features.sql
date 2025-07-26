-- +goose Up
create table features (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    description text,
    uses integer,
    reset varchar(255), --'short rest', 'long rest', 'daily'
    target varchar(255),
    modificator int
);

-- +goose Down
drop table features;
