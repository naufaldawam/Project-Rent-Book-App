create database rent_book_app;

use rent_book_app;

create table users(
	id int primary key not null auto_increment,
	name_user varchar(100),
	email varchar(100),
	phone varchar(13),
	pass varchar(50),
	created_at datetime default current_timestamp,
	updated_at datetime default current_timestamp,
	deleted_at datetime default current_timestamp
); 

create table books(
	id int primary key not null auto_increment,
	user_id int,
	name_book varchar(100),
	content_book varchar(100),
	created_at datetime default current_timestamp,
	updated_at datetime default current_timestamp,
	deleted_at datetime default current_timestamp,
	foreign key (user_id) references users(id) on delete cascade
);

create table rent_books(
	id int primary key not null auto_increment,
	book_id int,
	user_id int,
	rent_date datetime default current_timestamp,
	return_date datetime,
	foreign key (book_id) references books(id) on delete cascade,
	foreign key (user_id) references users(id) on delete cascade
);

-- Check user history for book rental :
select u.name_user, u.email, u.phone, rb.rent_date, rb.return_date
from users u
inner join rent_books rb on rb.id_user = u.id_user;

