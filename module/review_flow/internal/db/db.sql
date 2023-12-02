create table `t_apply_config`
(
    `id`          bigint(20) unsigned not null auto_increment comment '主键ID',
    `name`        varchar(32)         not null comment '配置名',
    `desc`        varchar(128)        not null comment '描述',
    `content`     mediumtext          not null comment '配置详情',
    `operator`    bigint(20)          not null default 0 comment '操作人id',
    `is_deleted`  tinyint(1)          NOT NULL COMMENT '是否删除',
    `update_time` datetime            not null default current_timestamp on update current_timestamp comment '最后修改时间',
    `create_time` datetime            not null default current_timestamp comment '创建时间',
    primary key (`id`)
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '审批流配置表';

create table `t_apply_detail`
(
    `id`            bigint(20) unsigned not null auto_increment comment '主键ID',
    `config_id`     bigint(20) unsigned not null comment '配置id',
    `title`         varchar(32)         not null comment '标题',
    `content`       mediumtext          not null comment '内容',
    `o_bpm_content` mediumtext          not null comment '原流转信息',
    `bpm_content`   mediumtext          not null comment '流转信息',
    `status`        bigint(20) unsigned not null comment '状态: 1-审批中;2-撤销;3-审批通过;4-审批驳回',
    `is_deleted`    tinyint(1)          not null comment '是否删除',
    `update_time`   datetime            not null default current_timestamp on update current_timestamp comment '最后修改时间',
    `create_time`   datetime            not null default current_timestamp comment '创建时间',
    primary key (`id`)
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '审批明细表';

create table `t_apply_log`
(
    `id`          bigint(20) unsigned not null auto_increment comment '主键ID',
    `detail_id`   bigint(20) unsigned not null comment '审批明细ID',
    `opt_type`    tinyint(1)          not null comment '操作人类型: 1-申请人;2-审批人',
    `operator`    bigint(20)          not null comment '操作人',
    `title`       varchar(32)         not null comment '标题',
    `content`     mediumtext          not null comment '内容',
    `status`      tinyint(1)          not null comment '状态: 1-操作中;2-操作完成',
    `is_deleted`  tinyint(1)          NOT NULL COMMENT '是否删除',
    `update_time` datetime            not null default current_timestamp on update current_timestamp comment '最后修改时间',
    `create_time` datetime            not null default current_timestamp comment '创建时间',
    primary key (`id`)
) engine InnoDB
  character set utf8mb4
  collate utf8mb4_general_ci
    comment '审批流水表';