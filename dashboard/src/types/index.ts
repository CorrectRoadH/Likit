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
type BusinessType = {
    title: string;
    id: string;
    type: string;
  };

  

export type { DatabaseConnectionConfig, BusinessType }