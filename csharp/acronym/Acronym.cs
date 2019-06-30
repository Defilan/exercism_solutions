using System;
using System.Linq;
using System.Runtime.InteropServices.ComTypes;
using System.Text;

public static class Acronym
{
    public static string Abbreviate(string phrase)
    {
        string[] tester = phrase.ToUpper().Split(" ");
        StringBuilder returner = new StringBuilder();
        for (int i = 0; i < tester.Length; i++)
        {
            if (tester[i] != "-")
            {
                returner.Append(tester[i].First());
            }
        }

        return returner.ToString();
    }
}