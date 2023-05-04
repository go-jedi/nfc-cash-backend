CREATE OR REPLACE FUNCTION bots_get(_uid character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
	_response JSONB;
BEGIN
	SELECT *
	FROM users
	WHERE uid = _uid
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким uid не существует';
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
			WHERE tb.is_work = TRUE
		) ag
	) bg
	INTO _response;

	RETURN _response;
END;
$function$