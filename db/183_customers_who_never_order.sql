# 1.左连接
select a.Name Customers
from Customers a
         left join Orders b
                   on a.Id = b.CustomerId
where b.CustomerId is null;

# 2.子查询和not in
select c.name Customers
from Customers c
where c.Id not in (
    select CustomerId
    from Orders
);

# 3.not exists
select c.Name Customers
from Customers c
where not exists(
        select CustomerId
        from Orders o
        where c.Id = o.CustomerId
    );