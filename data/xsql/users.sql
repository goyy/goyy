
-- mysql
drop table if exists users;
create table users (
  id varchar(50) not null,
  code varchar(50),
  name varchar(255),
  password varchar(255),
  memo varchar(1000),
  genre varchar(50),
  status varchar(50),
  roles varchar(1000),
  posts varchar(1000),
  org varchar(50),
  area varchar(50),
  creater varchar(50),
  created bigint,
  modifier varchar(50),
  modified bigint,
  version int(11),
  deletion int(11),
  primary key (id)
);

-- oracle
create table users (
  id varchar2(50) not null,
  code varchar2(50),
  name varchar2(255),
  password varchar2(255),
  memo varchar2(1000),
  genre varchar2(50),
  status varchar2(50),
  roles varchar2(1000),
  posts varchar2(1000),
  org varchar2(50),
  area varchar2(50),
  creater varchar2(50),
  created number(20),
  modifier varchar2(50),
  modified number(20),
  version number(11),
  deletion number(11),
  primary key (id)
);

insert into users (id, code, name, password, memo, genre, status, roles, posts, org, area, creater, created, modifier, modified, version, deletion) values ('aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 'aa', 1443253329, 'aa', 1443253329, 1, 0);
