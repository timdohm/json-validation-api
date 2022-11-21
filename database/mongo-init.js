db.createUser(
    {
        user: "api",
        pwd: "testpass",
        roles: [
            {
                role: "readWrite",
                db: "api"
            }
        ]
    }
);