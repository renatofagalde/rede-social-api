drop database if exists redesocial;
create database redesocial default character set utf8 default collate utf8_general_ci;
use redesocial;
drop table if exists usuarios;
create table usuarios
(
    id        int auto_increment primary key,
    nome      varchar(50) not null,
    nick      varchar(50) not null unique,
    email     varchar(50) not null unique,
    senha     varchar(150) not null,
    criado_em timestamp default current_timestamp()
) engine = innodb character set utf8 collate utf8_unicode_ci;;

insert into usuarios (nome, nick, email, senha)
    value ("Ronaldo Nazário", "Fenomeno", "fenomeno@teste.com", 1),
    ("Ayrton Senna da Silva", "Senna", "senna@teste.com", 1),
    ("Romário", "Baixinho", "baixinho@teste.com", 1),
    ("Zico", "Zico", "zico@teste.com", 1);