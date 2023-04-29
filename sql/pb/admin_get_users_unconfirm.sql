CREATE OR REPLACE FUNCTION admin_get_users_unconfirm(_id INTEGER)
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

	IF _u.role = 'user' THEN
		RAISE EXCEPTION 'пользователь не является администратором';
	END IF;

	SELECT
		COALESCE(aguu.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.username, u.email, u.role
			FROM users u
			WHERE u.is_confirm_account = FALSE
		) ag
	) aguu
	INTO _response;

	RETURN _response;
END;
$function$