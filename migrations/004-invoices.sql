CREATE TABLE invoices (
  id UUID NOT NULL,
  invoice_date DATE NOT NULL DEFAULT now(),
  email VARCHAR(36) NOT NULL,
  is_product BOOLEAN NOT NULL DEFAULT false,
  is_subscription BOOLEAN NOT NULL DEFAULT false,
  product_id UUID,
  subscription_id UUID,
  price NUMERIC(7,2) NOT NULL DEFAULT 0,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  CONSTRAINT invoices_id_pk PRIMARY KEY (id),
  CONSTRAINT invoices_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT invoices_subscription_id_fk FOREIGN KEY (subscription_id) REFERENCES subscriptions (id) ON UPDATE RESTRICT ON DELETE RESTRICT
);
