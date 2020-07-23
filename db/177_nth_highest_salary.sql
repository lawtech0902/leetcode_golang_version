# 1.单表查询
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
    SET
        N := N - 1;
    RETURN (
        SELECT salary
        FROM employee
        GROUP BY salary
        ORDER BY salary DESC
        LIMIT N, 1
    );
END