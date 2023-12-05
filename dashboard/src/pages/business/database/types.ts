interface DatabaseConnectionConfig{
    id: string;
    title: string;
    databaseType: string;
    host: string;
    port: number;
    username: string;
    password: string;
    database: string;
}

export type {DatabaseConnectionConfig}