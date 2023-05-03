CREATE OR REPLACE FUNCTION ip_block(_ip character varying)
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

	IF _ib.is_block = FALSE THEN
		UPDATE ip_blocks SET is_block = TRUE WHERE address = _ip;
		RETURN TRUE;
	END IF;

	IF _ib.id ISNULL THEN
		INSERT INTO ip_blocks(address) VALUES(_ip);
		RETURN TRUE;
	END IF;
END;
$function$