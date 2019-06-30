public static class ResistorColor
{
    public static int ColorCode(string color)
    {
        for(int i = 0; i < Colors().Length; i++)
        {
            if (Colors()[i] == color)
            {
                return i;
            }
        }
        return 0;
    }

    public static string[] Colors()
    {
        return new[] {"black", "brown", "red", "orange", "yellow", "green", "blue", "violet", "grey", "white"};
    }
}