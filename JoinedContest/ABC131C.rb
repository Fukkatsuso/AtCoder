def gcd(a, b)
    r = a % b
    while r != 0 do
        a = b
        b = r
        r = a % b
    end
    return b
end

a, b, c, d = gets.chomp.split(" ").map(&:to_i)
e = c * d / gcd(c, d)

can_c = b/c - (a-1)/c
can_d = b/d - (a-1)/d
can_e = b/e - (a-1)/e

cannot = (b - a + 1) - (can_c + can_d - can_e)
puts cannot