using System;

public static class Leap
{
    public static bool IsLeapYear(int year)
    {
        var returner = year % 4 == 0;

        switch (year % 100)
        {
            case 0 when year % 400 == 0:
                returner = true;
                break;
            case 0 when year % 400 != 0:
                returner = false;
                break;
        }
        
        return returner;
    }
}