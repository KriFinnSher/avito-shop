CREATE TABLE IF NOT EXISTS Users (
                                     id UUID PRIMARY KEY,
                                     username TEXT NOT NULL,
                                     balance BIGINT NOT NULL DEFAULT 1000,
                                     password_hash TEXT NOT NULL
);
