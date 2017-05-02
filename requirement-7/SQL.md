# Requirement-7 Solution

Please note that i have used MySQL instead of SQL Server, as i'm running a MacOS operating system

###First answer:

    SELECT DISTINCT
        s.Name
    FROM
        `Orders` AS o
            INNER JOIN
        `Salesperson` AS s ON o.SalespersonID = s.SalespersonID
            INNER JOIN
        `Customer` AS c ON o.CustomerID = c.CustomerID
    WHERE
        c.Name = 'George';

###Second answer:

    SELECT DISTINCT s1.Name
    FROM `Salesperson` AS s1
    WHERE s1.SalespersonID NOT IN (
      SELECT s.SalespersonID
      FROM `Orders` AS o
      INNER JOIN `Salesperson` AS s
        ON o.SalespersonID = s.SalespersonID
      INNER JOIN `Customer` AS c
        ON o.CustomerID = c.CustomerID
      WHERE c.Name = 'George'
    );

###Third answer:

    SELECT
        s1.Name
    FROM
        `Salesperson` AS s1
    WHERE
        s1.SalespersonID IN (
          SELECT
              o.SalespersonID
          FROM
              `Orders` AS o
          GROUP BY o.SalespersonID
          HAVING COUNT(*) > 1
        );
