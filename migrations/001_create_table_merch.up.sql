CREATE TABLE IF NOT EXISTS Merch (
                                  id UUID PRIMARY KEY,
                                  name TEXT NOT NULL,
                                  cost BIGINT NOT NULL CHECK (cost >= 0)
    );
