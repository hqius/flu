create table role
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(50)  not null comment '角色名',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '角色';

create table user
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(50)  not null comment '用户名',
    password    VARCHAR(50)  not null comment '密码',
    email       VARCHAR(100) not null comment '邮箱',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户表';

create table role_permission
(
    id            int primary key auto_increment comment '主键ID',
    role_id       int       not null comment '角色',
    permission_id int       not null comment '权限',
    update_time   timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time   timestamp not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '角色-权限';

create table permission
(
    id          int primary key auto_increment comment '主键ID',
    type        tinyint     not null default 0 comment '类型',
    name        varchar(64) not null default '' comment '权限名',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '权限表';

create table position
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(64) not null default '' comment '职位名',
    did         int         not null comment '部门ID',
    pid         int         not null comment '上级职位ID',
    update_time timestamp   not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp   not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '职位表';

create table position_role
(
    id          int primary key auto_increment comment '主键ID',
    rid         int       not null comment '角色ID',
    pid         int       not null comment '权限ID',
    update_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '职位-角色表';

create table user_position
(
    id          int primary key auto_increment comment '主键ID',
    uid         int       not null comment '用户ID',
    pid         int       not null comment '职位ID',
    update_time timestamp not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户-职位表';

create table department
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(64)  not null comment '部门名',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '部门表';

create table usergroup
(
    id          int primary key auto_increment comment '主键ID',
    name        varchar(64)  not null comment '用户组',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户组表';

create table user_usergroup
(
    id          int primary key auto_increment comment '主键ID',
    uid         int       not null comment '用户ID',
    gid         int       not null comment '组ID',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户-用户组表';

create table usergroup_role
(
    id          int primary key auto_increment comment '主键ID',
    uid         int       not null comment '用户ID',
    rid         int       not null comment '角色ID',
    update_time timestamp    not null default current_timestamp on update current_timestamp comment '更新时间',
    create_time timestamp    not null default current_timestamp comment '创建时间'
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '用户-角色表';


