--users
create or replace function upd_timestamp() returns trigger as
$$
begin
    new.gmt_update = current_timestamp;
    return new;
end
$$
language plpgsql;

drop table if exists users;

create table users (
id serial primary key,
user_id varchar(64) not null unique,
user_name varchar(255) not null,
password varchar(255) not null,
gender smallint default 0,
gmt_create timestamp default current_timestamp,
gmt_update timestamp default current_timestamp
);

create trigger upd_timestamp before update on users for each row execute procedure upd_timestamp();

--relationships
drop table if exists user_relationships;

create table user_relationships (
id serial primary key,
user_id varchar(64) not null,
follow_user_id varchar(64) not null,
state smallint default 0,
gmt_create timestamp default current_timestamp,
gmt_update timestamp default current_timestamp,
constraint uniq_user_follow unique (user_id, follow_user_id)
);

create trigger upd_timestamp before update on user_relationships for each row execute procedure upd_timestamp();

--test data
insert into users (user_id, user_name, password, gender)
values ('fad85910-57c9-415f-9eb0-ced37e9309b9', 'Taki', '', 1),
('c40e039e-c428-4adf-9dca-4060a486f624', 'Hotaka', '', 1),
('2f9a17c6-f520-4de0-a2ac-4a4e71a8f173', 'Mitsuha', '', 2),
('c54f2fe8-c144-4c1e-b8a5-9606a477dcd4', 'Hina', '', 2);

insert into user_relationships (user_id, follow_user_id, state)
values ('fad85910-57c9-415f-9eb0-ced37e9309b9', '2f9a17c6-f520-4de0-a2ac-4a4e71a8f173', 3),
('c40e039e-c428-4adf-9dca-4060a486f624', 'c54f2fe8-c144-4c1e-b8a5-9606a477dcd4', 1),
('2f9a17c6-f520-4de0-a2ac-4a4e71a8f173', 'fad85910-57c9-415f-9eb0-ced37e9309b9', 3),
('c54f2fe8-c144-4c1e-b8a5-9606a477dcd4', 'c40e039e-c428-4adf-9dca-4060a486f624', 2);
