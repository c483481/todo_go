CREATE TABLE todos (
   id BIGSERIAL PRIMARY KEY,
   xid VARCHAR(26) UNIQUE NOT NULL,
   version INTEGER NOT NULL DEFAULT 1,
   created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
   title VARCHAR(255) NOT NULL CHECK (LENGTH(title) <= 255),
   description TEXT NOT NULL
);