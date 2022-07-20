create table "user"(
    id uuid primary key,
    username varchar(255) not null,
    hash_password text not null
)