-- name: CreateUser :one
insert into users
    (id, username, password)
        values ($1, $2, $3) returning *;
