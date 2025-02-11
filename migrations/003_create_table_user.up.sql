CREATE TABLE IF NOT EXISTS Users (
                       id UUID PRIMARY KEY,
                       name TEXT NOT NULL,
                       balance BIGINT NOT NULL DEFAULT 1000
);