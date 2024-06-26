-- +migrate Up
create table customers (
    id serial primary key,
    name varchar(64) not null
);

create table products (
    id serial primary key,
    sku varchar(64)
);

create table invoices (
    id serial primary key,
    customer_id int not null,
    total int not null,
    created_at timestamp default current_timestamp,
    FOREIGN key (customer_id) REFERENCES customers(id)
);

create table invoice_items (
    product_id int not null,
    invoice_id int not null,
    primary key (product_id, invoice_id),
    FOREIGN key (product_id) REFERENCES products(id),
    FOREIGN key (invoice_id) REFERENCES invoices(id)
);

insert into customers (name) values
('Customer 1'),
('Customer 2'),
('Customer 3');

insert into products (sku) values
('Product 1'),
('Product 2'),
('Product 3');

insert into invoices (customer_id, total) values
(1, 100),
(2, 200);

insert into invoice_items (invoice_id, product_id) values
(1, 1),
(1, 2),
(1, 3),
(2, 1),
(2, 3);

-- +migrate Down
drop table invoice_items;
drop table invoices;
drop table products;
drop table customers;
