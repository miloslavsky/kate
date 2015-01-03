drop table `users`
    DROP TABLE IF EXISTS `users`

drop table `blogposts`
    DROP TABLE IF EXISTS `blogposts`

drop table `followers`
    DROP TABLE IF EXISTS `followers`

drop table `comments`
    DROP TABLE IF EXISTS `comments`

create table `users` 
    -- --------------------------------------------------
    --  Table Structure for `Kate/models.User`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `users` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL,
        `password` varchar(255) NOT NULL
    ) ENGINE=InnoDB;

create table `blogposts` 
    -- --------------------------------------------------
    --  Table Structure for `Kate/models.Blogpost`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `blogposts` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `name` varchar(255) NOT NULL,
        `content` text NOT NULL,
        `owner` bigint NOT NULL
    ) ENGINE=InnoDB;

create table `followers` 
    -- --------------------------------------------------
    --  Table Structure for `Kate/models.Follower`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `followers` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `follow` bigint NOT NULL,
        `follower` bigint NOT NULL
    ) ENGINE=InnoDB;

create table `comments` 
    -- --------------------------------------------------
    --  Table Structure for `Kate/models.Comment`
    -- --------------------------------------------------
    CREATE TABLE IF NOT EXISTS `comments` (
        `id` bigint AUTO_INCREMENT NOT NULL PRIMARY KEY,
        `content` text NOT NULL,
        `owner_name` varchar(255) NOT NULL,
        `owner` bigint NOT NULL,
        `post` bigint NOT NULL
    ) ENGINE=InnoDB;

