CREATE SCHEMA `users_db`;
CREATE table `users_db`.`users`(
    `id` bigint(20) not null auto_increment,
    `first_name` varchar(45) null,
    `last_name` varchar(45) null,
    `email` varchar(45) null,
    `date_created` varchar(45) null,
    primary key (`id`),
    unique index `email_unique` (`email` ASC)
);


select * from users_db.users;