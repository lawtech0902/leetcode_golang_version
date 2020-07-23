# 1.抵制一切花里胡哨
select max(Salary) as SecondHighestSalary
from Employee
where Salary < (select max(Salary) from Employee);

# 2.使用子查询和limit子句
select (
           select distinct Salary
           from Employee
           order by Salary desc
           limit 1 offset 1)
           as SecondHighestSalary;

# 3.使用ifnull和limit子句
select ifnull(
               (select distinct Salary
                from Employee
                order by Salary desc
                limit 1 offset 1), null)
           as SecondHighestSalary;