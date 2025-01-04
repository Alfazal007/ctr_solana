-- name: GetCreatorBalance :one
select * from creator_balance
	where creator_id=$1;

