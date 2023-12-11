import useSWR from "swr";
import admin_api from "./api";

const openapiFetcher = async (url) =>{
   return (await admin_api.getBusinesses()).data.businesses
};

const useBusiness = ()=> {
    const { data,isLoading,error, mutate } = useSWR("/admin/v1/businesses",openapiFetcher)

    const createBusiness = async (business)=>{
        mutate([...data,business],false)
        return admin_api.createBusiness(business)
    }
    
    const updateBusiness = (business)=>{
        mutate(data.map((item)=>{
            if(item.id == business.id){
                return business
            }else{
                return item
            }
        }),false)

        return admin_api.updateBusiness(business)
    }

    const deleteBusiness = (business)=>{
        mutate(data.filter((item)=>item.id != business.id),false)

        return admin_api.deleteBusiness(business.id)
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