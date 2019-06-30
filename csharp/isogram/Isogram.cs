using System;
using System.Collections;

public static class Isogram
{
    public static bool IsIsogram(string word)
    {
        var letters = new ArrayList();
        foreach (var c in word.ToLower())
        {
            if (letters.Contains(c))
            {
                return false;
            }

            if (c.Equals('-') | c.Equals(' '))
            {
                continue;
            }

            letters.Add(c);
        }
        return true;
    }
}
