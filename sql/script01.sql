drop database if exists redesocial;
create database redesocial default character set utf8 default collate utf8_general_ci;
use redesocial;
drop table if exists usuarios;
drop table if exists seguidores;
create table usuarios
(
    id        int auto_increment primary key,
    nome      varchar(50)  not null,
    nick      varchar(50)  not null unique,
    email     varchar(50)  not null unique,
    senha     varchar(150) not null,
    criado_em timestamp default current_timestamp()
) engine = innodb
  character set utf8
  collate utf8_unicode_ci;;

create table seguidores
(
    usuario_id  int not null,
    foreign key (usuario_id) references usuarios (id)
        on delete cascade,

    seguidor_id int not null,
    foreign key (usuario_id) references usuarios (id)
        on delete cascade,

    primary key (usuario_id, seguidor_id)
) engine = innodb;


insert into usuarios (nome, nick, email, senha)
    value
    ("Ronaldo Nazário", "Fenomeno", "fenomeno@teste.com",
     '$2a$12$S7wbqJTXtgp6kyEkrVTiGOr.fbeLX0ShD2J9ftMr3DMRiJyY.VNTS'),
    ("Ayrton Senna da Silva", "Senna", "senna@teste.com",
     '$2a$12$S7wbqJTXtgp6kyEkrVTiGOr.fbeLX0ShD2J9ftMr3DMRiJyY.VNTS'),
    ("Romário", "Baixinho", "baixinho@teste.com", '$2a$12$S7wbqJTXtgp6kyEkrVTiGOr.fbeLX0ShD2J9ftMr3DMRiJyY.VNTS'),
    ("Zico", "Zico", "zico@teste.com", '$2a$12$S7wbqJTXtgp6kyEkrVTiGOr.fbeLX0ShD2J9ftMr3DMRiJyY.VNTS');

insert into seguidores (usuario_id, seguidor_id)
values (1, 2),
       (1, 3),
       (1, 4),
       (2, 1)
