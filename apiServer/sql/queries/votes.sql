-- name: GetExistingVote :one
select * from votes
	where voter_id=$1 AND project_id=$2;

-- name: CreateVote :one
insert into votes
    (voter_id, project_id, public_id)
        values ($1, $2, $3) returning *;

-- name: GetVotesForProject :many
SELECT
    v.public_id,
    COUNT(v.public_id) AS vote_count,
    pi.secure_url
		FROM
			votes v
			JOIN
				project_images pi
				ON v.public_id = pi.public_id
				WHERE
					v.project_id = $1
					GROUP BY
						v.public_id,
						pi.secure_url;
