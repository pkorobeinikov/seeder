create schema app;

create table app.person
(
    id      uuid primary key,
    name    text not null,
    surname text not null
);

create table app.record
(
    id    uuid primary key,
    value text
);
