create schema app;

create table app.person (
    id      uuid primary key,
    name    text not null,
    surname text not null
);
