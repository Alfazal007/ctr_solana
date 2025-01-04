-- name: GetCreatorBalance :one
select * from creator_balance
	where creator_id=$1;

-- name: DeductCreatorBalance :exec
update creator_balance
	set lamports=$1 where creator_id=$2;

