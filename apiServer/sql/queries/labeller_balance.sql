-- name: GetLabellerBalance :one
select * from labeller_balance
	where labeller_id=$1;

-- name: UpsertLabellerBalance :exec
INSERT INTO labeller_balance (labeller_id, lamports)
	VALUES ($1, $2)
		ON CONFLICT (labeller_id)
			DO UPDATE SET lamports = $2;
