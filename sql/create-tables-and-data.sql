-- Table: authors
drop table if exists authors cascade;
create table authors
(
  author_id serial,
  firstname character varying(255) not null,
  lastname character varying(255) not null,
  middlename character varying(255),
  about character varying(1000) not null,
  constraint authors_pkey primary key (author_id)
)
with (
  oids=false
);
alter table authors
  owner to postgres;

-- Table: books
drop table if exists books;
create table books
(
  isbn character(14) not null,
  title character varying(255) not null,
  price numeric(5,2) not null,
  constraint books_pkey primary key (isbn)
)
with (
  oids=false
);
alter table books
  owner to postgres;

-- Table: users
drop table if exists users cascade;
create table users
(
  id serial,
  firstname character varying(255) not null,
  lastname character varying(255) not null,
  middlename character varying(255),
  email character varying(255),
  username character varying(255) not null,
  password character varying(255) not null,
  constraint users_pkey primary key (id)
)
with (
  oids=false
);
alter table users owner to postgres;
drop index if exists username_index;
drop index if exists email_index;
create unique index username_index on users using btree (username);
create unique index email_index on users using btree (email);

-- Table: book_authors
drop table if exists book_authors;
create table book_authors
(
  id serial not null,
  book_isbn character varying(14) not null,
  author_id integer not null,
  constraint book_authors_pkey primary key (id)
)
with (
  oids=false
);
alter table book_authors
  owner to postgres;

-- password: password123
insert into users(firstname, lastname, email, username, password) values('Rodrigo', 'Duterte', 'admin@danubee.com', 'admin@danubee.com', '$2a$04$qao77tWi3XyupSmedOyFteQzr3fJHU/L4pB9iyfc3gdO4T1UaMnJa');
-- password: password123
insert into users(firstname, lastname, email, username, password) values('Rae', 'Burawes', 'user@danubee.com', 'user@danubee.com', '$2a$04$vM7m/jyMF1byJZo6vvW0muT36uOoIhHh7aUnGcqSj/lXqqqtLOe.m');

insert into authors (firstname, lastname, middlename, about) values ('Ryan', 'Holiday', null, 'RYAN HOLIDAY is a bestselling author and media strategist. After dropping out of college at nineteen to apprentice under Robert Greene, author of The 48 Laws of Power, he went on to advise many bestselling authors and multiplatinum musicians. He served as director of marketing at American Apparel for many years, where his campaigns have been used as case studies by Twitter, YouTube, and Google and written about in Ad Age, the New York Times, and Fast Company.
His first book, Trust Me I’m Lying: Confessions of a Media Manipulator—which the Financial Times called an “astonishing, disturbing book”—was a debut bestseller and is now taught in colleges around the world. He is also the author of Growth Hacker Marketing, Ego is the Enemy and The Daily Stoic. He lives in Austin, Texas, and writes for Thought Catalog and the New York Observer.');

insert into authors (firstname, lastname, middlename, about) values ('Adam', 'Grant', 'M.', 'Adam Grant has been Wharton’s top-rated teacher for four straight years. He has been recognized as one of the world’s 25 most influential management thinkers and the world’s top 40 business professors under 40.
Adam is the author of two New York Times bestselling books translated into 34 languages. Originals explores how individuals champion new ideas and leaders fight groupthink; it is a #1 national bestseller and one of Amazon’s best books of February 2016. Give and Take examines why helping others drives our success, and was named one of the best books of 2013 by Amazon, Apple, the Financial Times, and The Wall Street Journal—as well as one of Oprah’s riveting reads and Harvard Business Review’s ideas that shaped management.');

insert into authors (firstname, lastname, middlename, about) values ('Eric', 'Ries', null, 'ERIC RIES is an entrepreneur and author of the popular blog Startup Lessons Learned. He co-founded and served as CTO of IMVU, his third startup,  and has had plenty of startup failures along the way. He is a frequent speaker at business events, has advised a number of startups, large companies, and venture capital firms on business and product strategy, and is an Entrepreneur-in-Residence at Harvard Business School. His Lean Startup methodology has been written about in the New York Times, the Wall Street Journal, the Harvard Business Review, the Huffington Post, and many blogs. He lives in San Francisco.');

insert into books(isbn, title, price) values('978-0297868439', 'Give and Take: Why Helping Others Drives Our Success', 19.99);
insert into books(isbn, title, price) values('978-0307887894', 'The Lean Startup: How Constant Innovation Creates Radically Successful Businesses', 23.90);
insert into books(isbn, title, price) values('978-1591846352', 'The Obstacle Is the Way: The Timeless Art of Turning Trials into Triumph', 20.00);
insert into books(isbn, title, price) values('978-0143128854', 'Originals: How Non-Conformists Move the World', 16.90);
insert into books(isbn, title, price) values('978-1101903209', 'The Startup Way', 16.50);
insert into books(isbn, title, price) values('978-1591847816', 'Ego Is the Enemy', 18.95);

insert into book_authors(book_isbn, author_id) values ('978-0297868439', 2);
insert into book_authors(book_isbn, author_id) values ('978-0307887894', 3);
insert into book_authors(book_isbn, author_id) values ('978-1591846352', 1);
insert into book_authors(book_isbn, author_id) values ('978-1101903209', 3);
insert into book_authors(book_isbn, author_id) values ('978-0143128854', 2);
insert into book_authors(book_isbn, author_id) values ('978-1591847816', 1);