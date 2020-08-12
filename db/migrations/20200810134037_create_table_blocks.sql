-- migrate:up
CREATE TABLE blocks (
    id SERIAL,
    id_project INTEGER,
    ordernum INTEGER,
    type VARCHAR(255),
    body JSONB,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    deleted_at TIMESTAMPTZ
)

-- migrate:down
DROP TABLE blocks