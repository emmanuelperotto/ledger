create table if not exists events
(
    id          integer generated always as identity
        constraint events_pk
            primary key,
    event_type  varchar                                not null,
    entity_type varchar                                not null,
    entity_id   varchar                                not null,
    event_data  json                                   not null,
    created_at  timestamp with time zone default now() not null
);