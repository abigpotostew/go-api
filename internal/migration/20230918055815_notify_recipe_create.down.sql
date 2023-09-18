SET statement_timeout = 0;

--bun:split

DROP TRIGGER recipes_after_create_trigger ON recipes;

--bun:split

DROP FUNCTION recipes_after_create_trigger;