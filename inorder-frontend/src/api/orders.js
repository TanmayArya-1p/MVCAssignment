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

export async function getAllOrders(limit, offset) {
    const params = {};
    if (limit != null) params.limit = limit;
    if (offset != null) params.offset = offset;

    const response = await axios.get(`${API_URL}/api/orders`, {params,withCredentials: true});
    console.log(response.data);
    return response.data;   
}

export async function getOrder(orderId) {
    const response = await axios.get(`${API_URL}/api/orders/${orderId}`, {withCredentials: true});
    console.log(response.data);
    return response.data;
}

export async function resolveBill(orderID,markAsBilled=false) {
    const response = await axios.get(`${API_URL}/api/orders/${orderID}/bill?resolve=`+markAsBilled, {withCredentials: true});
    console.log(response.data);
    return response.data;
}

export async function markAsPaid(orderID, amountPaid) {
    const response = await axios.post(`${API_URL}/api/orders/${orderID}/bill/pay`, {"amount": parseInt(amountPaid)}, {withCredentials: true});
    console.log(response.data);
    return response.data;
}

export async function addItemToOrder(orderId, item_id, quantity, instructions) {
    const response = await axios.post(`${API_URL}/api/orders/${orderId}/items`, { item_id : parseInt(item_id),quantity: quantity, instructions }, { withCredentials: true });
    return response.data;
}


export async function createOrder(tableNo) {
    const response = await axios.post(`${API_URL}/api/orders`, { table_no: parseInt(tableNo) }, { withCredentials: true });
    console.log(response.data);
    return response.data;
}

export async function deleteOrder(orderID) {
    const response = await axios.delete(`${API_URL}/api/orders/${orderID}`, { withCredentials: true });
    console.log(response.data);
    return response.data;
}