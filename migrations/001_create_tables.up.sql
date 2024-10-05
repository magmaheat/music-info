create table if not exists music-library (
    id uuid primary key,
    group varchar(50),
    name varchar(50),
    create_at timestamp default current_timestamp
)