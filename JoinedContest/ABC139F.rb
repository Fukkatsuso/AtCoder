# WA

def distance(x, y)
  return ((x ** 2) + (y ** 2)) ** 0.5
end

n = gets.chomp.to_i
xy = []
xsum = 0
ysum = 0
n.times do
  x, y = gets.chomp.split(" ").map(&:to_i)
  xsum += x
  ysum += y
  xy << [x, y]
end

dist = distance(xsum, ysum)
d = [[0, 0]] * 4
dist = [0] * 4

(0..n-1).each do |i|
  x = xy[i][0]
  y = xy[i][1]
  if x > 0 && y > 0
    d[0][0] += x
    d[0][1] += y
  elsif x < 0 && y > 0
    d[1][0] += x
    d[1][0] += y
  elsif x < 0 && y < 0
    d[2][0] += x
    d[2][1] += y
  elsif x > 0 && y < 0
    d[3][0] += x
    d[3][1] += y
  else
    if x == 0 && y > 0
      d[0][0] += x
      d[0][1] += y
      d[1][0] += x
      d[1][0] += y
    elsif x < 0 && y == 0
      d[1][0] += x
      d[1][0] += y
      d[2][0] += x
      d[2][1] += y
    elsif x == 0 && y < 0
      d[2][0] += x
      d[2][1] += y
      d[3][0] += x
      d[3][1] += y
    elsif x > 0 && y == 0
      d[3][0] += x
      d[3][1] += y
      d[0][0] += x
      d[0][1] += y
    end
  end
end

(0..3).each do |i|
  dist[i] = distance(d[i][0], d[i][1])
end

puts dist.max