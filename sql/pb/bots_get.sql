CREATE OR REPLACE FUNCTION bots_get(_id INTEGER)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_response JSONB;
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

	SELECT
		COALESCE(bg.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT tb.id, tb.name, tb.token, tb.chat_id, tb.is_admin, tb.is_work, tb.created
			FROM telegram_bots tb
			WHERE tb.is_deleted = FALSE
		) ag
	) bg
	INTO _response;

	RETURN _response;
END;
$function$