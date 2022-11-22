create table if not exists events
(
    id          serial
        constraint events_pk
            primary key,
    type        varchar(255)            not null,
    entity_type varchar(255)            not null,
    entity_id   integer                 not null,
    event_data  jsonb                   not null,
    created_at  timestamp default now() not null
);

