# Requirement-8 Solution

Please note that i have used MySQL instead of SQL Server, as i'm running a MacOS operating system

### First answer:

    SELECT
      *
    FROM
      Salesperson
    ORDER BY Salary DESC
    LIMIT 2 , 1;

The above query will display one row starting from the 2nd position, which means
it'll return the 3rd highest earning person because the rows are ordered by Salary
in a descending order

### Second answer:

    CREATE VIEW BigOrders AS
    SELECT
      o.CustomerID, SUM(o.CostOfUnit) AS 'TotalOrderValue'
    FROM
      `Orders` AS o
    GROUP BY
      o.CustomerID
    HAVING TotalOrderValue>1000

I chose to create a view as the summary table because SQL will always return the updated values in a view

### Third answer:

    SELECT
        MONTH(o.OrderDate) AS 'Month',
        YEAR(o.OrderDate) AS 'Year',
        SUM(o.CostOfUnit) AS 'Total'
    FROM
        `Orders` AS o
    GROUP BY MONTH(o.OrderDate) , YEAR(o.OrderDate)
    ORDER BY `Year` , `Month`;
