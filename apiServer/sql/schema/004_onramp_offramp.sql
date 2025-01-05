-- +goose Up
create table creator_balance (
	creator_id uuid primary key,
	lamports text not null,
	creator_pk_bs64 text,
  FOREIGN KEY(creator_id) REFERENCES users(id) on delete cascade
);

create table labeller_balance (
	labeller_id uuid primary key,
	lamports text not null,
	FOREIGN KEY(labeller_id) REFERENCES users(id) on delete cascade
);

-- +goose Down
drop table creator_balance;
drop table labeller_balance;
