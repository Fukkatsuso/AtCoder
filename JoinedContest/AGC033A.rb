line = gets.chomp.split(" ").map(&:to_i)
height = line[0]
width = line[1]
a = []
(0..height-1).each do
    a.push(gets.chomp.split(""))
end

b = Marshal.load(Marshal.dump(a))
cnt = 0
(0..height*width).each do
    cnt_white = 0
    (0..height-1).each do |h|
        (0..width-1).each do |w|
            # if a[h][w] == "#"
            #     cnt_black += 1
            #     b[h-1][w] = "#" if h > 0
            #     b[h][w-1] = "#" if w > 0
            #     b[h+1][w] = "#" if h + 1 < height
            #     b[h][w+1] = "#" if w + 1 < width
            # end
            if b[h][w] == "."
                if h > 0 && a[h-1][w] == "#"
                    b[h][w] = "#"
                elsif w > 0 && a[h][w-1] == "#"
                    b[h][w] = "#"
                elsif h + 1 < height && a[h+1][w] == "#"
                    b[h][w] = "#"
                elsif w + 1 < width && a[h][w+1] == "#"
                    b[h][w] = "#"
                end
                cnt_white += 1
            end 
        end
    end
    break if cnt_white == 0
    a = Marshal.load(Marshal.dump(b))
    cnt += 1
end

puts cnt