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

-- name: StartProject :one
update project set started=true
	where id=$1 returning *;

-- name: EndProject :one
update project set completed=true
	where id=$1 returning *;

-- name: CountRunningProjects :one
select count(*) from project 
	where started=true 
		AND completed=false
			AND creator_id=$1 limit 1;
