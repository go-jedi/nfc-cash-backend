CREATE OR REPLACE FUNCTION user_check_recovery_password(_uid character varying)
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

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким uid уже существует';
	END IF;

	IF _u.is_change_password = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$