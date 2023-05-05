CREATE OR REPLACE FUNCTION admin_change_user(_id INTEGER, js json)
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
		-- do nothing
	ELSE
		RAISE EXCEPTION 'пользователь не является администратором';
	END IF;

	UPDATE users SET username = js->>'username', tele_id = (js->>'tele_id')::BIGINT, email = js->>'email', role = js->>'role' WHERE id = (js->>'id')::INTEGER;
	RETURN TRUE;
END;
$function$