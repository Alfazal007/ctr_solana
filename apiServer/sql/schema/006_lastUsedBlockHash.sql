-- +goose Up
create table lastusedblock (
	lastUsedAddress text primary key
);

-- +goose Down
drop table lastusedblock;
