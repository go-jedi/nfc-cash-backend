CREATE OR REPLACE FUNCTION bot_delete(js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_tb telegram_bots;
BEGIN
	SELECT *
	FROM users
	WHERE uid = js->>'uid'
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким uid не существует';
	END IF;

	IF _u.role = 'admin' OR _u.role = 'super-admin' THEN
		-- do nothing
	ELSE
		RAISE EXCEPTION 'пользователь не является администратором';
	END IF;

	SELECT *
	FROM telegram_bots
	WHERE token = js->>'token'
	INTO _tb;

	IF _tb.id ISNULL THEN
		RAISE EXCEPTION 'бота с таким токеном не существует';
	END IF;

	UPDATE telegram_bots SET is_work = FALSE WHERE token = js->>'token';
	RETURN TRUE;
END;
$function$