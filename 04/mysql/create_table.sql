-- 
drop table if exists table1 cascade;

create table table1 (
  id INT not null AUTO_INCREMENT comment 'id'
  , display_name VARCHAR(100) comment '表示名'
  , sex CHAR(1) not null comment '性別'
  , birthday DATETIME comment '誕生日'
  , age INT comment '年齢'
  , married boolean comment '既婚'
  , rate FLOAT comment 'レート'
  , salary DECIMAL(10,0) comment '給料'
  , constraint table1_PKC primary key (id)
) comment '' ;