CREATE OR REPLACE FUNCTION user_launch_recovery_password(_uid character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE uid = _uid
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким uid не найден';
	ELSE
		UPDATE users SET is_change_password = TRUE WHERE uid = _uid;
	END IF;
END;
$function$