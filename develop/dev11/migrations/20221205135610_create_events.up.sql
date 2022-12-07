CREATE TABLE events (
    id bigserial not null primary key,
    user_id int not null,
    name varchar not null,
    date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);