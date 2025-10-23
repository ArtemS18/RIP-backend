DO $$
DECLARE
    i INTEGER := 1;
    max_rows INTEGER := 100000;
BEGIN
    LOOP
        INSERT INTO components(title, type, mtbf, mttr, available, description)
        VALUES ('Component ' || i, 'Type' || i, 100, 100, 1, 'Description'||i);
        i := i + 1;
        EXIT WHEN i > max_rows;
    END LOOP;
END $$;
