-- +goose Up
create table characters (
    id UUID primary key default gen_random_uuid(),
    name varchar(255) not null,
    index varchar(255),
    class varchar(255),
    race varchar(255),
    alignment varchar(255),
    user_id uuid references users(id),
    armour_class integer not null,
    current_experience integer default 0,
    death_save_pos integer default 0,
    death_save_neg integer default 0,
    hit_dice_left varchar(255),
    total_hit_dice varchar(255),
    max_hp integer not null,
    current_hp integer not null,
    temp_hp integer not null default 0,
    proficiency_bonus integer,
    inspiration boolean default false,
    speed integer not null default 30
);