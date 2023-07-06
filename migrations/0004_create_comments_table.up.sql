IF NOT EXISTS (SELECT *
FROM sys.tables
WHERE name = 'comments')
CREATE TABLE comments
(
    ID uniqueidentifier PRIMARY KEY,
    Slug nvarchar(max),
    Author nvarchar(max),
    Body nvarchar(max)
);
