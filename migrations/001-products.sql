CREATE TABLE products (
  id UUID NOT NULL,
  name VARCHAR(50) NOT NULL,
  description VARCHAR(50) NOT NULL,
  image VARCHAR(1024) NOT NULL,
  is_subscription BOOLEAN NOT NULL DEFAULT false,
  months INTEGER NOT NULL DEFAULT 0,
  price NUMERIC(7,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  CONSTRAINT products_id_pk PRIMARY KEY (id)
);
