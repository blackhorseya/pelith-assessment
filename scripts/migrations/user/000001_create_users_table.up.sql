CREATE TABLE IF NOT EXISTS users
(
    address  bytea PRIMARY KEY check ( octet_length(address) = 20 ),
    username VARCHAR(50) UNIQUE NOT NULL
);