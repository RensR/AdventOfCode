using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{
    class Day20 : ASolution
    {
        readonly List<Tile> Tiles;
        readonly List<Tile> Corners = new();

        public Day20() : base(20, 2020, "Jurassic Jigsaw")
        {
            Tiles = Input.Split("\n\n").Select(tileLines => new Tile(tileLines)).ToList();
        }

        protected override string SolvePartOne()
        {
            foreach (var tile in Tiles)
            {
                var (u, r, d, l) = tile.MatchesSides(Tiles);
                if ((l == 0 || r == 0) && (u == 0 || d == 0))
                    Corners.Add(tile);
            }
            return Corners.Select(cor => cor.Id).Aggregate((a, b) => a * b).ToString();
        }

        protected override string SolvePartTwo()
        {
            var topLeft = Corners[0];

            topLeft.FlipUpDown();
            while(topLeft.MatchingSides.up != 0 || topLeft.MatchingSides.left != 0)
            {
                topLeft.Rotate();
            }

            var matchedTiles = MatchAllTiles(topLeft);
            var map = new char[96][];

            for(var i = 0; i < matchedTiles.Count; i++)
            {
                for(var j = 0; j < matchedTiles[i][0].Pixels.Length; j++)
                {
                    var row = string.Empty;
                    foreach (var mapPart in matchedTiles[i])
                    {
                        row += mapPart.Pixels[j];
                    }
                    map[i * matchedTiles[i][0].Pixels.Length + j] = row.ToCharArray();
                }
            }

            map = map.Reverse().ToArray();
            var notNessy = map.ToArray();
            for(var turns = 0; turns < 4; turns++)
            {
                notNessy = FindNessy(map, notNessy);
                notNessy = Tile.RotateMatrixCounterClockwise(notNessy);
                map = Tile.RotateMatrixCounterClockwise(map);
            }
            return notNessy.Select(row => row.Count(area => area == '#')).Sum().ToString();
        }

        public static char[][] FindNessy(char[][] map, char[][] notNessy)
        {
            var nessy = new List<(int, int)>
            {
                (18, -1),
                (0,0),(5,0), (6,0), (11,0), (12, 0), (17, 0), (18,0), (19,0),
                (1,1), (4,1), (7,1), (10,1), (13,1), (16,1)
            };
            for (var x = 0; x < map.Length; x++)
            {
                for (var y = 1; y < map[x].Length - 1; y++)
                {
                    // is a starting Nessy, is within bounds to contain nessy and actually contains Nessy
                    if (map[x][y] is '#' && x + 19 < map[x].Length &&
                        nessy.All(delta => map[x + delta.Item1][y + delta.Item2] is '#'))
                    {
                        // Mark Nessy on the Nessy map
                        nessy.ForEach(delta => notNessy[x + delta.Item1][y + delta.Item2] = 'N');
                    }
                }
            }

            return notNessy;
        }

        private List<List<Tile>> MatchAllTiles(Tile start)
        {
            var sortedTiles = new List<List<Tile>>();
            var availableTiles = new List<Tile>(Tiles);
            var current = start;

            while (current != null)
            {
                if (current.MatchingSides.left != 0)
                    current.FlipLeftRight();
                var row = new List<Tile>();
                while (current != null)
                {
                    if (sortedTiles.Count == 0)
                    {
                        if (current.MatchingSides.up != 0)
                        {
                            current.FlipUpDown();
                        }
                    }
                    else
                    {
                        if (current.Up != sortedTiles[^1][row.Count].Down)
                        {
                            current.FlipUpDown();
                        }
                    }
                    row.Add(current);
                    availableTiles.Remove(current);
                    current = FindNextRight(current, availableTiles);
                }
                sortedTiles.Add(row);
                current = FindNextBottom(row[0], availableTiles);
            }

            return sortedTiles;
        }

        private static Tile FindNextRight(Tile tile, List<Tile> availableTiles)
        {
            var toFind = tile.Right;
            var next = availableTiles.FirstOrDefault(t => t.Right == toFind || t.Up == toFind || t.Down == toFind || t.Left == toFind);

            if(next is null)
                return null;
            
            while (next.Left != toFind)
                next.Rotate();

            return next;
        }

        private static Tile FindNextBottom(Tile tile, List<Tile> availableTiles)
        {
            var toFind = tile.Down;
            var next = availableTiles.FirstOrDefault(t => t.Right == toFind || t.Up == toFind || t.Down == toFind || t.Left == toFind);

            if (next is null)
                return null;

            while (next.Up != toFind)
                next.Rotate();

            return next;
        }

        private class Tile
        {
            public readonly long Id;
            public string[] Pixels;
            public int Left;
            public int Right;
            public int Up;
            public int Down;
            public (int up, int right, int down, int left) MatchingSides;

            public Tile(string tileLayout)
            {
                var lines = tileLayout.SplitByNewline();
                Id = long.Parse(Regex.Match(lines[0], @"\d+").Value);
                // Skip the line with the ID
                lines = lines.Skip(1).ToList();

                Pixels = new string[lines.Count - 2];
                for (var i = 1; i < lines.Count - 1; i++)
                {
                    Pixels[i - 1] = lines[i][1..^1];
                }

                // Hash the sides
                Up = Math.Min(Convert.ToInt32(lines[0].Replace('#', '1').Replace('.', '0'), 2),
                    Convert.ToInt32(lines[0].Reverse().Replace('#', '1').Replace('.', '0'), 2));

                Down = Math.Min(Convert.ToInt32(lines[^1].Replace('#', '1').Replace('.', '0'), 2),
                      Convert.ToInt32(lines[^1].Reverse().Replace('#', '1').Replace('.', '0'), 2));

                var leftString = string.Join(string.Empty, new string(lines.Select(line => line[0]).ToArray())).Replace('#', '1').Replace('.', '0');
                Left = Math.Min(Convert.ToInt32(leftString, 2), Convert.ToInt32(leftString.Reverse(), 2));

                var rightString = string.Join(string.Empty, new string(lines.Select(line => line[^1]).ToArray())).Replace('#', '1').Replace('.', '0');
                Right = Math.Min(Convert.ToInt32(rightString, 2), Convert.ToInt32(rightString.Reverse(), 2));
            }

            public Tile Rotate()
            {
                var tmp = Left;
                Left = Up;
                Up = Right;
                Right = Down;
                Down = tmp;

                MatchingSides = (MatchingSides.right, MatchingSides.down, MatchingSides.left, MatchingSides.up);

                Pixels = RotateMatrixCounterClockwise(Pixels);
                return this;
            }

            public Tile FlipUpDown()
            {
                var tmp = Up;
                Up = Down;
                Down = tmp;

                MatchingSides = (MatchingSides.down, MatchingSides.right, MatchingSides.up, MatchingSides.left);

                Pixels = Pixels.Reverse().ToArray();
                return this;
            }

            public Tile FlipLeftRight()
            {
                return Rotate().FlipUpDown().Rotate().Rotate().Rotate();
            }

            private static string[] RotateMatrixCounterClockwise(string[] currentMatrix)
            {
                var oldMatrix = currentMatrix.Select(s => s.ToCharArray()).ToArray();
                return RotateMatrixCounterClockwise(oldMatrix).Select(m => new string(m)).ToArray();
            }

            public static char[][] RotateMatrixCounterClockwise(IReadOnlyList<char[]> oldMatrix)
            {
                var newMatrix = new char[oldMatrix.Count][];
                for (var i = 0; i < oldMatrix.Count; i++)
                {
                    newMatrix[i] = new char[oldMatrix[0].Length];
                }

                var newRow = 0;
                for (var oldColumn = oldMatrix.Count - 1; oldColumn >= 0; oldColumn--)
                {
                    var newColumn = 0;
                    for (var oldRow = 0; oldRow < oldMatrix[0].Length; oldRow++)
                    {
                        newMatrix[newRow][newColumn] = oldMatrix[oldRow][oldColumn];
                        newColumn++;
                    }
                    newRow++;
                }
                return newMatrix;
            }

            public (int, int, int, int) MatchesSides(List<Tile> otherTiles)
            {
                int left = 0, right = 0, up = 0, down = 0;

                foreach (var otherTile in otherTiles.Where(otherTile => otherTile.Id != Id))
                {
                    if (Left == otherTile.Left || Left == otherTile.Right || Left == otherTile.Up || Left == otherTile.Down )
                    {
                        left += 1;
                    }
                    if (Right == otherTile.Left || Right == otherTile.Right || Right == otherTile.Up || Right == otherTile.Down)
                    {
                        right += 1;
                    }
                    if (Up == otherTile.Left || Up == otherTile.Right || Up == otherTile.Up || Up == otherTile.Down)
                    {
                        up += 1;
                    }
                    if (Down == otherTile.Left || Down == otherTile.Right || Down == otherTile.Up || Down == otherTile.Down)
                    {
                        down += 1;
                    }
                }

                MatchingSides = (up, right, down, left);
                return (up, right, down, left);
            }
        }
    }
}
