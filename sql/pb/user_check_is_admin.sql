CREATE OR REPLACE FUNCTION user_check_is_admin(_id INTEGER)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE id = _id
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким id не существует';
	END IF;

	IF _u.role = 'admin' OR _u.role = 'super-admin' THEN
		RETURN TRUE;
	ELSE
		RETURN FALSE;
	END IF;
END;
$function$
