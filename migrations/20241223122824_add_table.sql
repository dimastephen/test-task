-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh(
    id serial primary key ,
    user_ip varchar(45) not null ,
    user_id INTEGER not null ,
    token_hash varchar(255) not null ,
    expires_at timestamp not null ,
    created_at timestamp not null default NOW(),
    updated_at timestamp null
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
