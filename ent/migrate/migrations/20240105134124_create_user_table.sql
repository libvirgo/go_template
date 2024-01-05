-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `create_time` datetime NOT NULL, `update_time` datetime NOT NULL, `name` text NULL, `email` text NULL, `password` text NOT NULL);
-- Create index "user_email" to table: "users"
CREATE UNIQUE INDEX `user_email` ON `users` (`email`);
