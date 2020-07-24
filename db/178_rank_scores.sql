# 1.窗口函数
select Score,
       dense_rank() over (order by Score desc) as `Rank`
from Scores;

# 2.普通解法
select s.Score, count(distinct t.Score) `Rank`
from Scores s join Scores t
on t.Score >= s.Score
group by s.Id
order by s.Score desc;