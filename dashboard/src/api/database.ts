import useSWR from "swr";
import fetcher from "./fetcher";
import axios from "axios";

const useDatabase = ()=> {
    const {data,isLoading,error, mutate} = useSWR("/admin/v1/database",fetcher)

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
        database:data,
        isLoading,
        isError:error,
        createDatabase,
        updateDatabase,
        deleteDatabase
    }
}
export default useDatabase;