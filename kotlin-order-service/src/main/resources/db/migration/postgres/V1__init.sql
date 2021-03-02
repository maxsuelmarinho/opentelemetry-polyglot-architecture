CREATE TABLE IF NOT EXISTS orders (
  id BIGSERIAL PRIMARY KEY,
  uuid varchar NOT NULL UNIQUE,
  user_id varchar NOT NULL,
  payment_method varchar NULL,
  tax_price numeric(19,0) NOT NULL,
  shipping_price numeric(19,0) NOT NULL,
  total_price numeric(19,0) NOT NULL,
  is_paid boolean NOT NULL,
  paid_at timestamp NULL,
  is_delivered boolean NOT NULL,
  delivered_at timestamp NULL,
  shipping_address JSONB NOT NULL,
  payment_result JSONB NULL,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE order_items (
    id BIGSERIAL PRIMARY KEY,
    uuid varchar NOT NULL,
    name varchar NOT NULL,
    qty INTEGER NOT NULL,
    image varchar NOT NULL,
    price numeric(19,0) NOT NULL,
    product_id varchar NOT NULL,
    order_id bigint NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW(),
    CONSTRAINT order_items_orders_fkey FOREIGN KEY (order_id) REFERENCES orders(id)
);
