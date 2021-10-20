set map

while read line
    set search (math "2020 - $line")
    if test -n "$map[$search]"
        echo "Nums: $line, $search"
        echo (math "$search * $line")
    else
        set map[$line] 1
    end

end

