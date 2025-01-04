-- name: GetExistingVote :one
select * from votes
	where voter_id=$1 AND project_id=$2;

-- name: CreateVote :one
insert into votes
    (voter_id, project_id, public_id)
        values ($1, $2, $3) returning *;

