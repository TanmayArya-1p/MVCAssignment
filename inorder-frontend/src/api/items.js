import { API_URL } from "../config";
import axios from "axios";

export async function getAllItems(limit,offset) {
    const params = {};
    if (limit != null) params.limit = limit;
    if (offset != null) params.offset = offset;

    const response = await axios.get(`${API_URL}/api/items`, {params,withCredentials: true});
    console.log(response.data);
    return response.data;
}



export async function getItemsOfTags(tags) {
    const response = await axios.get(`${API_URL}/api/items/bytags?tags=`+tags.join(","), {withCredentials: true});
    console.log(response.data);
    return response.data;
}