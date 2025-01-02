-- name: CreateUser :one
insert into users
    (id, username, password, role)
        values ($1, $2, $3, $4) returning *;

-- name: CheckSimilarUserExists :one
select count(*) from users
	where username=$1 limit 1;

-- name: GetUserByUsername :one
select * from users
	where username=$1 limit 1;

