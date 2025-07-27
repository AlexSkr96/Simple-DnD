-- +goose Up
-- +goose StatementBegin
create table user_sessions
(
    id         UUID primary key default gen_random_uuid(),
    user_id    uuid references users (id) on delete cascade,
    token      varchar(255) not null unique,
    expires_at timestamp    not null,
    created_at timestamp        default now()
);

create index idx_user_sessions_token on user_sessions(token);
create index idx_user_sessions_expires_at on user_sessions(expires_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop index if exists idx_user_sessions_expires_at;
drop index if exists idx_user_sessions_token;
drop table user_sessions;
-- +goose StatementEnd
