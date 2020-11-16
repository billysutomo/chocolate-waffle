-- migrate:up
CREATE TABLE users (
    id SERIAL,
    name VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
) 

-- migrate:down
DROP TABLE users