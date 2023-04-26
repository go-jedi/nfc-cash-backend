CREATE OR REPLACE FUNCTION user_compare_password(_uid character varying, _pswd character varying)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE uid = _uid
	INTO _u;

	IF _u.password = _pswd THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$