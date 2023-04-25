CREATE OR REPLACE FUNCTION user_create(js json, _uid character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_eml users;
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE email = js->>'email'
	INTO _eml;

	SELECT *
	FROM users
	WHERE username = js->>'username'
	INTO _u;

	IF _eml ISNULL THEN
		-- do nothing
	ELSE
		RAISE EXCEPTION 'пользователь с таким email уже существует';
	END IF;
	
	IF _u.id ISNULL THEN
		INSERT INTO users(uid, username, email, password) VALUES(_UID, js->>'username', js->>'email', js->>'password');
	ELSE
		RAISE EXCEPTION 'пользователь с таким именем уже существует';
	END IF;
END;
$function$