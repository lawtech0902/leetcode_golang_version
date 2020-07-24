select distinct a.Num ConsecutiveNums
from Logs a, Logs b, Logs c
where
a.Id = b.Id + 1 and a.Id = c.Id + 2 and a.Num = b.Num and a.Num = c.Num;