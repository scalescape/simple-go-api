CREATE TABLE users (
    id varchar NOT NULL,
    name varchar NOT NULL,
    created_at timestamp with time zone DEFAULT (now() at time zone 'utc')
);

