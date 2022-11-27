create table if not exists outbox
(
    event_id       integer not null
        primary key,
    event_type     varchar not null,
    aggregate_type varchar not null,
    aggregate_id   varchar not null,
    payload        json    not null,
    created_at     timestamp with time zone default now()
);



