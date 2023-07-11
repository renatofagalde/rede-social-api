
select "criando database";
create database if not exists redesocial;

select "usando database";
use redesocial;

select "dropando tabela usuários";
drop table if exists usuarios;

select "criando tabela usuários";
create table usuarios
(
    id        int auto_increment primary key,
    nome      varchar(50) not null,
    nick      varchar(50) not null unique,
    email     varchar(50) not null unique,
    senha     varchar(50) not null,
    criado_em timestamp default current_timestamp()
);