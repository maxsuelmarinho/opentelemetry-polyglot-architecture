DROP TABLE IF EXISTS orders;
CREATE TABLE orders (
    id BIGSERIAL auto_increment NOT NULL,
    uuid varchar NOT NULL,
    user_id varchar NOT NULL,
    payment_method varchar NULL,
    tax_price numeric(19,0) NOT NULL,
    shipping_price numeric(19,0) NOT NULL,
    total_price numeric(19,0) NOT NULL,
    is_paid boolean NOT NULL,
    paid_at timestamp NULL,
    is_delivered boolean NOT NULL,
    delivered_at timestamp NULL,
    shipping_address TEXT NOT NULL,
    payment_result TEXT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
    CONSTRAINT orders_pkey PRIMARY KEY (id)
);

-- TODO: default current_timestamp is not working
DROP TABLE IF EXISTS order_items;
CREATE TABLE order_items (
    id BIGSERIAL auto_increment NOT NULL,
    uuid varchar NOT NULL,
    name varchar NOT NULL,
    qty INTEGER NOT NULL,
    image varchar NOT NULL,
    price numeric(19,0) NOT NULL,
    product_id varchar NOT NULL,
    order_id bigint NOT NULL,
    created_at timestamp DEFAULT NOW(),
    updated_at timestamp DEFAULT NOW(),
    CONSTRAINT order_items_pkey PRIMARY KEY (id),
    CONSTRAINT order_items_orders_fkey FOREIGN KEY (order_id) REFERENCES orders(id)
);
