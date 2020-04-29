create database bienes_raices;
use bienes_raices;
create table if not exists propiedades( 
id int auto_increment not null,
precio decimal(16,2),
descripcion varchar(450),
propietario varchar(200),
vendedor varchar(200),
creado timestamp null default current_timestamp(),
primary key(`id`)
)Engine= InnoDB default CHARSET=utf8

show tables;