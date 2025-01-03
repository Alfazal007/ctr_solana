-- name: CreateProject :one
insert into project
    (id, name, creator_id)
        values ($1, $2, $3) returning *;

-- name: GetExistingProject :one
select * from project
	where name=$1;

-- name: GetExistingProjectById :one
select * from project
	where id=$1;

