DO $$
BEGIN

    IF NOT EXISTS (
    SELECT FROM pg_catalog.pg_tables
    WHERE schemaname = 'public'
        AND tablename = 'comments'
) THEN
    CREATE TABLE comments
    (
        ID uuid PRIMARY KEY,
        Slug text,
        Author text,
        Body text
    );
END IF;

END $$;
