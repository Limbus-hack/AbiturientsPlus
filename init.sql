create table if not exists public.users
(
    id         bigint unique not null,
    name       varchar       not null,
    last_name  varchar       not null,
    region     int           not null,
    prediction int           not null,
    status     varchar       not null
);