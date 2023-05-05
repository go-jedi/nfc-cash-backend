CREATE OR REPLACE FUNCTION bot_create(_id INTEGER, js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_tb telegram_bots;
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

	SELECT *
	FROM telegram_bots
	WHERE token = js->>'token'
	INTO _tb;

	IF _tb.id ISNULL THEN
		INSERT INTO telegram_bots(name, token, chat_id) VALUES(js->>'name', js->>'token', js->>'chat_id');
		RETURN TRUE;
	END IF;

	IF _tb.token = js->>'token' THEN
		UPDATE telegram_bots SET name = js->>'name', chat_id = js->>'chat_id', is_work = TRUE, is_deleted = FALSE WHERE token = js->>'token';
		RETURN TRUE;
	END IF;
END;
$function$