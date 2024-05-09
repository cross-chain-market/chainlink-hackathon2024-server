BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS items (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    collection_id UUID NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    base_image_path TEXT,
    fiat_price DECIMAL,
    address TEXT,
    total_amount BIGINT,
    listed_amount BIGINT,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT fk_collections_id FOREIGN KEY (collection_id) REFERENCES collections(id),
    CONSTRAINT uc_item_name_collection_id UNIQUE (collection_id, name)
);


CREATE INDEX IF NOT EXISTS idx_items_collection_id ON items (collection_id);
CREATE INDEX IF NOT EXISTS idx_listed_amount ON items (listed_amount);

COMMIT;