CREATE TABLE bots(
    "id" character varying(255) PRIMARY KEY NOT NULL,
    "ip" character varying(255) NOT NULL,
    "who_am_i" character varying(255) NOT NULL,
    "os" character varying(255) NOT NULL,
    "install_date" character varying(255) NOT NULL,
    "admin" boolean NOT NULL,
    "av" character varying(255) NOT NULL,
    "cpu" character varying(255) NOT NULL,
    "gpu" character varying(255) NOT NULL,
    "version" character varying(255) NOT NULL, 
    "last_checkin" character varying(255),
    "last_command" character varying(255),
    "new_command" character varying(255) 
);