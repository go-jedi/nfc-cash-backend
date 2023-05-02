CREATE OR REPLACE FUNCTION room_get(_uidr character varying)
	RETURNS json
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
	_response JSONB;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = _uidr
	INTO _r;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid уже существует';
	END IF;

	SELECT
		COALESCE(rg.s, '[]')
	FROM
	(
		SELECT json_agg(ag.*)::JSONB s
		FROM
		(
			SELECT r.id, r.uid_room, r.member_count, r.is_works
			FROM rooms r
			WHERE r.uid_room = _uidr
		) ag
	) rg
	INTO _response;

	RETURN _response;
END;
$function$