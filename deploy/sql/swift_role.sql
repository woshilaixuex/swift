IF OBJECT_ID('role_salt', 'U') IS NOT NULL
DROP TABLE role_salt;
CREATE TABLE role_salt (
                           id bigint IDENTITY(1,1) NOT NULL,
                           role_id bigint NOT NULL,
                           salt varchar(128) NOT NULL,
                           PRIMARY KEY (id),
                           CONSTRAINT idx_role_id UNIQUE (role_id)
)

IF OBJECT_ID('role', 'U') IS NOT NULL
DROP TABLE role;
CREATE TABLE role (
                      id bigint IDENTITY(1,1) NOT NULL,
                      create_time datetime NOT NULL DEFAULT GETDATE(),
                      update_time datetime NOT NULL DEFAULT GETDATE(),
                      delete_time datetime NOT NULL DEFAULT GETDATE(),
                      del_state tinyint NOT NULL DEFAULT '0',
                      version bigint NOT NULL DEFAULT '0',
                      role tinyint NOT NULL DEFAULT '0',
                      mobile char(11) NOT NULL DEFAULT '',
                      password varchar(255) NOT NULL DEFAULT '',
                      nickname varchar(255) NOT NULL DEFAULT '',
                      PRIMARY KEY (id),
                      CONSTRAINT idx_mobile UNIQUE (mobile)
)