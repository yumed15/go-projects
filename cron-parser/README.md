# Cron Expression Parser

Given a cron expression and a command, outputs the parsed cron as formatted table with the field name taking the first 14 columns and
the times as a space-separated list following it.

Currently, the application doesn't support:
- alpha characters for month (JAN-DEC) or day of week (SUN-SAT)
- the special characters `L`, `W` and `#`
- and predefined cron expressions (e.g. `@annually`, `@yearly` etc)

#Run
Minimum requirements
```
go version go1.20
```

Run the application from cli
```
$ go build main.go
$ go run main.go '*/15' 0 '1,15' '*' '1-5' /usr/bin/find
```

NB: It's a common problem in Unix-like shells (e.g. zsh) When you run the command with special characters (e.g. */15) the shell interprets it as a glob pattern and tries to expand it to match files or directories in the current directory.

To avoid this issue, you should enclose the `*/15` part in quotes `'*/15'` or escape it `\*/15`

#Usage
The cron-parser utility takes a single argument; a cron expression as follows:
```
cron-parser "*/15 0 1,15 * 1-5 /usr/bin/find"
```

It then prints a summary of the execution schedule of the command as well as the command itself. For example:
```
minute        0 15 30 45
hour          0
day of month  1 15
month         1 2 3 4 5 6 7 8 9 10 11 12
day of week   1 2 3 4 5
command       /usr/bin/find
```

# Implementation
The cron interpretation is based on the [wiki documentation](https://en.wikipedia.org/wiki/Cron#CRON_expression): 

    Field name     Mandatory?   Allowed values    Allowed special characters
    ----------     ----------   --------------    --------------------------
    Seconds        No           0-59              * / , -
    Minutes        Yes          0-59              * / , -
    Hours          Yes          0-23              * / , -
    Day of month   Yes          1-31              * / , - L W
    Month          Yes          1-12 or JAN-DEC   * / , -
    Day of week    Yes          0-6 or SUN-SAT    * / , - L #
    Year           No           1970–2099         * / , -

#### Asterisk ( * )
The asterisk indicates that the cron expression matches for all values of the field. 
e.g., using an asterisk in the 4th field (month) indicates every month.

#### Slash ( / )
Slashes describe increments of ranges. For example `3-59/15` in the minute field indicate the third minute of the hour and every 15 minutes thereafter. The form `*/...` is equivalent to the form "first-last/...", that is, an increment over the largest possible range of the field.

#### Comma ( , )
Commas are used to separate items of a list. For example, using `MON,WED,FRI` in the 5th field (day of week) means Mondays, Wednesdays and Fridays.

#### Hyphen ( - )
Hyphens define ranges. For example, 2000-2010 indicates every year between 2000 and 2010 AD, inclusive.

#### L
`L` stands for "last". When used in the day-of-week field, it allows you to specify constructs such as "the last Friday" (`5L`) of a given month. In the day-of-month field, it specifies the last day of the month.

#### W
The `W` character is allowed for the day-of-month field. This character is used to specify the business day (Monday-Friday) nearest the given day. As an example, if you were to specify `15W` as the value for the day-of-month field, the meaning is: "the nearest business day to the 15th of the month."

So, if the 15th is a Saturday, the trigger fires on Friday the 14th. If the 15th is a Sunday, the trigger fires on Monday the 16th. If the 15th is a Tuesday, then it fires on Tuesday the 15th. However if you specify `1W` as the value for day-of-month, and the 1st is a Saturday, the trigger fires on Monday the 3rd, as it does not 'jump' over the boundary of a month's days.

The `W` character can be specified only when the day-of-month is a single day, not a range or list of days.

The `W` character can also be combined with `L`, i.e. `LW` to mean "the last business day of the month."

#### Hash ( # )
`#` is allowed for the day-of-week field, and must be followed by a number between one and five. It allows you to specify constructs such as "the second Friday" of a given month.

Predefined cron expressions
---------------------------
(Copied from <https://en.wikipedia.org/wiki/Cron#Predefined_scheduling_definitions>, with text modified according to this implementation)

    Entry       Description                                                             Equivalent to
    @annually   Run once a year at midnight in the morning of January 1                 0 0 0 1 1 * *
    @yearly     Run once a year at midnight in the morning of January 1                 0 0 0 1 1 * *
    @monthly    Run once a month at midnight in the morning of the first of the month   0 0 0 1 * * *
    @weekly     Run once a week at midnight in the morning of Sunday                    0 0 0 * * 0 *
    @daily      Run once a day at midnight                                              0 0 0 * * * *
    @hourly     Run once an hour at the beginning of the hour                           0 0 * * * * *
    @reboot     Not supported



## Possible Expressions
`*/...`