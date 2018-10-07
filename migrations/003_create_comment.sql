-- +migrate Up
drop table if exists comment;
create table comment(
    id int(10) not null auto_increment,
    article_id int(10) default null,
    nickname varchar(15) default null,
    url varchar(255) default null,
    content text,
    created int(10) default 0,
    status tinyint(1) default 1 comment '0屏蔽，1正常',
    primary key(id)
)engine=InnoDB auto_increment=17 default charset=utf8;
-- +migrate Down 
drop table comment;