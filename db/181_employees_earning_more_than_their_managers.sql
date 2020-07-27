select E1.name as Employee
from Employee E1
left outer join Employee E2
on E1.ManagerId = E2.Id
where E1.Salary > E2.Salary;