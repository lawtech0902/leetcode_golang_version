# 1.自连接
select distinct p1.Email Email
from Person p1
         join Person p2
              on p1.Id != p2.Id
where p1.Email = p2.Email;

# 2.group by和临时表
select Email
from (
         select Email, count(Email) num
         from Person
         group by Email
     ) as statistic
where num > 1;

# 3.group by和having
select Email
from Person
group by Email
having count(Email) > 1;