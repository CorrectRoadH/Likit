import useSWR from "swr";
import axios from "axios";

import admin_api from "./api";

const databaseFetcher = async (key) =>{
    return (await admin_api.getDatabaseConfigureList()).data.dataSourceConfig
 };
 
 
const useDatabase = ()=> {
    const {data,isLoading,error, mutate} = useSWR("database",databaseFetcher)

    const createDatabase = (database)=>{
        mutate(
            [...data,database],
            false
        )
        return axios.post("/admin/v1/database",database)
    }
    
    const updateDatabase = (business)=>{
        console.log(business)
    }

    const deleteDatabase = (business)=>{
        console.log(business)
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