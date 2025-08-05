-- This script only contains the table creation statements and does not fully represent the table in the database. It's still missing: sequences, indices, triggers. Do not use it as a backup.

CREATE TABLE [dbo].[items] (
    [id] bigint,
    [menu_id] bigint,
    [name] nvarchar(255) NOT NULL,
    [index_order] float,
    [created_at] datetimeoffset,
    [updated_at] datetimeoffset,
    [deleted_at] datetimeoffset,
    CONSTRAINT [fk_items_menu] FOREIGN KEY ([menu_id]) REFERENCES [dbo].[menus]([id]),
    PRIMARY KEY ([id])
);