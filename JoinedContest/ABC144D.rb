a, b, x = gets.chomp.split(" ").map(&:to_i)
h = x.to_f / (a * a).to_f
 
if h > b / 2
  rad = Math.atan2(2.0*(b-h).to_f, a)
  puts rad * 180.0 / Math::PI
else
  rad = Math.atan2(b, 2.0*(a*h).to_f/b.to_f)
  puts rad * 180.0 / Math::PI
end
