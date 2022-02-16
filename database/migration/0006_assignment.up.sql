DROP TABLE session;
DROP TABLE users;

CREATE TABLE users(
                      id UUID,
                      name VARCHAR(25) NOT NULL,
                      userid VARCHAR(20) NOT NULL PRIMARY KEY,
                      email VARCHAR(50) NOT NULL,
                      mobile VARCHAR(10) NOT NULL,
                      password VARCHAR(50) NOT NULL,
                      unique (email, mobile,id)
);
CREATE TABLE session(
                        id TEXT NOT NULL,
                        userid TEXT NOT NULL ,
                        createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                        UNIQUE (id),
                        CONSTRAINT fk_userid
                            FOREIGN KEY(userid)
                                REFERENCES users(userid)
)