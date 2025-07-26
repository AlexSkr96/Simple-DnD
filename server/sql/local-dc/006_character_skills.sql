-- +goose Up
create table character_skills (
    id UUID primary key default gen_random_uuid(),
    skill_id uuid references skills(id),
    character_id uuid references characters(id),
    value integer,
    is_proficient bool,
    unique(character_id, skill_id)
);
