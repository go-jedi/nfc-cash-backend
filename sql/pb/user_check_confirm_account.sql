CREATE OR REPLACE FUNCTION user_check_confirm_account(js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE username = js->>'username'
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким username не существует';
	END IF;

	IF _u.is_confirm_account = TRUE THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$
