-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create the videos table
CREATE TABLE videos (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Create the video_tags table as a join table between users and videos
CREATE TABLE video_tags (
    user_id INT NOT NULL,
    video_id INT NOT NULL,
    PRIMARY KEY (user_id, video_id),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (video_id) REFERENCES videos(id)
);

-- Insert sample data into users table
INSERT INTO users (name) VALUES 
('Alice'),
('Bob'),
('Charlie');

-- Insert sample data into videos table
INSERT INTO videos (name) VALUES 
('Video 1'),
('Video 2'),
('Video 3');

-- Insert sample data into video_tags table
INSERT INTO video_tags (user_id, video_id) VALUES 
(1, 1),
(1, 2),
(2, 3),
(3, 1),
(3, 3);


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
