
explain select city from ( select range_end, city from ipaddress where range_start <= 841171341 order by range_start desc limit 1 ) subqry where 841171341 <= range_end;
