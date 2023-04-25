CREATE OR REPLACE FUNCTION user_verify_email(_uid character varying)
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
		RAISE EXCEPTION 'пользователь с таким uid не существует';
	END IF;

	IF _u.is_verify_email = TRUE THEN
		RAISE EXCEPTION 'верификация по почте уже пройдена';
	END IF;

	UPDATE users SET is_verify_email = TRUE WHERE uid = _uid;
END;
$function$ 