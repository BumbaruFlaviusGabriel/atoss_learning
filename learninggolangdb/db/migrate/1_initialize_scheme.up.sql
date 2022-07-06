BEGIN;

create extension if not exists "uuid-ossp";
CREATE TABLE customers (
    id uuid DEFAULT uuid_generate_v4(),
    name varchar(255) NOT NULL UNIQUE,
    address varchar(255) NOT NULL,
    description text NOT NULL,
    CONSTRAINT pk_customer_id PRIMARY KEY (id)
);


CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4(),
    email varchar(255) NOT NULL UNIQUE,
    password varchar(255) NOT NULL,
    role varchar(255) NOT NULL,
    first_name varchar(255) NOT NULL,
    last_name varchar(255) NOT NULL,
    contact_number varchar(255) NOT NULL,
    customer_id uuid NOT NULL,
    CONSTRAINT pk_user_id PRIMARY KEY (id),
    CONSTRAINT fk_user_customer FOREIGN KEY (customer_id) REFERENCES customers(id)
);

COMMIT;