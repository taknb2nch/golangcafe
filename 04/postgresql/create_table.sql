-- 
drop table if exists table1 cascade;

create table table1 (
  id serial not null
  , display_name character varying(100)
  , sex character(1) not null
  , birthday timestamp
  , age integer
  , married boolean
  , rate real
  , salary numeric(10,0)
  , constraint table1_PKC primary key (id)
) ;

comment on column table1.id is 'id';
comment on column table1.display_name is '表示名';
comment on column table1.sex is '性別';
comment on column table1.birthday is '誕生日';
comment on column table1.age is '年齢';
comment on column table1.married is '既婚';
comment on column table1.rate is 'レート';
comment on column table1.salary is '給料';

