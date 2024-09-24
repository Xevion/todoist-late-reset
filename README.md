# todoist-late-reset

A basic Go program to reschedule certain tasks in Todoist if they're completed 'late'.

---

For tasks that are scheduled daily, or repeat at least 2 days in a row, this program helps reschedule them automatically if they're completed after midnight.

For example, if you have a task to complete every weekday (Mon-Fri) and you complete it at 2AM on Tuesday, this program will reschedule it for the same day (Tuesday).

- Yes, rescheduling it is simple and easy, but it'd be nicer to have confidence when completing tasks late that they'll be available the next day.