create table clients
(
    id                uuid not null
        primary key,
    last_name         varchar(20),
    first_name        varchar(20),
    patronymic        varchar(20),
    age               integer,
    registration_date varchar(20)
);

alter table clients
    owner to postgres;

create table markets
(
    id      uuid not null
        primary key,
    name    varchar(20),
    address varchar(50),
    active  boolean,
    owner   varchar(20)
);

alter table markets
    owner to postgres;
