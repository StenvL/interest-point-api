CREATE FUNCTION get_distance(
	lon DECIMAL(9, 6),
    lat DECIMAL(9, 6),
	curLon DECIMAL(9, 6),
    curLat DECIMAL(9, 6)
)
RETURNS INT DETERMINISTIC
RETURN 6371 * 
    ACOS( 
        COS(RADIANS(curLat)) * 
        COS(RADIANS(lat)) * 
        COS(RADIANS(lon) - RADIANS(curLon)) + 
        SIN(RADIANS(curLat)) * 
        SIN(RADIANS(lat)) 
    )