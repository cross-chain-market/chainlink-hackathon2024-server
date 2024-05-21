BEGIN;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS collections (
    id  BIGSERIAL PRIMARY KEY,
    owner_address_hex TEXT NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    base_hash TEXT,
    address TEXT,
    status TEXT,
    network_id TEXT NOT NULL,
    chain_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ,

    CONSTRAINT uc_collection_name_owner_address_chain_id UNIQUE (owner_address_hex, name, chain_id)

    );


CREATE INDEX IF NOT EXISTS idx_collections_owner_address_hex ON collections(owner_address_hex);

COMMIT;