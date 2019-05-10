mass = gets.chomp.split(" ").map(&:to_i)
height = mass[0]
width = mass[1]
grid = []
(0..height-1).each do
    grid.push(gets.chomp.split(""))
end

can = true
(0..height-1).each do |h|
    (0..width-1).each do |w|
        next if grid[h][w] == "." #白なら無視
        #周囲のマスに黒があるか確認
        next if h > 0 && grid[h-1][w] == "#" #上のマス
        next if h < height-1 && grid[h+1][w] == "#" #下のマス
        next if w > 0 && grid[h][w-1] == "#" #左のマス
        next if w < width-1 && grid[h][w+1] == "#" #右のマス
        can = false
    end
    
end

if can
    puts "Yes"
else
    puts "No"
end