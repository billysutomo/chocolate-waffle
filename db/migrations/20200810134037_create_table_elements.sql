-- migrate:up
CREATE TABLE elements (
    id SERIAL,
    id_project INTEGER,
    ordernum INTEGER,
    type VARCHAR(255),
    body VARCHAR(255),
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
)

-- migrate:down
DROP TABLE elements