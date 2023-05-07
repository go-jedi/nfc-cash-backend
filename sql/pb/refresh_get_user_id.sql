CREATE OR REPLACE FUNCTION refresh_get_user_id(_rtkn character varying)
	RETURNS integer
	LANGUAGE plpgsql
AS $function$
DECLARE
	_rt refresh_tokens;
BEGIN
	SELECT *
	FROM refresh_tokens
	WHERE token = _rtkn
	INTO _rt;

	IF _rt.id ISNULL THEN
		RAISE EXCEPTION 'пользователя с таким токеном не существует';
	END IF;

	RETURN _rt.user_id;
END;
$function$ 