CREATE TABLE image_metadata (
                                id VARCHAR(255) PRIMARY KEY,
                                filename VARCHAR(255) NOT NULL default '',
                                filepath VARCHAR(255) NOT NULL default '',
                                size_file int NOT NULL default  0,
                                mimetype VARCHAR(255) NOT NULL default '',
                                upload_date bigint DEFAULT 0
);
