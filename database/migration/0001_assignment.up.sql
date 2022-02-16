CREATE TABLE users(
                      id UUID,
                      name TEXT NOT NULL,
                      userid text not null PRIMARY KEY,
                      email TEXT NOT NULL,
                      mobile_no TEXT NOT NULL,
                      password TEXT NOT NULL,
                      unique (email, mobile_no,id)
);
CREATE TABLE session (
                         id UUID,
                         userid text not null PRIMARY KEY,
                         password VARCHAR(50) NOT NULL,
                         createdAt TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
                         UNIQUE (id),
                         CONSTRAINT fk_userid
                             FOREIGN KEY(userid)
                             REFERENCES users(userid)
);