# sync

Notes on the synchronization process by Todoist.

## How do edited, completed, deleted, rescheduled and uncompleted items respond?

Edited: Full object
Completed: `is_completed`, full object
Deleted: `is_deleted`, full object

TODO: Rescheduled, uncompleted

## How should synchronized items be identified/updated?

The `id` detail.

TODO: It is not well specified in the docs what the meaning of `v2_id` is.

## What data structure should be used to store synchronized items?

A dictionary with the `id` as the key.

## How and when should synchronization be triggered?

With the cron schedule, first thing.

## Is there a limit to the delay between synchronization?

Likely: if there is a limit, it should be handled automatically by the `full_sync` attribute.