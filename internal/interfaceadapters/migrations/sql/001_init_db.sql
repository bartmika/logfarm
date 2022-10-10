-- +goose Up

-- DEVELOPERS NOTE:
-- We need to make sure that we don't use `foreign keys` and other features
-- found via link [0] to support the `TiDB`.
-- [0] https://docs.pingcap.com/tidb/stable/mysql-compatibility#unsupported-features

CREATE TABLE IF NOT EXISTS records (
    id  VARCHAR (63) PRIMARY KEY NOT NULL,
    client VARCHAR (127) NOT NULL DEFAULT '',
    content VARCHAR (1027) NOT NULL DEFAULT '',
    facility INT NOT NULL DEFAULT 0,
    hostname VARCHAR (255) NOT NULL DEFAULT '',
    priority SMALLINT NOT NULL DEFAULT 0,
    severity SMALLINT NOT NULL DEFAULT 0,
    tag VARCHAR (255) NOT NULL DEFAULT (now()),
    timestamp DATETIME NOT NULL DEFAULT (now()),
    tls_peer VARCHAR (127) NOT NULL DEFAULT ''
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_records_id
ON records (id);
CREATE INDEX IF NOT EXISTS idx_records_timestamp
ON records (timestamp);

-- +goose Down

DROP TABLE records CASCADE;
