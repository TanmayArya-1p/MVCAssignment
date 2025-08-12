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


export async function getAllOrderedItems() {
    const response = await axios.get(`${API_URL}/api/orders/items`, {withCredentials: true});
    console.log(response.data);
    return response.data;
}

export async function bumpOrderItemStatus(itemId) {
    const response = await axios.post(`${API_URL}/api/orders/item/${itemId}/bump`,{}, { withCredentials: true });
    console.log(response.data);
    return response.data;
}

export async function deleteItem(itemId) {
    const response = await axios.delete(`${API_URL}/api/items/${itemId}`, { withCredentials: true });
    console.log(response.data);
    return response.data;
}


export async function uploadImage(image) {
    const formData = new FormData()
    formData.append("image", image);
    const response = await axios.post(`${API_URL}/api/items/upload`, formData, {
        headers: {
            'Content-Type': 'multipart/form-data'
        },
        withCredentials: true
    });
    console.log(response.data);
    return response.data;
}


export async function createItem({name,price,description,tags,image}) {
    const response = await axios.post(`${API_URL}/api/items`, {
        name,
        price: parseFloat(price),
        description,
        tags: tags,
        image
    }, { withCredentials: true });
    console.log(response.data);
    return response.data;
}   