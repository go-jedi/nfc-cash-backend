CREATE OR REPLACE FUNCTION bots_get_hidden()
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_response JSONB;
BEGIN
	SELECT
		COALESCE(bg.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM (
			SELECT tb.id, tb.name, tb.token, tb.chat_id, tb.is_admin, tb.is_work, tb.created
			FROM telegram_bots tb
			WHERE tb.is_work = TRUE
		) ag
	) bg
	INTO _response;

	RETURN _response;
END;
$function$