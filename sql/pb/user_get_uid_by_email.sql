CREATE OR REPLACE FUNCTION user_get_uid_by_email(js json)
	RETURNS text
	LANGUAGE plpgsql
AS $function$
DECLARE
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE email = js->>'email'
	INTO _u;

	IF _u.id ISNULL THEN
		RAISE EXCEPTION 'пользователь с таким email не существует';
	END IF;

	RETURN _u.uid;
END;
$function$