class Graph():
    def __init__(self, inputGraph) -> None:
        self.inputGraph = inputGraph
        self.length = len(inputGraph)
        self.width = len(inputGraph[0])
        self.minCost = float('inf')
        self.S = [0, 0]
        self.E = [0, 0]
        self.found = False
        # self.adjacencyLists = {}
        self.adjacencyLists = self.createAdjacencyLists()
        pass

    def getVertexKey(self, i, j):
        return str(i) + '_' + str(j)

    def checkSAdjacent(self, i, j, k, l):
        if self.inputGraph[i][j] == 'z' and self.inputGraph[k][l] == 'E':
            return True
        if self.inputGraph[i][j] == 'S' and self.inputGraph[k][l] == 'a':
            return True
        return False

    def createAdjacencyLists(self) -> None:
        adjacencyLists = {}

        for i in range(self.length):
            for j in range(self.width):
                key = self.getVertexKey(i, j)
                adjacencyLists[key] = []

                if self.inputGraph[i][j] == 'S':
                    self.S = [i, j]

                if self.inputGraph[i][j] == 'E':
                    self.E = [i, j]
                    continue

                # Look down
                if (i + 1 < self.length) and ((abs(ord(self.inputGraph[i+1][j]) - ord(self.inputGraph[i][j])) <= 1) or self.checkSAdjacent(i, j, i+1, j)):
                    adjacencyLists[key].append(self.getVertexKey(i+1, j))

                # Look up
                if (i - 1 >= 0) and ((abs(ord(self.inputGraph[i-1][j]) - ord(self.inputGraph[i][j])) <= 1) or self.checkSAdjacent(i, j, i-1, j)):
                    adjacencyLists[key].append(self.getVertexKey(i-1, j))

                # Look right
                if (j + 1 < self.width) and ((abs(ord(self.inputGraph[i][j+1]) - ord(self.inputGraph[i][j])) <= 1) or self.checkSAdjacent(i, j, i, j+1)):
                    adjacencyLists[key].append(self.getVertexKey(i, j+1))

                # Look down
                if (j - 1 < self.width) and ((abs(ord(self.inputGraph[i][j-1]) - ord(self.inputGraph[i][j])) <= 1) or self.checkSAdjacent(i, j, i, j-1)):
                    adjacencyLists[key].append(self.getVertexKey(i, j-1))

        return adjacencyLists

    def search(self, node1, node2, cost):

        if node1 == node2:
            print('Found')
            self.found = True
            self.minCost = min(self.minCost, cost)
            return

        for node in self.adjacencyLists[node1]:
            self.search(node, node2, cost + 1)
        pass

def generateGraph(inputGraph):
    # file1 = open('dec12input.txt', 'r')
    file1 = open('sample.in', 'r')
    lines = file1.readlines()
    for line in lines:
        inputGraph.append([*line.strip()])

if __name__ == '__main__':
    inputGraph = []
    generateGraph(inputGraph)
    g = Graph(inputGraph)
    g.search(g.getVertexKey(g.S[0], g.S[1]), g.getVertexKey(g.E[0], g.E[1]), 0)
