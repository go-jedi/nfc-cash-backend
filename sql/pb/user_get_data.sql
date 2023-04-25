CREATE OR REPLACE FUNCTION user_get_data(_username character varying, _password character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugd.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.uid, u.username, u.password, u.role
			FROM users u
			WHERE u.username = _username AND u.password = _password
		) ag
	) ugd
	INTO _response;

	RETURN _response;
END;
$function$