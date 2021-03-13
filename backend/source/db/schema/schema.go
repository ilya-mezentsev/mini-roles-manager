package scheme

const Schema = `
create table if not exists account(
	id serial,
	hash character(32) primary key
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
	resource_id character(100) primary key,
	title character(100) not null,
	links_to character(100)[]
	account_hash character(32)
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
	operation character(10),
	effect character(6) not null,
	resource_id character(100)
);
alter table permission drop constraint if exists fk_permission;
alter table permission
add constraint fk_permission
foreign key(resource_id)
references resource(resource_id)
on delete cascade;


create table if not exists role(
	id serial,
	role_id character(32) unique,
	title character(100),
	permissions character(32)[],
	extends character(32)[],
	account_hash character(32)	
);
alter table role drop constraint if exists fk_role;
alter table role
add constraint fk_role
foreign key(account_hash)
references account(hash)
on delete cascade;
`
