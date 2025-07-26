-- +goose Up
create table skills (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    base_ability_id uuid references abilities(id)
);

insert into skills (name, base_ability_id)
select skill_name, a.id from abilities a
cross join (values
    ('Acrobatics', 'Dexterity'),
    ('Animal Handling', 'Wisdom'),
    ('Arcana', 'Intelligence'),
    ('Athletics', 'Strength'),
    ('Deception', 'Charisma'),
    ('History', 'Intelligence'),
    ('Insight', 'Wisdom'),
    ('Intimidation', 'Charisma'),
    ('Investigation', 'Intelligence'),
    ('Medicine', 'Wisdom'),
    ('Nature', 'Intelligence'),
    ('Perception', 'Wisdom'),
    ('Performance', 'Charisma'),
    ('Persuasion', 'Charisma'),
    ('Religion', 'Intelligence'),
    ('Sleight of Hand', 'Dexterity'),
    ('Stealth', 'Dexterity'),
    ('Survival', 'Wisdom')
) as skill_data(skill_name, ability_name)
where a.name = skill_data.ability_name;

-- +goose Down
drop table skills;
