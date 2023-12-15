CREATE TABLE IF NOT EXISTS t_log(
    tx_hash character varying(128) REFERENCES t_transaction(tx_hash) NOT NULL,
    index bigint NOT NULL,
    data varchar NOT NULL,
    create_time timestamp NOT NULL DEFAULT (now())::timestamp,
    update_time timestamp NOT NULL DEFAULT (now())::timestamp,
    PRIMARY KEY(tx_hash, index)
);

CREATE INDEX IF NOT EXISTS tx_hash_index_idx ON t_log (tx_hash, index);