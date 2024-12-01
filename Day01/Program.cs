// const string filename = "../../../test.txt";
const string filename = "../../../../../AdventOfCodeInputs/2024/Day01/input.txt";

var left = new List<int>();
var right = new List<int>();

var sr = new StreamReader(filename);
while (sr.ReadLine() is { } line)
{
    var tokens = line.Split(' ', StringSplitOptions.RemoveEmptyEntries);
    if (tokens.Length < 2) break;
    // Console.WriteLine($"{tokens[0]} | {tokens[1]}");
    left.Add(int.Parse(tokens[0]));
    right.Add(int.Parse(tokens[1]));
}

left.Sort();
right.Sort();

var total = 0;
for (var i = 0; i < left.Count; i++)
{
    total += Math.Abs(right[i] - left[i]);
}

Console.WriteLine($"Part 1: {total}");

var similarity = 0;
var j = 0;
for (var i = 0; i < left.Count; i++)
{

    var currentId = left[i];
    var rightCount = 0;
    for (; j < right.Count; j++)
    {
        if (right[j] < currentId)
        {
            continue;
        } else if (right[j] > currentId)
        {
            break;
        }

        rightCount++;
    }

    do
    {
        similarity += left[i] * rightCount;
        i++;
    } while (i < left.Count && left[i] == left[i-1]);
    i--;
}

Console.WriteLine($"Part 2: {similarity}");
