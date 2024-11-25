-- Create the transactions table
CREATE TABLE IF NOT EXISTS transactions
(
    tx_hash      TEXT PRIMARY KEY, -- Transaction hash
    block_number INTEGER,          -- Block number
    timestamp    TIMESTAMP,        -- Transaction timestamp
    from_address TEXT,             -- Sender address
    to_address   TEXT,             -- Receiver address
    value        NUMERIC,          -- Transaction value
    gas_used     INTEGER,          -- Gas used
    gas_price    NUMERIC,          -- Gas price
    status       BOOLEAN           -- Transaction status (success or failure)
);

-- Create the swap_events table
CREATE TABLE IF NOT EXISTS swap_events
(
    id                 SERIAL PRIMARY KEY,                  -- Primary key
    tx_hash            TEXT NOT NULL,                       -- Associated transaction hash
    from_token_address TEXT NOT NULL,                       -- Source token address
    to_token_address   TEXT NOT NULL,                       -- Destination token address
    from_token_amount  TEXT NOT NULL,                       -- Source token amount
    to_token_amount    TEXT NOT NULL,                       -- Destination token amount
    pool_address       TEXT,                                -- Swap pool address (if applicable)
    FOREIGN KEY (tx_hash) REFERENCES transactions (tx_hash) -- Foreign key to transactions
);

-- Create indexes for the swap_events table
CREATE INDEX idx_swap_events_tx_hash ON swap_events (tx_hash);
CREATE INDEX idx_swap_events_from_token ON swap_events (from_token_address);
CREATE INDEX idx_swap_events_to_token ON swap_events (to_token_address);
