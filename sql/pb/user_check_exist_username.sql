CREATE OR REPLACE FUNCTION user_check_exist_username(js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE 
	_u users;
BEGIN
	SELECT *
	FROM users
	WHERE username = js->>'username'
	INTO _u;

	IF _u.id ISNULL THEN
		RETURN FALSE;
	ELSE
		RETURN TRUE;
	END IF;
END;
$function$ 