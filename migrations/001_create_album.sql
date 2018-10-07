-- +migrate Up
drop table if exists album;
create table album(
    id int(10) not null auto_increment,
    title varchar(255) not null default '' comment '文章标题',
    picture varchar(255) default '' comment 'picture',
    keywords varchar(2550) default '' comment '关键词',
    summary varchar(255) default '',
    created int(10) default 0 comment '发布时间',
    viewnum int(10) default 0 comment '阅读次数',
    status tinyint(1) default 1 comment '状态: 0草稿，1已发布',
    primary key(id),
    key INDEX_TCVS (title,created,viewnum,status) using btree
) engine=InnoDB auto_increment=13 default charset=utf8 comment='相册';

-- +migrate Down
drop table album;