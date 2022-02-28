CREATE TABLE subscriptions (
  id UUID NOT NULL,
  email VARCHAR(100) NOT NULL,
  status VARCHAR(100) NOT NULL,
  type_subs VARCHAR(100) NOT NULL,
  begins_at DATE NOT NULL DEFAULT now(),
  ends_at DATE NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT now(),
  updated_at TIMESTAMP,
  CONSTRAINT subscriptions_id_pk PRIMARY KEY (id)
);
