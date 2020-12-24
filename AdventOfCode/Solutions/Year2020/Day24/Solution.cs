using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day24 : ASolution
    {
        private readonly List<List<string>> _directions = new();
        private Dictionary<(int, int), bool> _tiles = new();

        public Day24() : base(24, 2020, "Lobby Layout")
        {
            foreach (var line in Input.SplitByNewline())
            {
                var lineDirections = new List<string>();
                for (var i = 0; i < line.Length; i++)
                {
                    switch (line[i])
                    {
                        case 'e':
                        case 'w':
                            lineDirections.Add(line[i].ToString());
                            break;
                        case 's':
                        case 'n':
                            lineDirections.Add($"{line[i]}{line[i + 1]}");
                            i++;
                            break;
                    }
                }

                _directions.Add(lineDirections);
            }
        }

        protected override string SolvePartOne()
        {
            foreach (var direction in _directions)
            {
                var finalTile = direction.Aggregate((x: 0, y: 0),
                    (current, directionStep) => directionStep switch
                    {
                        "e" => (current.x + 1, current.y),
                        "se" => (current.x + 1, current.y - 1),
                        "sw" => (current.x, current.y - 1),
                        "w" => (current.x - 1, current.y),
                        "nw" => (current.x - 1, current.y + 1),
                        "ne" => (current.x, current.y + 1),
                        _ => current
                    });

                _tiles[finalTile] = !_tiles.ContainsKey(finalTile) || !_tiles[finalTile];
            }

            return _tiles.Values.Count(tile => tile).ToString();
        }

        protected override string SolvePartTwo()
        {
            for (var i = 1; i <= 100; i++)
            {
                _tiles = UpdateFloor(_tiles);
            }

            return _tiles.Values.Count(tile => tile).ToString();
        }

        private static Dictionary<(int, int), bool> UpdateFloor(Dictionary<(int, int), bool> floorTiles)
        {
            var newFloor = new Dictionary<(int, int), bool>();

            // foreach black tile
            foreach (var floorTilesKey in floorTiles.Keys
                .Where(floorTilesKey => floorTiles[floorTilesKey]))
            {
                newFloor = UpdateTiles(floorTiles, newFloor, floorTilesKey);
            }

            return newFloor;
        }

        private static Dictionary<(int, int), bool> UpdateTiles(IReadOnlyDictionary<(int, int), bool> floorTiles,
            Dictionary<(int, int), bool> newFloor, (int x, int y) tile)
        {
            // All directions in a hexagonal grid and (0,0) to also reference itself for lonely tiles
            // that should become white because they are not near any other black tiles.
            var lookAround = new[] {(0, 0), (1, 0), (1, -1), (0, -1), (-1, 0), (-1, 1), (0, 1)};
            
            foreach (var (dx, dy) in lookAround)
            {
                var current = (x: tile.x + dx, y: tile.y + dy);
                if (newFloor.ContainsKey(current))
                    continue;

                var nextToBlack = 0;
                // skip the first element of lookAround as it references itself 
                foreach (var (ddx, ddy) in lookAround[1..])
                {
                    var around = (current.x + ddx, current.y + ddy);
                    if (floorTiles.ContainsKey(around) && floorTiles[around])
                    {
                        nextToBlack++;
                    }
                }
                
                // if current is black and 1 or 2 stay black. If current is white and 2 become black
                // otherwise be white.
                newFloor[current] = floorTiles.ContainsKey(current) && floorTiles[current]
                    ? nextToBlack is 1 || nextToBlack is 2
                    : nextToBlack is 2;
            }

            return newFloor;
        }
    }
}
