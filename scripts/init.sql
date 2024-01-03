CREATE TABLE IF NOT EXISTS t_block(
    block_num bigint UNIQUE NOT NULL,
    block_hash character varying(128) UNIQUE NOT NULL,
    block_time bigint NOT NULL,
   	parent_hash character varying(128) NOT NULL,
    create_time timestamp NOT NULL DEFAULT (now())::timestamp,
    update_time timestamp NOT NULL DEFAULT (now())::timestamp,
    PRIMARY KEY(block_num)
);

CREATE INDEX IF NOT EXISTS block_num_idx ON t_block (block_num);

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

CREATE TABLE IF NOT EXISTS t_log(
    tx_hash character varying(128) REFERENCES t_transaction(tx_hash) NOT NULL,
    index bigint NOT NULL,
    data varchar NOT NULL,
    create_time timestamp NOT NULL DEFAULT (now())::timestamp,
    update_time timestamp NOT NULL DEFAULT (now())::timestamp,
    PRIMARY KEY(tx_hash, index)
);

CREATE INDEX IF NOT EXISTS tx_hash_index_idx ON t_log (tx_hash, index);