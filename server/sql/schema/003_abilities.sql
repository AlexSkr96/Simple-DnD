-- +goose Up
create table abilities (
    id UUID primary key default gen_random_uuid(),
    name varchar(255)
);

insert into abilities (name) values ('Strength');
insert into abilities (name) values ('Dexterity');
insert into abilities (name) values ('Constitution');
insert into abilities (name) values ('Intelligence');
insert into abilities (name) values ('Wisdom');
insert into abilities (name) values ('Charisma');

-- +goose Down
drop table abilities;
