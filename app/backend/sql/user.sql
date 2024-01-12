CREATE TABLE users (
    user_id integer primary key,
    user_name varchar(100) not null,
    email_address varchar(100) not null,
    hashed_password varchar(100) not null
);