CREATE TABLE IF NOT EXISTS Transactions (
                                            id UUID PRIMARY KEY,
                                            from_user TEXT NOT NULL,
                                            type TEXT CHECK (type IN ('transfer', 'purchase')) NOT NULL,
                                            amount BIGINT NOT NULL,
                                            to_user TEXT NOT NULL,
                                            item TEXT NOT NULL,
                                            date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
