def count_bomb(grid, height, width, h, w)
    cnt = 0
    #上のマス
    if h > 0
        cnt += 1 if grid[h-1][w] == "#"
        if 0 < w #左のマスが存在
            cnt += 1 if grid[h-1][w-1] == "#"
        end
        if w < width-1 #右のマスが存在
            cnt += 1 if grid[h-1][w+1] == "#"
        end
    end

    #同じ段のマス
    if 0 < w #左のマスが存在
        cnt += 1 if grid[h][w-1] == "#"
    end
    if w < width-1 #右のマスが存在
        cnt += 1 if grid[h][w+1] == "#"
    end

    # 下のマス
    if h < height-1
        cnt += 1 if grid[h+1][w] == "#"
        if 0 < w #左のマスが存在
            cnt += 1 if grid[h+1][w-1] == "#"
        end
        if w < width-1 #右のマスが存在
            cnt += 1 if grid[h+1][w+1] == "#"
        end
    end

    return cnt
end

mass = gets.chomp.split(" ").map(&:to_i)
height = mass[0]
width = mass[1]
grid = []
(0..height-1).each do
    grid.push(gets.chomp.split(""))
end

(0..height-1).each do |h|
    (0..width-1).each do |w|
        if grid[h][w] == "."
            grid[h][w] = count_bomb(grid, height, width, h, w).to_s
        end
    end
end

(0..height-1).each do |h|
    puts grid[h].join
end