-- name: GetCreatorBalance :one
select * from creator_balance
	where creator_id=$1;

-- name: DeductCreatorBalance :exec
update creator_balance
	set lamports=$1 where creator_id=$2;

-- name: AddPublicKey :exec
update creator_balance
	set creator_pk_bs64=$1 where creator_id=$2;

-- name: GetCreatorBalanceViaPK :one
select * from creator_balance
	where creator_pk_bs64=$1;
