
create table ipaddress(
id int not null auto_increment,
range_start int(11),
range_end int(11),
country_code varchar(4),
country varchar(255),
state varchar(255),
city varchar(255),
primary key(id));
