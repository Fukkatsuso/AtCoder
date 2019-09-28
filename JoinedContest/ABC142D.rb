# TLE

a, b = gets.chomp.split(" ").map(&:to_i)
ans = 1
g = a.gcd(b)
gcdab = g
d = 2
dtmp = 1

while g > 1 do
  while g % d == 0 do
    if dtmp != d
      ans += 1 
      dtmp = d
    end
    g /= d
  end
  d += 1
  break if d > gcdab
end

puts ans