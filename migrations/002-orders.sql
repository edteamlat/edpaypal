CREATE TABLE orders (
  id UUID NOT NULL,
  email VARCHAR(100) NOT NULL,
  is_product BOOLEAN NOT NULL DEFAULT false,
  is_subscription BOOLEAN NOT NULL DEFAULT false,
  product_id UUID,
  type_subs VARCHAR(10),
  price NUMERIC(7,2) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  CONSTRAINT orders_id_pk PRIMARY KEY (id),
  CONSTRAINT orders_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);
