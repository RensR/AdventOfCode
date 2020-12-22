using System.Collections.Generic;
using System.Linq;

namespace AdventOfCode.Solutions.Year2020
{
    class Day22 : ASolution
    {
        private readonly Queue<int> _playerOne = new();
        private readonly Queue<int> _playerTwo = new();
        private readonly string[] _decks;

        public Day22() : base(22, 2020, "Crab Combat")
        {
            _decks = Input.Split("\n\n");
        }

        protected override string SolvePartOne()
        {
            DrawDeck();
            while (_playerOne.Count > 0 && _playerTwo.Count > 0)
            {
                PlayNormalRound(_playerOne, _playerTwo);
            }

            return GetScore().ToString();
        }

        protected override string SolvePartTwo()
        {
            DrawDeck();
            PlaySubGame(_playerOne, _playerTwo);
            return GetScore().ToString();
        }
        
        private void DrawDeck()
        {
            foreach (var card in _decks[0].Split("\n").Skip(1).Select(int.Parse))
            {
                _playerOne.Enqueue(card);
            }

            foreach (var card in _decks[1].Split("\n").Skip(1).Select(int.Parse))
            {
                _playerTwo.Enqueue(card);
            }
        }

        private static void PlayNormalRound(Queue<int> subDeckOne, Queue<int> subDeckTwo)
        {
            var cardOne = subDeckOne.Dequeue();
            var cardTwo = subDeckTwo.Dequeue();

            if (cardOne > cardTwo)
            {
                subDeckOne.Enqueue(cardOne);
                subDeckOne.Enqueue(cardTwo);
            }
            else
            {
                subDeckTwo.Enqueue(cardTwo);
                subDeckTwo.Enqueue(cardOne);
            }
        }

        private int GetScore()
        {
            var winner = _playerOne.Count > 0 ? _playerOne : _playerTwo;
            var score = 0;
            for (var i = winner.Count; i > 0; i--)
            {
                score += i * winner.Dequeue();
            }

            return score;
        }

        private static bool PlaySubGame(Queue<int> subDeckOne, Queue<int> subDeckTwo)
        {
            HashSet<(string, string)> seenDecks = new();
            while (subDeckOne.Count > 0 && subDeckTwo.Count > 0)
            {
                if (CheckRecursion(subDeckOne, subDeckTwo, seenDecks))
                {
                    return true;
                }
                var cardOne = subDeckOne.Peek();
                var cardTwo = subDeckTwo.Peek();

                if (cardOne <= subDeckOne.Count - 1 && cardTwo <= subDeckTwo.Count - 1)
                {
                    subDeckOne.Dequeue();
                    subDeckTwo.Dequeue();
                    
                    if (PlaySubGame(
                        new Queue<int>(subDeckOne.Take(cardOne)),
                        new Queue<int>(subDeckTwo.Take(cardTwo))
                    ))
                    {
                        subDeckOne.Enqueue(cardOne);
                        subDeckOne.Enqueue(cardTwo);
                    }
                    else
                    {
                        subDeckTwo.Enqueue(cardTwo);
                        subDeckTwo.Enqueue(cardOne);
                    }
                }
                else
                {
                    PlayNormalRound(subDeckOne, subDeckTwo);
                }
            }

            return subDeckOne.Count > 0;
        }

        private static bool CheckRecursion(IEnumerable<int> subDeckOne, IEnumerable<int> subDeckTwo, ISet<(string, string)> seenDecks)
        {
            var deckOne = subDeckOne.Select(num => num.ToString()).Aggregate((cardA, cardB) => cardA + ' ' + cardB);
            var deckTwo = subDeckTwo.Select(num => num.ToString()).Aggregate((cardA, cardB) => cardA + ' ' + cardB);

            if (seenDecks.Contains((deckOne, deckTwo)))
                return true;

            seenDecks.Add((deckOne, deckTwo));
            return false;
        }
    }
}
