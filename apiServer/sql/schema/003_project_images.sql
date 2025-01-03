-- +goose Up
create table project_images (
	public_id text primary key,
	project_id uuid not null,
	secure_url text not null,
	FOREIGN KEY(project_id) REFERENCES project(id) on delete cascade
);

-- +goose Down
drop table project_images;
