CREATE TABLE Transactions (
                              id UUID PRIMARY KEY,
                              user_id UUID REFERENCES users(id) ON DELETE CASCADE,
                              type TEXT CHECK (type IN ('transfer', 'purchase')) NOT NULL,
                              amount BIGINT NOT NULL,
                              target_user_id UUID REFERENCES users(id),
                              item_id UUID REFERENCES items(id),
                              date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
