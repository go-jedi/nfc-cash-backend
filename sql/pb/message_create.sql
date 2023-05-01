CREATE OR REPLACE FUNCTION message_create(js json)
	RETURNS boolean
	LANGUAGE plpgsql
AS $function$
DECLARE
	_r rooms;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = js->>'uidRoom'
	INTO _r;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid уже существует';
	END IF;

	INSERT INTO messages(uid_room, uid_user, message) VALUES(js->>'uidRoom', js->>'uidUser', js->>'message');
	RETURN TRUE;
END;
$function$