import useSWR from "swr";
import fetcher from "./fetcher";
import axios from "axios";

const useBusiness = ()=> {
    const {data,isLoading,error, mutate} = useSWR("/admin/v1/businesses",fetcher)

    const createBusiness = async (business)=>{
        mutate([...data,business],false)
        return axios.post('/admin/v1/business', business)
    }
    
    const updateBusiness = (business)=>{
        mutate([data.map((item)=>{
            if(item.id == business.id){
                return business
            }else{
                return item
            }
        })],false)
        return axios.put('/admin/v1/business',business)
    }

    const deleteBusiness = (business)=>{
        return axios.delete('/admin/v1/business')
    }

    return {
        businesses:data,
        isLoading,
        isError:error,
        createBusiness,
        updateBusiness,
        deleteBusiness
    }
}
export default useBusiness;