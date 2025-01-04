-- +goose Up
create table votes (
	voter_id uuid not null,
	project_id uuid not null,
	public_id text not null,
	FOREIGN KEY(project_id) REFERENCES project(id) on delete cascade,
	FOREIGN KEY(voter_id) REFERENCES users(id) on delete cascade,
	FOREIGN KEY(public_id) REFERENCES project_images(public_id) on delete cascade,
    PRIMARY KEY (voter_id, project_id)
);

-- +goose Down
drop table votes;
