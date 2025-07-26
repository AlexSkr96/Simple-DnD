-- +goose Up
create table character_features (
    id UUID primary key default gen_random_uuid(),
    character_id UUID references characters(id),
    feature_id UUID references features(id)
);

-- +goose Down
drop table character_features;
