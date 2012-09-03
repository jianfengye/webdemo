create table webdemo_admin
(
    admin_id int not null auto_increment,
    admin_name varchar(32) not null,
    admin_password varchar(32) not null,
    primary key(admin_id)
);