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