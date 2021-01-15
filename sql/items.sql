create table items
(
	id serial not null,
	name text,
	owner text not null
);

create unique index items_id_uindex
	on items (id);

alter table items
	add constraint items_pk
		primary key (id);

INSERT INTO items (name, owner) VALUES (?, ?)