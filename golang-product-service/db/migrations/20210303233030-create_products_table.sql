
-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS products (
  id BIGSERIAL PRIMARY KEY,
  uuid uuid DEFAULT uuid_generate_v4(),
  name VARCHAR NOT NULL,
  image VARCHAR NOT NULL,
  brand VARCHAR NOT NULL,
  category VARCHAR NOT NULL,
  description VARCHAR NOT NULL,
  rating NUMERIC(2, 1) NOT NULL,
  num_reviews INT NOT NULL,
  price numeric(19,2) NOT NULL,
  count_in_stock INT NOT NULL,
  user_id VARCHAR NULL,
  created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
  updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

-- +migrate Down
DROP TABLE IF EXISTS products;
