CREATE TABLE IF NOT EXISTS t_transaction(
    tx_hash character varying(128) UNIQUE NOT NULL,
    block_num bigint references t_block(block_num) NOT NULL,
    "from" character varying(128) NOT NULL,
    "to" character varying(128) NOT NULL,
    nonce bigint NOT NULL,
    data varchar NOT NULL,
    value character varying(300) NOT NULL,
    create_time timestamp NOT NULL DEFAULT (now())::timestamp,
    update_time timestamp NOT NULL DEFAULT (now())::timestamp,
    PRIMARY KEY(tx_hash)
);

CREATE INDEX IF NOT EXISTS tx_hash_idx ON t_transaction (tx_hash);