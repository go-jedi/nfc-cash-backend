CREATE OR REPLACE FUNCTION admin_user_confirm_account(_id INTEGER, js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_ua users;
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE id = _id
	INTO _ua;

	IF _ua.id ISNULL THEN
		RAISE EXCEPTION 'администратора с таким id не существует';
	END IF;

	IF _ua.role = 'user' THEN
		RAISE EXCEPTION 'пользователь не является администратором';
	END IF;

	SELECT *
	FROM users
	WHERE id = (js->>'id')::INTEGER
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким id не существует';
	END IF;

	UPDATE users SET is_confirm_account = TRUE WHERE id = (js->>'id')::INTEGER;
	RETURN TRUE;
END;
$function$