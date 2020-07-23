select p.FirstName, p.LastName, a.city, a.state
from Person p
left outer join address a
on p.PersonId = a.PersonId;