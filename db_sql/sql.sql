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

alter table `users_db`.`users` add column `status` varchar(45) not null after `email`, add column `password` varchar(45) not null after `status`;