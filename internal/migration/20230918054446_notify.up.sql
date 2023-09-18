SET statement_timeout = 0;

--bun:split

CREATE FUNCTION recipes_after_update_trigger()
    RETURNS TRIGGER AS $$
BEGIN
  PERFORM pg_notify('recipes:updated', NEW.id::text);
RETURN NULL;
END;
$$
LANGUAGE plpgsql;

--bun:split

CREATE TRIGGER recipes_after_update_trigger
    AFTER UPDATE ON recipes
    FOR EACH ROW EXECUTE PROCEDURE recipes_after_update_trigger();
