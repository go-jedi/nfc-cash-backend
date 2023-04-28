CREATE OR REPLACE FUNCTION user_get_profile(_id INTEGER)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(ugp.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT u.id, u.username, u.email, u.role
			FROM users u
		) ag
	) ugp
	INTO _response;

	RETURN _response;
END;
$function$