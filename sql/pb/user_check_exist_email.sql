CREATE OR REPLACE FUNCTION user_check_exist_email(js json)
	RETURNS boolean
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
		RETURN FALSE;
	ELSE
		RETURN TRUE;
	END IF;
END;
$function$ 