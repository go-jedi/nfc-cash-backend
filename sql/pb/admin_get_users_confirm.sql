CREATE OR REPLACE FUNCTION admin_get_users_confirm(_id INTEGER)
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
		COALESCE(aguc.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.username, u.email, u.role
			FROM users u
			WHERE u.role != 'super-admin'
			AND u.is_confirm_account = TRUE
		) ag
	) aguc
	INTO _response;

	RETURN _response;
END;
$function$