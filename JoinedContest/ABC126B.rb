s = gets.chomp.split("")
a = s[0..1].join.to_i
b = s[2..3].join.to_i

if 1 <= a && a <= 12
    if 1 <= b && b <= 12
        @ans = "AMBIGUOUS"
    else
        @ans = "MMYY"
    end
else
    if 1 <= b && b <= 12
        @ans = "YYMM"
    else
        @ans = "NA"
    end    
end
puts @ans