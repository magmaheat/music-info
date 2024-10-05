create table if not exists music (
    id uuid primary key,
    name varchar(50),
    group varchar(50),
    text_song text,
    genre varchar(50),
    release_year int,
    duration float
    create_at timestamp default current_timestamp
)