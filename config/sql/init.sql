
-- 创建数据表User
create table turingcup.users 
(
	id int auto_increment,
	name varchar(32) not null,
	password varchar(255) not null,
	phone varchar(11),
	email varchar(255) check (email like '%@%'),
	school varchar(255) not null,
	school_id varchar(255),
	created_at  datetime,
	updated_at  datetime,
	deleted_at   datetime,
	primary key(id)
);

-- 创建数据表Race
create table turingcup.races (
	id int auto_increment,
	name varchar(32) not null,
	createAt datetime,
	updateAt datetime,
	primary key(id)
);

-- 创建数据表Team
create table turingcup.teams (
	id int auto_increment,
	rid int not null,
	name varchar(32) not null,
	capID int not null,
	createAt datetime,
	updateAt datetime,
	primary key(id)
);

-- 创建数据表TeamRecord
create table turingcup.team_records (
	id int auto_increment,
	raceID int not null,
	uid int not null,
	tid int not null,
	primary key(id)
);

-- 创建数据表CircleRaceInfo
create table turingcup.circle_race_infos (
	tid int not null,
	rid int not null,
	score int,
	primary key(tid, rid)
);
-- 创建数据表Round
create table turingcup.rounds(
	id int auto_increment,
	tid1 int not null,
	tid2 int not null,
	tidWin int,
	roundType int not null,
	recordPath varchar(255),
	primary key(id)
);