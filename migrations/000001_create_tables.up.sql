create table if not exists music (
    id uuid primary key,
    song_name varchar(50),
    group_name varchar(50),
    text_song text,
    genre varchar(50),
    release_year int,
    duration float,
    create_at timestamp default current_timestamp
)