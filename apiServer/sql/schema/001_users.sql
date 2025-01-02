-- +goose Up
create type user_role AS enum ('creator', 'labeller');

create table users (
    id uuid primary key,
    username text not null unique,
    password text not null,
    role user_role default 'labeller'
);

-- +goose Down
drop table users;
drop type user_role;
