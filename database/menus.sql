-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: sequences, indices, triggers. Do not use it as a backup.

CREATE TABLE [dbo].[menus] (
    [id] bigint,
    [name] nvarchar(255) NOT NULL,
    [created_at] datetimeoffset,
    [updated_at] datetimeoffset,
    [deleted_at] datetimeoffset,
    PRIMARY KEY ([id])
);