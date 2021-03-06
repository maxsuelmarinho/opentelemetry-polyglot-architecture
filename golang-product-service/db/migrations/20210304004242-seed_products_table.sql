
-- +migrate Up
INSERT INTO products (name, image, description, brand, category, price, count_in_stock, rating, num_reviews)
VALUES (
  'Airpods Wireless Bluetooth Headphones',
  '/images/airpods.jpg',
  'Bluetooth technology lets you connect it with compatible devices wirelessly High-quality AAC audio offers immersive listening experience Built-in microphone allows you to take calls while working',
  'Apple',
  'Electronics',
  89.99,
  10,
  4.5,
  12
),
(
  'iPhone 11 Pro 256GB Memory',
  '/images/phone.jpg',
  'Introducing the iPhone 11 Pro. A transformative triple-camera system that adds tons of capability without complexity. An unprecedented leap in battery life',
  'Apple',
  'Electronics',
  599.99,
  7,
  4.0,
  8
),
(
  'Cannon EOS 80D DSLR Camera',
  '/images/camera.jpg',
  'Characterized by versatile imaging specs, the Canon EOS 80D further clarifies itself using a pair of robust focusing systems and an intuitive design',
  'Cannon',
  'Electronics',
  929.99,
  5,
  3,
  12
),
(
  'Sony Playstation 4 Pro White Version',
  '/images/playstation.jpg',
  'The ultimate home entertainment center starts with PlayStation. Whether you are into gaming, HD movies, television, music',
  'Sony',
  'Electronics',
  399.99,
  11,
  5,
  12
),
(
  'Logitech G-Series Gaming Mouse',
  '/images/mouse.jpg',
  'Get a better handle on your games with this Logitech LIGHTSYNC gaming mouse. The six programmable buttons allow customization for a smooth playing experience',
  'Logitech',
  'Electronics',
  49.99,
  7,
  3.5,
  10
),
(
  'Amazon Echo Dot 3rd Generation',
  '/images/alexa.jpg',
  'Meet Echo Dot - Our most popular smart speaker with a fabric design. It is our most compact smart speaker that fits perfectly into small space',
  'Amazon',
  'Electronics',
  29.99,
  0,
  4,
  12
);

-- +migrate Down
DELETE FROM products;
