-- migrate:up
CREATE TABLE projects (
    id SERIAL,
    id_user INTEGER,
    url VARCHAR(255),
    profile_picture VARCHAR(255),
    title VARCHAR(255),
    description VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
)

-- migrate:down
DROP TABLE projects

