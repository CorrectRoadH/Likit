import axios from "axios";

const fetcher =  (url: string) => {
    return axios
    .get(url)
    .then((res:any) => res.data);
}


export default fetcher;