# Sql说明
## turingcup.users 
- id 主键，唯一标识用户，设置为auto_increment
- name 用户姓名，约束为非空
- password 用户密码，约束为非空
- phone 用户手机号码
- email 用户邮箱，约束为邮箱格式
- school 用户所在学校中文名，约束为非空
- school_id 用户所在学校的id标识，即，教育部为高校编排的院校代码
- created_at 表项的创建时间，即，每一位用户的注册时间
- updated_at 表项的最后更改时间
- deleted_at 表项的删除时间

补充说明：该表项在用户注册时创建，并在后续用户信息更改或删除时更新数据。

## turingcup.races
- id 主键，唯一标识每届图灵杯项目，以本届图灵杯为第一届，设置为auto_increment
- name 每届图灵杯的主题，设置为非空
- created_at 表项的创建时间
- updated_at 表项的最后更改时间
- deleted_at 表项的删除时间

补充说明：为实现图灵杯项目的延拓性设计的数据表，用以支持在当前数据库中实现后续若干届图灵杯项目。

## turingcup.teams
- id 主键，唯一标识每支队伍，设置auto_increment
- rid 队伍所属的图灵杯的turingcup.races表项的id。如，队伍A为2023年第一届图灵杯的队伍，队伍B为2024年第二届图灵杯的队伍。设置为非空
- name 队伍名称，设置为非空
- cap_id 队长的turingcup.users表项的id
- created_at 表项的创建时间
- updated_at 表项的最后更改时间
- deleted_at 表项的删除时间


补充说明：用于存储重要的team信息。

## turingcup.team_records
- id 主键，唯一标识每一支队伍记录，设置auto_increment
- race_id 队伍记录所属的图灵杯的turingcup.races表项的id
- uid 队伍记录对应的用户的turingcup.users表项的id
- tid 队伍记录所属的队伍的turingcup.teams表项的id

补充说明：使用此数据表实现turingcup.races、turingcup.teams、turingcup.users三个数据表的联系，实现了解耦合的效果。

## turingcup.rounds
- id 主键，唯一标识每一轮对战，设置auto_increment
- tid1 参与对战的第一支队伍，设置为非空
- tid2 参与对战的第二支队伍，设置为非空
- tid_win 获胜队伍
- round_type 对战的类型
- record_path 对战视频的存储路径


补充说明：用于存储对战信息，仅适用于本届图灵杯。在round_type中根据不同的对战类型进行分类。根据本次图灵杯的赛程设计，具体对战类型有：1、初赛正常赛 2.初赛车轮战 3、复活赛正常赛 4.复活赛车轮战 5、决赛车轮战。

## circle_race_infos
- id 主键，唯一标识每一轮车轮战，设置auto_increment
- tid 参与该车轮战的一支队伍的turingcup.teams表项的id
- score 该队伍在本次车轮战中的得分

补充说明：用于存储每一轮车轮战的信息，仅适用于本届图灵杯，作为对普通对战的补充，用于在人数较少时，以更高的公平性决出胜者。每一轮车轮战完成后，根据每一个team的score得分排名。

