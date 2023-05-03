CREATE OR REPLACE FUNCTION ip_check_block(_ip character varying)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ib ip_blocks;
BEGIN
	SELECT *
	FROM ip_blocks
	WHERE address = _ip
	INTO _ib;

	IF _ib.id ISNULL THEN
		RETURN FALSE;
	END IF;

	IF _ib.is_block = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$