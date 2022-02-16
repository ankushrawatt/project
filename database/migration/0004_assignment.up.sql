DROP TABLE session;
CREATE TABLE session(
                        id TEXT NOT NULL,
                        userid TEXT NOT NULL ,
                        createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                        UNIQUE (id),
                        CONSTRAINT fk_userid
                            FOREIGN KEY(userid)
                                REFERENCES users(userid)
)