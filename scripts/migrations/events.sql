create table if not exists events
(
    id          integer generated always as identity
        constraint events_pk
            primary key,
    type        varchar                                not null,
    entity_type varchar                                not null,
    entity_id   integer                                not null,
    event_data  jsonb                                  not null,
    created_at  timestamp with time zone default now() not null
);
