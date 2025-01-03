-- +goose Up
create table project (
  id uuid primary key,
  name text not null,
  started boolean default false,
  completed boolean default false,
  creator_id uuid not null,
  FOREIGN KEY(creator_id) REFERENCES users(id) on delete cascade
);

-- +goose Down
drop table project;

