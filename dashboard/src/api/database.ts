import useSWR from "swr";

import admin_api from "./api";
import { DatabaseConnectConfig } from "./openapi";

const databaseFetcher = async (key) =>{
    return (await admin_api.getDatabaseConfigureList()).data.dataSourceConfig
 };
 
 
const useDatabase = ()=> {
    const {data,isLoading,error, mutate} = useSWR("database",databaseFetcher)

    const createDatabase = (database:DatabaseConnectConfig)=>{
        mutate(
            [...data,database],
            false
        )
        return admin_api.createDatabase(database)
    }
    
    const updateDatabase = (business)=>{
        console.log(business)
    }

    const deleteDatabase = (id:string)=>{
        return admin_api.deleteDatabase(id)
    }

    return {
        database:data||[],
        isLoading,
        isError:error,
        createDatabase,
        updateDatabase,
        deleteDatabase
    }
}
export default useDatabase;