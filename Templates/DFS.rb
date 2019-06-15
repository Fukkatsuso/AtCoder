class Node
    attr_reader :data, :nodes, :is_visited

    def initialize(data, *nodes)
        @data = data
        @nodes = nodes
        @is_visited = false
    end

    def visit
        @is_visited = true
    end
end


def dfs(start, target)
    stack = []
    stack.push(start)
    until stack.empty?
        new_node = stack.pop
        if new_node.data == target.data
        end
        unless new_node.is_visited
            new_node.visit
            new_node.nodes.each do |n|
                stack.push(n) unless n.is_visited
            end
        end
    end
    return route
end


n1 = Node.new(1)
n2 = Node.new(2)