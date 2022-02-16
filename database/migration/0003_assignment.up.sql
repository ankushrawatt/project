DROP TABLE session;
CREATE TABLE session(
                        id UUID,
                        userid text not null,
                        createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                        UNIQUE (id),
                        CONSTRAINT fk_userid
                            FOREIGN KEY(userid)
                                REFERENCES users(userid)
)