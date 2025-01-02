-- +goose Up
create table project (
  id uuid primary key,
  name text not null,
	creator_id uuid,
	FOREIGN KEY(creator_id) REFERENCES users(id) on delete cascade
);

-- +goose Down
drop table project;
