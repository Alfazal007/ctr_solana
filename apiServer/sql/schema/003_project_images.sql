-- +goose Up
create table project_images (
  id uuid primary key,
	project_id uuid,
	FOREIGN KEY(project_id) REFERENCES project(id) on delete cascade
);

-- +goose Down
drop table project_images;
