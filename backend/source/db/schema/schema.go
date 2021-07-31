package schema

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


create table if not exists permission(
	id serial,
	permission_id character(32) unique,
	operation character(10) not null,
	effect character(6) not null,
	resource_id character(100),
	account_hash character(32)
);
alter table permission drop constraint if exists fk_permission;
alter table permission
add constraint fk_permission
foreign key(resource_id, account_hash)
references resource(resource_id, account_hash)
on delete cascade;


create table if not exists roles_version(
	id serial,
	version_id character(32),
	title character(100) default '',
	account_hash character(32),
	created_at timestamp default current_timestamp,
	unique(version_id, account_hash)
);
alter table roles_version drop constraint if exists fk_roles_version;
alter table roles_version
add constraint fk_roles_version
foreign key(account_hash)
references account(hash)
on delete cascade;


create table if not exists role(
	id serial,
	role_id character(32),
	title character(100) default '',
	permissions character(32)[] default array[]::character(32)[],
	extends character(32)[] default array[]::character(32)[],
	account_hash character(32),
	version_id character(32),
	created_at timestamp default current_timestamp,
	unique(role_id, version_id, account_hash)
);
alter table role drop constraint if exists fk_role;
alter table role
add constraint fk_role
foreign key(version_id, account_hash)
references roles_version(version_id, account_hash)
on delete cascade;
`
