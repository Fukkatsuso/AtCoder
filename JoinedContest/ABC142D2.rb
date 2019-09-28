# AC

require 'prime'

a, b = gets.chomp.split(" ").map(&:to_i)
ans = 1
g = a.gcd(b)

ans += Prime.prime_division(g).length

puts ans