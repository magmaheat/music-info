create extension if not exists "uuid-ossp";


create table if not exists music (
    id uuid primary key default uuid_generate_v4(),
    song_name varchar(50),
    group_name varchar(50),
    text_song text default '',
    genre varchar(50),
    release_year int,
    duration float,
    create_at timestamp default current_timestamp
)