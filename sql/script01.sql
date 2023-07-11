select "criando database";
drop database if exists redesocial;
create database redesocial default character set utf8 default collate utf8_general_ci;

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
) engine = innodb character set utf8 collate utf8_unicode_ci;;

insert into usuarios value (null, "Ronaldo Nazário", "Fenomeno", "fenomeno@teste.com", 1, now()),
    (null, "Ayrton Senna da Silva", "Senna", "senna@teste.com", 1, now()),
    (null, "Romário", "Baixinho", "baixinho@teste.com", 1, now()),
    (null, "Zico", "Zico", "zico@teste.com", 1, now());