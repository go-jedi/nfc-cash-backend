CREATE OR REPLACE FUNCTION user_recovery_password(js json)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE uid = js->>'uid'
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким uid не существует';
	END IF;

	UPDATE users SET password = js->>'password' WHERE uid = js->>'uid';
END;
$function$