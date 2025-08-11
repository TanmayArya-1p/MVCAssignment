import { API_URL } from "../config";
import axios from "axios";

export async function getMyOrders(limit, offset) {
    const params = {};
    if (limit != null) params.limit = limit;
    if (offset != null) params.offset = offset;

    const response = await axios.get(`${API_URL}/api/orders/my`, {params,withCredentials: true});
    console.log(response.data);
    return response.data;   
}