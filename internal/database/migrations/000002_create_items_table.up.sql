CREATE TABLE items (
                       id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                       user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                       title VARCHAR(255) NOT NULL,
                       description TEXT NOT NULL,
                       file_path TEXT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);
