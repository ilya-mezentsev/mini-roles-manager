package schema

import "github.com/jmoiron/sqlx"

const Schema = `
create table if not exists account(
	id serial,
	hash character(32) primary key,
	created_at timestamp default current_timestamp
);

create table if not exists account_credentials(
	id serial,
	login character(100) unique,
	password character(32),
	account_hash character(32)
);
alter table account_credentials drop constraint if exists fk_account_credentials;
alter table account_credentials
add constraint fk_account_credentials
foreign key(account_hash)
references account(hash)
on delete cascade;


create table if not exists resource(
	id serial,
	resource_id character(100),
	title character(100) default '',
	links_to character(100)[] default array[]::character(100)[],
	account_hash character(32),
	created_at timestamp default current_timestamp,
	unique(resource_id, account_hash)
);
alter table resource drop constraint if exists fk_resource;
alter table resource
add constraint fk_resource
foreign key(account_hash)
references account(hash)
on delete cascade;


create table if not exists resource_permission(
	id serial,
	permission_id character(32) unique,
	operation character(10) not null,
	effect character(6) not null,
	resource_id character(100),
	account_hash character(32)
);
alter table resource_permission drop constraint if exists fk_resource_permission;
alter table resource_permission
add constraint fk_resource_permission
foreign key(resource_id, account_hash)
references resource(resource_id, account_hash)
on delete cascade;


create table if not exists role(
	id serial,
	role_id character(32),
	title character(100) default '',
	account_hash character(32),
	created_at timestamp default current_timestamp,
	unique(role_id, account_hash)
);
alter table role drop constraint if exists fk_role;
alter table role
add constraint fk_role
foreign key(account_hash)
references account(hash)
on delete cascade;


create table if not exists role_permission(
	id serial,
	permission_id character(32),
	role_id character(32),
	account_hash character(32),
	unique(permission_id, role_id, account_hash)
);
alter table role_permission drop constraint if exists fk_role_permission;
alter table role_permission
add constraint fk_role_permission
foreign key(role_id, account_hash)
references role(role_id, account_hash)
on delete cascade;

alter table role_permission drop constraint if exists fk_role_permission_permission_id;
alter table role_permission
add constraint fk_role_permission_permission_id
foreign key(permission_id)
references resource_permission(permission_id)
on delete cascade;


create table if not exists role_extending(
	id serial,
	extends_from character(32),
	role_id character(32),
	account_hash character(32),
	unique(extends_from, role_id, account_hash)
);
alter table role_extending drop constraint if exists fk_role_extending;
alter table role_extending
add constraint fk_role_extending
foreign key(extends_from, account_hash)
references role(role_id, account_hash)
on delete cascade;
`

func _(db *sqlx.DB) {
	db.MustExec(Schema)
}
