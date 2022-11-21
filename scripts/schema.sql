create table if not exists outbox
(
    event_id       uuid         not null primary key,
    event_type     varchar(255) not null,
    aggregate_type varchar(255) not null,
    aggregate_id   varchar(255) not null,
    payload        jsonb,
    created_at     timestamp default now()
);

