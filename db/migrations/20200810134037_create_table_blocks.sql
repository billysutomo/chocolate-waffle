-- migrate:up
CREATE TABLE blocks (
    id SERIAL,
    id_project INTEGER,
    ordernum INTEGER,
    url VARCHAR(255),
    icon VARCHAR(255),
    title VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
) 

-- migrate:down
DROP TABLE blocks