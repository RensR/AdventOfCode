namespace AdventOfCode.Solutions.Year2020
{
    class Day17 : ASolution
    {
        bool[,,] galaxy;
        bool[,,,] galaxy4D;
        const int maxTime = 6;
        readonly int size;

        public Day17() : base(17, 2020, "")
        {
            var rows = Input.SplitByNewline();
            size = rows[0].Length + 2 * maxTime;
            galaxy = new bool[size, size, size];
            galaxy4D = new bool[size, size, size, size];

            for (int i = 0; i < rows.Count; i++)
            {
                for (int j = 0; j < rows[0].Length; j++)
                {
                    galaxy[i + maxTime, j + maxTime, maxTime] = rows[i][j] == '#';
                    galaxy4D[i + maxTime, j + maxTime, maxTime, maxTime] = rows[i][j] == '#';
                }
            }

        }

        protected override string SolvePartOne()
        {
            for(var time = 0; time < maxTime; time++)
            {
                Update3D();
            }

            var count = 0;
            for (var x = 0; x < galaxy.GetLength(0); x++)
            {
                for (var y = 0; y < galaxy.GetLength(1); y++)
                {
                    for (var z = 0; z < galaxy.GetLength(2); z++)
                    {
                        if (galaxy[x, y, z])
                            count++;
                    }
                }
            }

            return count.ToString();
        }

        void Update3D()
        {
            var newGalaxy = new bool[size, size, size];
            for (var x = 0; x < galaxy.GetLength(0); x++)
            {
                for (var y = 0; y < galaxy.GetLength(1); y++)
                {
                    for (var z = 0; z < galaxy.GetLength(2); z++)
                    {
                        var neighbours = CountNeighbours3D(x, y, z);
                        if (galaxy[x, y, z])
                        {
                            newGalaxy[x, y, z] = neighbours is 2 or 3;
                        }
                        else
                        {
                            newGalaxy[x, y, z] = neighbours is 3;
                        }
                    }
                }
            }

            galaxy = newGalaxy;
        }

        int CountNeighbours3D(int x, int y, int z)
        {
            var neighbours = 0;
            for (var dx = -1; dx <= 1; dx++)
            {
                if (x + dx < 0 || x + dx >= galaxy.GetLength(0))
                    continue;
                for (var dy = -1; dy <= 1; dy++)
                {
                    if (y + dy < 0 || y + dy >= galaxy.GetLength(1))
                        continue;
                    for (var dz = -1; dz <= 1; dz++)
                    {
                        if (z + dz < 0 || z + dz >= galaxy.GetLength(2))
                            continue;
                        if (dx == 0 && dy == 0 && dz == 0)
                            continue;

                        if (galaxy[x + dx, y + dy, z + dz])
                            neighbours++;
                    }
                }
            }

            return neighbours;
        }

        void Update4D()
        {
            var newGalaxy4D = new bool[size, size, size, size];
            for (var x = 0; x < galaxy4D.GetLength(0); x++)
            {
                for (var y = 0; y < galaxy4D.GetLength(1); y++)
                {
                    for (var z = 0; z < galaxy4D.GetLength(2); z++)
                    {
                        for (var w = 0; w < galaxy4D.GetLength(3); w++)
                        {
                            var neighbours = CountNeighbours4D(x, y, z, w);
                            if (galaxy4D[x, y, z, w])
                            {
                                newGalaxy4D[x, y, z, w] = neighbours is 2 or 3;
                            }
                            else
                            {
                                newGalaxy4D[x, y, z, w] = neighbours is 3;
                            }
                        }
                    }
                }
            }

            galaxy4D = newGalaxy4D;
        }

        int CountNeighbours4D(int x, int y, int z, int w)
        {
            var neighbours = 0;
            for(var dx = -1; dx <= 1; dx++)
            {
                if (x + dx < 0 || x + dx >= galaxy4D.GetLength(0))
                    continue;
                for (var dy = -1; dy <= 1; dy++)
                {
                    if (y + dy < 0 || y + dy >= galaxy4D.GetLength(1))
                        continue;
                    for (var dz = -1; dz <= 1; dz++)
                    {
                        if (z + dz < 0 || z + dz >= galaxy4D.GetLength(2))
                            continue;

                        for (var dw = -1; dw <= 1; dw++)
                        {
                            if (w + dw < 0 || w + dw >= galaxy4D.GetLength(3))
                                continue;

                            if (dx == 0 && dy == 0 && dz == 0 && dw == 0)
                                continue;

                            if (galaxy4D[x + dx, y + dy, z + dz, w + dw])
                                neighbours++;
                        }
                    }
                }
            }

            return neighbours;
        }

        protected override string SolvePartTwo()
        {
            for (var time = 0; time < maxTime; time++)
            {
                Update4D();
            }

            var count = 0;
            for (var x = 0; x < galaxy4D.GetLength(0); x++)
            {
                for (var y = 0; y < galaxy4D.GetLength(1); y++)
                {
                    for (var z = 0; z < galaxy4D.GetLength(2); z++)
                    {
                        for (var w = 0; w < galaxy4D.GetLength(3); w++)
                        {
                            if (galaxy4D[x, y, z, w])
                                count++;
                        }
                    }
                }
            }

            return count.ToString();
        }
    }
}
