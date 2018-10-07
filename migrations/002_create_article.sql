-- +migrate Up
drop table if exists article;
create table article(
    id int(10) not null auto_increment,
    title varchar(255) not null default '' comment '文章标题',
    url varchar(255) default '' comment 'URL',
    keywords varchar(2550) default '' comment '关键词',
    summary varchar(255) default '',
    content longtext not null comment '正文',
    author varchar(20) default '' comment '作者',
    created int(10) default 0 comment '发布时间',
    viewnum int(10) default 0 comment '阅读次数',
    status tinyint(1) default 1 comment '状态: 0草稿，1已发布',
    primary key(id),
    key INDEX_TCVS(title,created,viewnum,status) using btree
) engine=InnoDB auto_increment=39 default charset=utf8 comment='文章';

-- +migrate Down
drop table article;