-- +migrate Up
drop table if exists user;
create table user(
    id int(10) not null auto_increment,
    phone varchar(20) not null default '' comment '用户名',
    password varchar(255) not null default '' comment '密码',
    user_profile_id int(10) default null,
    created int(10) default null comment '注册时间',
    changed int(10) default null comment  '编辑时间',
    status int(4) not null default 1 comment '状态: 0屏蔽，1正常',
    primary key(id)
)engine=InnoDB auto_increment=2 default charset=utf8 comment='用户';

insert into user(phone,password,user_profile_id)
values("15888888888","PBKDF2$sha512$1000$BAge9tI5qjmfihau1YCBKA==$jC/7JLU+GgLawRGD2JUvvVoiRmQErt8QUnRLaCqUNqfYKiPp7QXPF0MsArrs43TmnH2fgRXuWDzj3YoOc0Hf9A==",1);
-- +migrate Down
drop table user;