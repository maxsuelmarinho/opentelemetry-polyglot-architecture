
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS reviews (
  id BIGSERIAL PRIMARY KEY,
  uuid uuid DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL,
  rating INT NOT NULL,
  comment VARCHAR NOT NULL,
  user_id VARCHAR NOT NULL,
  product_id BIGINT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  CONSTRAINT reviews_products_fkey FOREIGN KEY (product_id) REFERENCES products(id)
);

-- +migrate Down
DROP TABLE IF EXISTS reviews;
