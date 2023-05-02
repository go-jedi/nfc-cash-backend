CREATE OR REPLACE FUNCTION room_join(_uidr character varying, _uidu character varying)
	RETURNS void
	LANGUAGE plpgsql 
AS $function$
DECLARE
	_r rooms;
	_ur users_room;
BEGIN
	SELECT *
	FROM rooms
	WHERE uid_room = _uidr
	INTO _r;

	SELECT *
	FROM users_room
	WHERE uid_user = _uidu
	AND entry_room = _uidr
	INTO _ur;

	IF _r.id ISNULL THEN
		RAISE EXCEPTION 'комната с таким uid не существует';
	END IF;

	IF _ur.id ISNULL THEN
		UPDATE rooms SET member_count = member_count+1, is_works = TRUE, members = array_append(members, _uidu) WHERE uid_room = _uidr;
		INSERT INTO users_room(uid_user, entry_room) VALUES(_uidu, _uidr);
	END IF;
END;
$function$