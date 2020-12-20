using System;
using System.Collections.Generic;
using System.Linq;
using System.Text.RegularExpressions;

namespace AdventOfCode.Solutions.Year2020
{

    class Day20 : ASolution
    {
        List<Tile> Tiles;
        public Day20() : base(20, 2020, "Jurassic Jigsaw")
        {
            Tiles = Input.Split("\n\n").Select(tileLines => new Tile(tileLines)).ToList();
        }

        protected override string SolvePartOne()
        {
            var corners = new List<Tile>();
            foreach (var tile in Tiles)
            {
                var (l, r, u, d) = tile.MatchesSides(Tiles);
                if ((l == 0 || r == 0) && (u == 0 || d == 0))
                    corners.Add(tile);
            }
            return corners.Select(cor => cor.Id).Aggregate((a, b) => a * b).ToString();
        }

        protected override string SolvePartTwo()
        {
            return null;
        }

        class Tile
        {
            public readonly long Id;
            bool[,] Pixels;
            int Left;
            int Right;
            int Up;
            int Down;

            public Tile(string tileLayout)
            {
                var lines = tileLayout.SplitByNewline();
                Id = long.Parse(Regex.Match(lines[0], @"\d+").Value);
                lines = lines.Skip(1).ToList();


                Pixels = new bool[lines.Count, lines[0].Length];

                for(int i = 0; i < lines.Count; i++)
                {
                    for (int j = 0; j < lines[i].Length; j++)
                    {
                        Pixels[i, j] = lines[i][j] == '#';
                    }
                }

                Up = Math.Min(Convert.ToInt32(lines[0].Replace('#', '1').Replace('.', '0'), 2),
                    Convert.ToInt32(lines[0].Reverse().Replace('#', '1').Replace('.', '0'), 2));

                Down = Math.Min(Convert.ToInt32(lines[^1].Replace('#', '1').Replace('.', '0'), 2),
                      Convert.ToInt32(lines[^1].Reverse().Replace('#', '1').Replace('.', '0'), 2));


                var leftString = string.Join(string.Empty, new string(lines.Select(line => line[0]).ToArray())).Replace('#', '1').Replace('.', '0');
                Left = Math.Min(Convert.ToInt32(leftString, 2), Convert.ToInt32(leftString.Reverse(), 2));

                var rightString = string.Join(string.Empty, new string(lines.Select(line => line[^1]).ToArray())).Replace('#', '1').Replace('.', '0');
                Right = Math.Min(Convert.ToInt32(rightString, 2), Convert.ToInt32(rightString.Reverse(), 2));
            }

            public (int, int, int, int) MatchesSides(List<Tile> otherTiles)
            {
                int left = 0, right = 0, up = 0, down = 0;

                foreach(var otherTile in otherTiles)
                {
                    if (otherTile.Id == Id)
                        continue;
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

                return (left, right, up, down);
            }
        }
    }
}
