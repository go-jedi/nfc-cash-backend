CREATE OR REPLACE FUNCTION room_create(_uid character varying)
	RETURNS void
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = _uid
	INTO _r;

	IF _r.id ISNULL THEN
		INSERT INTO rooms(uid_room) VALUES(_uid);
	ELSE
		RAISE EXCEPTION 'комната с таким uid уже существует';
	END IF;
END;
$function$