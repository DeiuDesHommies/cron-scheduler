# cron-scheduler
基于robfig/cron/v3的定时任务定时器

# 数据库表创建命令
`
create table backup_schedule
(
	id                      bigint auto_increment comment '''主键'''
	primary key,
	task_id                 bigint      default 0  not null comment '''计划id''',
	cron                    varchar(32) default '' not null comment '''cron表达式''',
	type                    varchar(32) default '' not null comment '''任务类型''',
	next_schedule_task_time bigint      default 0  not null comment '''下次计划定份任务时间''',
	cuser                   bigint      default 0  not null comment '''创建用户uid''',
	ctime                   bigint      default 0  not null comment '''计划创建时间''',
	utime                   bigint      default 0  not null comment '''最近修改时间''',
	dtime                   bigint      default 0  not null comment '''删除时间'''
);
`
